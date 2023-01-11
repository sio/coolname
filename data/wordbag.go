package data

type WordBag interface {
	Size() int
	Get(position int) []string
}

// The most straightforward WordBag implementation
type WordList []string

func (wl *WordList) Size() int {
	return len(*wl)
}
func (wl *WordList) Get(position int) []string {
	return []string{(*wl)[position]}
}
