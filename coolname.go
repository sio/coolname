package coolname

import (
	"io"

	"github.com/sio/coolname/data"
)

type Generator struct {
	config *data.Config
	words  *data.WordCollection
	ranInt func(max int) int
}
