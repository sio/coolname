package data

import (
	_ "embed"
	"encoding/json"
)

type Config map[string]ListRef

type ListRef struct {
	Kind  RefType  `json:"type"`
	Refs  []string `json:"lists"`
	Value string   `json:"value"`

	/* The following config.json fields are supported by upstream    // TODO
	   but are not implemented here:
	        WORDS = 'words'
	        PHRASES = 'phrases'
	        NUMBER_OF_WORDS = 'number_of_words'
	        GENERATOR = 'generator'
	        MAX_LENGTH = 'max_length'
	        MAX_SLUG_LENGTH = 'max_slug_length'

		These fields were briefly supported but were removed
		because we make no use of them
			Comment      string   `json:"comment"`
			Unique       bool     `json:"ensure_unique"`
			UniquePrefix int      `json:"ensure_unique_prefix"`
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

func (c *Config) Item(name string) (value ListRef, exists bool) {
	value, exists = (*c)[name]
	return value, exists
}
