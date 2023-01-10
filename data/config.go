package data

import (
	"fmt"
)

type WordCollection map[string]WordList
type WordList []string

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
