package coolname

import (
	"math/rand"
	"time"

	"github.com/sio/coolname/data"
)

type Generator struct {
	config *data.Config
	words  *data.WordCollection
	random func(max int) int
}

// Return a slice of random words (default length is 4)
func (g *Generator) Generate() []string {
	return g.GenerateN(4)
}

// Return a slice of N random words
//
// Prepositions and articles (of, from, the) are not counted as words,
// so the resulting slice may contain more elements than `count`
func (g *Generator) GenerateN(count int) []string {
	g.init()
	result := make([]string, 0, count)
	return result
}

func (g *Generator) init() {
	if g.config == nil {
		g.config = data.DefaultConfig()
	}
	if g.words == nil {
		g.words = &data.Words
	}
	if g.random == nil {
		rand.Seed(time.Now().UnixNano())
		g.random = rand.Intn
	}
}
