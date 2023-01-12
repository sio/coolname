package main

import (
	"testing"

	"github.com/sio/coolname/data"
)

func TestWords(t *testing.T) {
	for _, category := range upstreamLists {
		if !data.Words.Contains(category) {
			t.Errorf("category not found in dataset: %s", category)
		}
		var bag data.WordBag
		bag = data.Words.Bag(category)
		if bag.Size() == 0 {
			t.Errorf("empty category: %s", category)
		}
		for i := 0; i < bag.Size(); i++ {
			word, err := data.Words.Get(category, i)
			if len(word) == 0 {
				t.Errorf("empty word in category %s at position %d", category, i)
			}
			if err != nil {
				t.Errorf("failed to fetch word #%d from category %s: %v", i, category, err)
			}
		}
	}
}
