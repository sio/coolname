package data

import (
	"fmt"
)

type WordCollection map[string]*WordList

func (c *WordCollection) Contains(category string) bool {
	_, exists := (*c)[category]
	return exists
}

func (c *WordCollection) Bag(category string) WordBag {
	bag := &wordCollectionCategoryBag{
		category:   category,
		collection: c,
	}
	return bag
}

func (c *WordCollection) Get(category string, position int) (value string, err error) {
	var list *WordList
	list, exists := (*c)[category]
	if !exists {
		return value, fmt.Errorf("category not found: %s", category)
	}
	if position >= len(*list) || position < 0 {
		return value, fmt.Errorf("invalid word index: %d (category %s contains %d entries)", position, category, len(*list))
	}
	return (*list)[position], nil
}

// Helpful WordBag wrapper for WordCollection
type wordCollectionCategoryBag struct {
	category   string
	collection *WordCollection
}

func (b *wordCollectionCategoryBag) Get(position int) []string {
	value, _ := b.collection.Get(b.category, position)
	return []string{value}
}
func (b *wordCollectionCategoryBag) Size() int {
	var list *WordList
	list = (*b.collection)[b.category]
	return len(*list)
}
