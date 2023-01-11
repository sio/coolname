package data

import (
	"testing"
)

func TestWords(t *testing.T) {
	for _, category := range UpstreamLists {
		if !Words.Contains(category) {
			t.Errorf("category not found in dataset: %s", category)
		}
		var bag WordBag
		bag = Words.Bag(category)
		if bag.Size() == 0 {
			t.Errorf("empty category: %s", category)
		}
		for i := 0; i < bag.Size(); i++ {
			word, err := Words.Get(category, i)
			if len(word) == 0 {
				t.Errorf("empty word in category %s at position %d", category, i)
			}
			if err != nil {
				t.Errorf("failed to fetch word #%d from category %s: %v", i, category, err)
			}
		}
	}
}
