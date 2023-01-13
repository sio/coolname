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
