package tests

import (
	"github.com/sio/coolname"
	"testing"

	"fmt"
	"regexp"
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

func TestShowOutput(t *testing.T) {
	var g coolname.Generator
	for i := 0; i < 10; i++ {
		slug, err := g.Slug()
		if err != nil {
			t.Fatal(err)
		}
		fmt.Println(slug)
	}
}

func TestSlug(t *testing.T) {
	var g coolname.Generator
	valid := regexp.MustCompile(`[a-z-]+`)
	for i := 0; i < 100; i++ {
		slug, err := g.Slug()
		if err != nil {
			t.Errorf("could not generate slug: %v", err)
			continue
		}
		if len(slug) == 0 {
			t.Errorf("empty slug: %q", slug)
			continue
		}
		if strings.Contains(slug, " ") {
			t.Errorf("slug contains whitespace: %q", slug)
		}
		if !valid.MatchString(slug) {
			t.Errorf("does not match regex %v: %q", valid, slug)
		}
	}
}

func BenchmarkSlug(b *testing.B) {
	var g coolname.Generator
	var err error
	for i := 0; i < b.N; i++ {
		_, err = g.Slug()
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkSlugMulti(b *testing.B) {
	var g coolname.Generator
	var err error
	for i := 0; i < b.N; i++ {
		for j := 0; j < 100; j++ {
			_, err = g.Slug()
			if err != nil {
				b.Error(err)
			}
		}
	}
}
