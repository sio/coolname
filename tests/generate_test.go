package tests

import (
	"github.com/sio/coolname"
	"testing"

	"strings"
)

func TestGenerate(t *testing.T) {
	var g coolname.Generator
	words := g.Generate()
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
