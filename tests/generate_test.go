package tests

import (
	"github.com/sio/coolname"
	"testing"

	"fmt"
	"regexp"
	"strings"
)

// TODO: use deterministic RandInt with a small custom WordCollection & Config to verify Generator

func TestGenerate(t *testing.T) {
	var g coolname.Generator
	tests := []struct {
		Func   func() ([]string, error)
		MinLen int
	}{
		{Func: coolname.Generate, MinLen: 2},
		{Func: g.Generate, MinLen: 2},
		{Func: func() ([]string, error) { return g.GenerateN(2) }, MinLen: 2},
		{Func: func() ([]string, error) { return g.GenerateN(3) }, MinLen: 3},
		{Func: func() ([]string, error) { return g.GenerateN(4) }, MinLen: 4},
	}
	for index, test := range tests {
		prefix := fmt.Sprintf("[%d]", index)
		words, err := test.Func()
		if err != nil {
			t.Fatalf("%s %v", prefix, err)
		}
		if len(words) < test.MinLen {
			t.Errorf("%s output shorter than %d elements: %v", prefix, test.MinLen, words)
		}
		for i := 0; i < len(words); i++ {
			if len(words[i]) == 0 {
				t.Errorf("%s word %d is empty: %v", prefix, i, words)
			}
			if strings.Contains(words[i], " ") {
				t.Errorf("%s word %d contains a space: %q", prefix, i, words[i])
			}
			if strings.Contains(words[i], "=") {
				t.Errorf("%s word %d contains an equal sign: %q", prefix, i, words[i])
			}
		}
	}
}

func TestShowOutput(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode (useful only for interactive runs)")
	}
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
