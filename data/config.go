package data

import (
	_ "embed"
	"encoding/json"
	"fmt"
)

type WordBag interface {
	Size() int
	Get(position int) string
}

type WordCollection map[string]WordList

func (c *WordCollection) Contains(category string) bool {
	_, exists := (*c)[category]
	return exists
}

func (c *WordCollection) Get(category string, position int) (value string, err error) {
	_, exists := (*c)[category]
	if !exists {
		return value, fmt.Errorf("category not found: %s", category)
	}
	if position >= len((*c)[category]) || position < 0 {
		return value, fmt.Errorf("invalid word index: %d (category %s contains %d entries)", position, category, len((*c)[category]))
	}
	return (*c)[category][position], nil
}

// The most straightforward WordBag implementation
type WordList []string

func (wl *WordList) Size() int {
	return len(*wl)
}
func (wl *WordList) Get(position int) string {
	return (*wl)[position]
}

type Config map[string]ListRef

type ListRef struct {
	Comment      string   `json:"comment"`
	Kind         RefType  `json:"type"`
	Refs         []string `json:"lists"`
	Unique       bool     `json:"ensure_unique"`
	UniquePrefix int      `json:"ensure_unique_prefix"`
	Value        string   `json:"value"`

	/* The following config.json fields are supported by upstream    // TODO
	   but are not implemented here:
	        WORDS = 'words'
	        PHRASES = 'phrases'
	        NUMBER_OF_WORDS = 'number_of_words'
	        GENERATOR = 'generator'
	        MAX_LENGTH = 'max_length'
	        MAX_SLUG_LENGTH = 'max_slug_length'
	*/
}

type RefType string

const (
	Nested    RefType = "nested"
	Cartesian         = "cartesian"
	List              = "words"
	Phrases           = "phrases"
	Const             = "const"
)

//go:embed config.json
var rawDefaultConfig []byte

func DefaultConfig() (c *Config) {
	c = &Config{}
	c.Parse(rawDefaultConfig)
	return c
}

// Parse json configuration
func (c *Config) Parse(input []byte) {
	err := json.Unmarshal(input, c)
	if err != nil {
		panic(err)
	}
}
