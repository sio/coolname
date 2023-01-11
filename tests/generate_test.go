package tests

import (
	"github.com/sio/coolname"
	"testing"

	"fmt"
)

func TestGenerate(t *testing.T) {
	var g coolname.Generator
	words := g.Generate()
	fmt.Println(words)
	t.Fatal("TODO")
}
