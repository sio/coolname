package coolname

import (
	"github.com/sio/coolname/data"

	"fmt"
)

//
// This file contains implementations of WordBag interface for non-trivial cases
//

// Single-value word bag
type constWordBag string

func (bag constWordBag) Size() int {
	return 1
}

func (bag constWordBag) Get(position int) []string {
	if position != 0 {
		panic(fmt.Sprintf("index out of range: %d", position))
	}
	return []string{string(bag)}
}

// Concatenation of word bags
type nestedWordBag []data.WordBag

func (bag *nestedWordBag) Size() int {
	var size int
	for _, element := range *bag {
		size += element.Size()
	}
	return size
}

func (bag *nestedWordBag) Get(position int) []string {
	for _, element := range *bag {
		if position < element.Size() {
			return element.Get(position)
		}
		position -= element.Size()
	}
	panic("index out of range")
}

// Cartesian product
type cartesianWordBag []data.WordBag

func (bag *cartesianWordBag) Size() int {
	var size int
	size = 1
	for _, element := range *bag {
		size *= element.Size()
		if size == 0 {
			return size
		}
	}
	return size
}

func (bag *cartesianWordBag) Get(index int) []string {
	var result = make([]string, 0, len(*bag))
	for _, element := range *bag {
		var size int = element.Size()
		result = append(result, element.Get(index%size)...)
		index = index / size
	}
	return result
}
