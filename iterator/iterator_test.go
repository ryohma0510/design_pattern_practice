package iterator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIterator(t *testing.T) {

	bs := BookShelf{
		Books: []*Book{},
		Last:  0,
	}

	books := []string{"A", "B", "C"}

	bs.AppendBook(&Book{Name: books[0]})
	bs.AppendBook(&Book{Name: books[1]})
	bs.AppendBook(&Book{Name: books[2]})

	it := bs.Iterator()

	for i := 0; it.HasNext(); i++ {
		assert.Equal(t, books[i], it.Next().(*Book).Name)
	}
}
