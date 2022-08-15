package iterator

// Aggregate
// Iteratorを作り出す役を決める
type Aggregate interface {
	Iterator() Iterator
}

// もしBookShelfにループ系の実装がされていたら、Booksの管理の仕方が変更されると数え上げの処理も変更をしなくてはいけなくなる
type BookShelf struct {
	Books []*Book
	Last  int
}

func (bs BookShelf) GetBookAt(idx int) *Book {
	return bs.Books[idx]
}

func (bs *BookShelf) AppendBook(book *Book) {
	bs.Books = append(bs.Books, book)
	bs.Last++
}

func (bs BookShelf) GetLength() int {
	return bs.Last
}

func (bs BookShelf) Iterator() Iterator {
	return &BookShelfIterator{
		BookShelf: bs,
		idx:       0,
	}
}

// Iterator
// 要素を順番にスキャンしていく
type Iterator interface {
	HasNext() bool
	Next() interface{}
}

// Concrete Iterator
// 数える対象とインデックスを覚える
type BookShelfIterator struct {
	BookShelf
	idx int
}

func (i BookShelfIterator) HasNext() bool {
	if i.idx < i.BookShelf.GetLength() {
		return true
	} else {
		return false
	}
}

func (i *BookShelfIterator) Next() interface{} {
	book := i.GetBookAt(i.idx)
	i.idx++

	return book
}

type Book struct {
	Name string
}
