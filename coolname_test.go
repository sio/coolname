package coolname

import (
	"testing"
)

func TestEntropy(t *testing.T) {
	tests := []struct {
		Dictionary string
		Entropy    float64
	}{
		{"all", 6.26e10},
		{"4", 6.23e10},
		{"3", 2.8e8},
		{"2", 3.2e5},
	}
	g := defaultGenerator
	g.init()
	for _, tt := range tests {
		bag := g.bags[tt.Dictionary]
		got := float64(bag.Size())
		want := tt.Entropy
		t.Logf("entropy for dictionary %q: want %.1e, got %.1e (%v)", tt.Dictionary, want, got, got)
		if got < want {
			t.Errorf("entropy too low for dictionary %q", tt.Dictionary)
		}
	}
}

func TestUniqueness(t *testing.T) {
	// It's only feasible to test uniqueness for the smallest dictionary
	// ("2" has 320k permutations)
	const dict = "2"

	// Some adjectives occur in multiple lists, like "remarkable" in
	// "of_modifier" and "adjective".
	// Helpful command to find all other such words:
	//    sort data/data.go | uniq -c | grep -vP '^\s*1\s'
	const threshold = 0.99

	if testing.Short() {
		t.Skip("skipping test in short mode")
	}

	var seq incrementer
	g := Generator{random: seq.FakeRandInt}
	g.init()
	size := g.bags[dict].Size()
	if size == 0 {
		t.Fatalf("empty dictionary: %q", dict)
	}
	unique := make(map[string]struct{})
	var repeats int
	const repeatsShow = 10
	for seq.Value() < size {
		slug, err := g.SlugFrom(dict)
		if err != nil {
			t.Fatalf("generator failure on iteration %d/%d: %v", seq.Value(), size, err)
		}
		_, exists := unique[slug]
		if exists {
			if repeats < repeatsShow {
				t.Logf("slug repeated %q, iteration %d/%d", slug, seq.Value(), size)
			}
			if repeats == repeatsShow {
				t.Logf("further repeated entries are omitted from output")
			}
			repeats++
		}
		unique[slug] = struct{}{}
	}
	t.Logf("unique entries: %d/%d", len(unique), size)
	ratio := float32(len(unique)) / float32(size)
	if ratio < threshold {
		t.Errorf("uniqueness ratio lower than threshold: %.2f < %.2f", ratio, threshold)
	}
}

type incrementer int

func (i *incrementer) FakeRandInt(max int) int {
	return i.Next()
}
func (i *incrementer) Next() int {
	result := i.Value()
	*i++
	return result
}
func (i *incrementer) Value() int {
	return int(*i)
}
