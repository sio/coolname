package data

import (
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
	Comment string
	Type refType
	Refs []string
	Unique bool
	UniquePrefix bool
}

type refType string

const (
	Nested refType = "nested"
	Cartesian = "cartesian"
	Words = "words"
	Phrases = "phrases"
	Const = "const"
)
