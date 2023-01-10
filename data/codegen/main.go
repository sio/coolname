package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/sio/coolname/data"
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
	if ref == data.UpstreamRef {
		fmt.Printf("Upstream data up to date: branch %s at commit %s\n", target, ref)
		return
	}
	os.WriteFile("upstream.ref", []byte(ref), 0666)

	var file string
	for i := -1; i < len(data.UpstreamLists); i++ {
		if i < 0 {
			file = data.UpstreamConfig
		} else {
			file = data.UpstreamLists[i] + ".txt"
		}
		err = fetch(file, ref)
		if err != nil {
			log.Fatal(err)
		}
	}
}

// Resolve upstream branch/tag name to a commit hash
func commit(head string) (hash string, err error) {
	resp, err := http.Get(refsUrl)
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
	dest := filepath.Join("..", filename)

	fmt.Printf("Fetching %s to %s\n", filename, dest)

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return fmt.Errorf("HTTP error: %s (%s)", resp.Status, url)
	}

	output, err := os.Create(dest)
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
