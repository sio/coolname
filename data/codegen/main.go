package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/sio/coolname/data"
)

const (
	baseUrl = "https://github.com/alexanderlukanin13/coolname"
	dataUrl = baseUrl + "raw/$REF/coolname/data/"
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

	fetch(data.UpstreamConfig, ref)
	for _, list := range data.UpstreamLists {
		fetch(list+".txt", ref)
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
func fetch(filename string, commit string) {
	url := strings.ReplaceAll(dataUrl, "$REF", commit) + filename
	fmt.Printf("Fetching %s from %s\n", filename, url)
}
