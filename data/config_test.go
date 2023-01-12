package data

import (
	"testing"

	"encoding/json"
)

func TestDefaultConfig(t *testing.T) {
	var conf *Config
	conf = DefaultConfig()
	if len(*conf) < 27 {
		t.Errorf("default config too short (%d keys): %v", len(*conf), conf)
	}
}

func TestListRefParsing(t *testing.T) {
	input := []byte(`
		{
			"comment": "Entry point",
			"type": "nested",
			"lists": ["2", "3", "4"],
			"ensure_unique": true,
			"ensure_unique_prefix": 4,
			"max_slug_length": 50
		}
	`)
	var ref ListRef
	err := json.Unmarshal(input, &ref)
	if err != nil {
		t.Fatalf("failed to parse a ListRef: %v", err)
	}
	if ref.Kind != Nested {
		t.Errorf("invalid reference type: %v", ref)
	}
	if len(ref.Refs) != 3 {
		t.Errorf("invalid refs: %v", ref)
	}
}
