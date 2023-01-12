package coolname

// TODO: benchmark memory allocations
// TODO: review exported symbols, make all non-essential ones private (hint: go doc)

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/sio/coolname/data"
)

// coolname generator
//
// Zero value is perfectly usable (will load default configuration)
//
// Tunable knobs:
//   - Upon instanciation user may choose the default dictionary and random
//     generator by specifying field values
//   - After instanciation user may provide a custom configuration and word
//     list via Configure() method
type Generator struct {
	// Which dictionary to use by default
	dictionary string

	// Random number generator, default: rand.Intn (not cryptographically secure)
	random func(max int) int

	// A collection of word bags to draw from
	bags  map[string]data.WordBag
	sizes map[string]int
}

// Return a slice of random words (most likely will result in length of 4)
func (g *Generator) Generate() (words []string, err error) {
	return g.GenerateFrom("all")
}

// Return a slice of N random words
//
// Prepositions and articles (of, from, the) are not counted as words,
// so the resulting slice may contain more elements than `count`
//
// Currently only dictionaries for 2, 3 and 4 words are defined upstream
// (see config.json)
func (g *Generator) GenerateN(count int) (words []string, err error) {
	return g.GenerateFrom(fmt.Sprintf("%d", count))
}

const timeout = time.Second

// Return a random word combination from dictionary specified by name
//
// Valid dictionary names are top level keys from config.json
func (g *Generator) GenerateFrom(dictionary string) (words []string, err error) {
	res := make(chan result)
	go func() {
		var r result
		r.words, r.err = g.generate(dictionary)
		res <- r
	}()
	select {
	case r := <-res:
		return r.words, r.err
	case <-time.After(timeout):
		return words, fmt.Errorf("generator timed out")
	}
}

type result struct {
	words []string
	err   error
}

func (g *Generator) generate(dictionary string) (words []string, err error) {
	g.init()

	var dict data.WordBag
	var ok bool
	dict, ok = g.bags[dictionary]
	if !ok {
		return words, fmt.Errorf("dictionary does not exist: %s", dictionary)
	}

	var size int
	size, ok = g.sizes[dictionary]
	if !ok {
		size = dict.Size()
		g.sizes[dictionary] = size
	}

	words = dict.Get(g.random(size))

	// check for repeated words in output
	unique := make(map[string]struct{})
	for _, word := range words {
		_, exists := unique[word]
		if exists {
			return g.generate(dictionary) // try again
		}
		unique[word] = struct{}{}
	}

	// break phrases into words
	for {
		var i int
		var word string
		for i, word = range words {
			if !strings.Contains(word, " ") {
				continue
			}
			words = concat(words[:i], strings.Fields(word), words[i+1:])
			break
		}
		if i == len(words)-1 {
			break
		}
	}

	return words, nil
}

const slugSeparator = "-"

// Generate URL safe slug
func (g *Generator) Slug() (slug string, err error) {
	return g.SlugFrom("all")
}

// Generate URL safe slug with specified amount of meaningful words
// (see GenerateN)
func (g *Generator) SlugN(count int) (slug string, err error) {
	return g.SlugFrom(fmt.Sprintf("%d", count))
}

// Generate URL safe slug from a given dictionary (see GenerateFrom)
func (g *Generator) SlugFrom(dictionary string) (slug string, err error) {
	words, err := g.GenerateFrom(dictionary)
	return strings.Join(words, slugSeparator), err
}

// Provide default values for uninitialized fields
func (g *Generator) init() {
	if g.random == nil {
		rand.Seed(time.Now().UnixNano())
		g.random = rand.Intn
	}
	if g.bags == nil {
		var err error
		err = g.Configure(data.DefaultConfig(), &data.Words)
		if err != nil {
			panic(err) // default configuration must apply without errors
		}
	}
	if g.dictionary == "" {
		g.dictionary = "all"
	}
}

// Load configuration
func (g *Generator) Configure(conf *data.Config, words *data.WordCollection) (err error) {
	g.bags = make(map[string]data.WordBag)
	g.sizes = make(map[string]int)

	var category string
	for category = range *words {
		g.bags[category] = words.Bag(category)
	}

	done := make(map[string]bool)
	for len(done) != len(*conf) {
		var previous int
		previous = len(done)

		var spec data.ListRef
		for category, spec = range *conf {
			if done[category] {
				continue
			}

			switch spec.Kind {
			case data.Const:
				g.bags[category] = constWordBag(spec.Value)

			case data.Nested:
				children, ok := g.bagsByName(spec.Refs)
				if !ok {
					continue // we are not ready to parse higher levels yet
				}
				wrapped := nestedWordBag(children)
				g.bags[category] = &wrapped

			case data.Cartesian:
				children, ok := g.bagsByName(spec.Refs)
				if !ok {
					continue // we are not ready to parse higher levels yet
				}
				wrapped := cartesianWordBag(children)
				g.bags[category] = &wrapped

			default:
				return fmt.Errorf("ListRef.Kind not implemented: %s", spec.Kind)
			}
			done[category] = true
		}

		if len(done) == previous { // nothing added during this iteration
			var keys []string
			keys = make([]string, len(done))
			var i int
			for category = range done {
				keys[i] = category
				i++
			}
			return fmt.Errorf("encountered configuration loop, parsed %d keys: %v", len(done), keys)
		}
	}

	return nil
}

// Return a slice of WordBags for provided names
func (g *Generator) bagsByName(names []string) (result []data.WordBag, ok bool) {
	var name string
	var exists bool
	for _, name = range names {
		_, exists = g.bags[name]
		if !exists {
			return result, false
		}
	}
	result = make([]data.WordBag, len(names))
	for i := 0; i < len(names); i++ {
		result[i] = g.bags[names[i]]
	}
	return result, true
}

// Slice concatenation
// https://stackoverflow.com/a/40678026
func concat[T any](slices ...[]T) []T {
	var total int
	var slice, result []T
	for _, slice = range slices {
		total += len(slice)
	}
	result = make([]T, total)
	var i int
	for _, slice = range slices {
		i += copy(result[i:], slice)
	}
	return result
}
