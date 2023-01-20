package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

const (
	baseUrl = "https://github.com/alexanderlukanin13/coolname"
	dataUrl = baseUrl + "/raw/$REF/coolname/data/"
	gitUrl  = baseUrl + ".git"
	refsUrl = gitUrl + "/info/refs?service=git-upload-pack" // smart HTTP protocol
	target  = "master"
)

// CLI entrypoint for go generate
func main() {
	ref, err := commit(target)
	if err != nil {
		log.Fatal(err)
	}
	if ref == upstreamRef {
		fmt.Printf("Word lists up to date: upstream branch %s at commit %s\n", target, ref)
		return
	}
	os.WriteFile("codegen/upstream.ref", []byte(ref), 0666)

	errors := make(chan error)
	var wg sync.WaitGroup

	for i := -1; i < len(upstreamLists); i++ {
		var file string
		if i < 0 {
			file = upstreamConfig
		} else {
			file = upstreamLists[i] + ".txt"
		}
		wg.Add(1)
		go func() {
			defer wg.Done()
			var err error
			err = fetch(file, ref)
			if err != nil {
				errors <- err
			}
		}()
	}
	go func() {
		wg.Wait()
		errors <- nil
	}()
	err = <-errors
	if err != nil {
		log.Fatal(err)
	}

	err = convert()
	if err != nil {
		log.Fatal(err)
	}
}

// Resolve upstream branch/tag name to a commit hash
func commit(head string) (hash string, err error) {
	client := http.Client{
		Timeout: 5 * time.Second,
	}
	resp, err := client.Get(refsUrl)
	if err != nil {
		return hash, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return hash, fmt.Errorf("HTTP error: %s", resp.Status)
	}

	scanner := bufio.NewScanner(resp.Body)
	target := [...][]byte{
		[]byte(" " + head),
		[]byte("/" + head),
	}
	for scanner.Scan() {
		line := scanner.Bytes()
		for i := 0; i < len(target); i++ {
			if bytes.HasSuffix(line, target[i]) {
				return string(bytes.Split(line, []byte(" "))[0][4:]), nil
			}
		}
	}
	return hash, fmt.Errorf("head not found: %s", head)
}

// Fetch data files from upstream
func fetch(filename string, commit string) (err error) {
	url := strings.ReplaceAll(dataUrl, "$REF", commit) + filename
	fmt.Printf("Fetching %s\n", filename)

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return fmt.Errorf("HTTP error: %s (%s)", resp.Status, url)
	}

	output, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer output.Close()

	_, err = io.Copy(output, resp.Body)
	if err != nil {
		return err
	}
	return nil
}
