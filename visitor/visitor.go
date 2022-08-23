package visitor

import "fmt"

// DirEntryInterface はディレクトリとファイルを同一視するためのInterface
type DirEntryInterface interface {
	Name() string
	Size() int
	String() string
	Accept(visitor VisitInterface) string
}

type VisitInterface interface {
	VisitFile(file *File) string
	VisitDirectory(directory *Directory) string
}

// File はAcceptInterfaceを実装している
type File struct {
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

// Accept には具体的な処理を書かずに、Visitorに具体的な処理を記述することで、データ構造と処理を分離することができる
func (f File) Accept(visitor VisitInterface) string {
	return visitor.VisitFile(&f)
}

// Directory はAcceptInterfaceを実装している
type Directory struct {
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
	return visitor.VisitDirectory(d)
}

// ListVisitor は具体的な処理を実行する
type ListVisitor struct {
	currentDir string
}

func (l *ListVisitor) VisitFile(file *File) string {
	return fmt.Sprintf("%s/%s\n", l.currentDir, file)
}

func (l *ListVisitor) VisitDirectory(directory *Directory) string {
	result := fmt.Sprintf("%s/%s\n", l.currentDir, directory)

	saveDir := l.currentDir
	l.currentDir = fmt.Sprintf("%s/%s", l.currentDir, directory.Name())
	for _, innerEntry := range directory.Entries() {
		result += innerEntry.Accept(l)
	}
	l.currentDir = saveDir

	return result
}
