package data

import (
	"testing"
)

func TestWords(t *testing.T) {
	for _, category := range UpstreamLists {
		if !Words.Contains(category) {
			t.Errorf("category not found in dataset: %s", category)
		}
		if len(Words[category]) == 0 {
			t.Errorf("empty category: %s", category)
		}
		for i := 0; i < len(Words[category]); i++ {
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
