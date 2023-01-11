package tests

import (
	"github.com/sio/coolname"
	"testing"

	"fmt"
	"strings"
)

func TestGenerate(t *testing.T) {
	var g coolname.Generator
	words, err := g.Generate()
	if err != nil {
		t.Fatal(err)
	}
	if len(words) < 4 {
		t.Errorf("output too short: %v", words)
	}
	for i := 0; i < len(words); i++ {
		if len(words[i]) == 0 {
			t.Errorf("word %d is empty: %v", i, words)
		}
		if strings.Contains(words[i], " ") {
			t.Errorf("word %d contains a space: %q", i, words[i])
		}
		if strings.Contains(words[i], "=") {
			t.Errorf("word %d contains an equal sign: %q", i, words[i])
		}
	}
}

func TestMultiple(t *testing.T) {
	var g coolname.Generator
	for i := 0; i < 10; i++ {
		words, err := g.Generate()
		if err != nil {
			t.Fatal(err)
		}
		fmt.Println(words)
	}
}
