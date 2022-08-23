package visitor

import "fmt"

// DirEntryInterface はディレクトリとファイルを同一視するためのInterface
type DirEntryInterface interface {
	Name() string
	Size() int
	Add(dirEntry DirEntryInterface)
	String() string
	acceptInterface
}

type VisitInterface interface {
	Visit(entry DirEntryInterface) string
}

type acceptInterface interface {
	Accept(visitor VisitInterface) string
}

// File はAcceptInterfaceを実装している
type File struct {
	DirEntryInterface
	name string
	size int
}

func NewFile(name string, size int) *File {
	return &File{name: name, size: size}
}

func (f File) Name() string {
	return f.name
}

func (f File) Size() int {
	return f.size
}

func (f File) String() string {
	return fmt.Sprintf("%s (%d)", f.Name(), f.Size())
}

func (f File) Accept(visitor VisitInterface) string {
	return visitor.Visit(f)
}

// Directory はAcceptInterfaceを実装している
type Directory struct {
	DirEntryInterface
	name    string
	entries []DirEntryInterface
}

func (d *Directory) Entries() []DirEntryInterface {
	return d.entries
}

func NewDirectory(name string) *Directory {
	return &Directory{name: name}
}

func (d *Directory) Name() string {
	return d.name
}

func (d *Directory) Size() int {
	var size int

	for _, entry := range d.entries {
		size += entry.Size()
	}

	return size
}

func (d *Directory) Add(dirEntry DirEntryInterface) {
	d.entries = append(d.entries, dirEntry)
}

func (d *Directory) String() string {
	return fmt.Sprintf("%s (%d)", d.Name(), d.Size())
}

func (d *Directory) Accept(visitor VisitInterface) string {
	return visitor.Visit(d)
}

type ListVisitor struct {
	currentDir string
}

func (l *ListVisitor) Visit(entry DirEntryInterface) string {
	// これで分けるの微妙な気がするが、一度書籍のJavaと同じ雰囲気で実装してみる
	result := fmt.Sprintf("%s/%s", l.currentDir, entry)

	// ディレクトリの場合は下層を再起的にAcceptする
	if dir, ok := entry.(*Directory); ok {
		saveDir := l.currentDir
		l.currentDir = fmt.Sprintf("%s/%s", l.currentDir, dir.Name())
		for _, innerEntry := range dir.Entries() {
			innerEntry.Accept(l)
		}
		l.currentDir = saveDir
	}

	return result
}
