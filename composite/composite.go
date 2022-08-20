package composite

import (
	"fmt"
	"strconv"
)

type DirectoryEntryInterface interface {
	Name() string
	Size() int
	Add(entry DirectoryEntryInterface)
	PrintList(prefix string) string
}

type DirectoryEntry struct {
	DirectoryEntryInterface
	name string
}

func (e *DirectoryEntry) Name() string {
	return e.name
}

func (e *DirectoryEntry) print(callerEntry DirectoryEntryInterface) string {
	return fmt.Sprintf(
		"%s (%s)\n",
		callerEntry.Name(),
		strconv.Itoa(callerEntry.Size()),
	)
}

// file

type File struct {
	*DirectoryEntry
	size int
}

func NewFile(name string, size int) *File {
	return &File{
		DirectoryEntry: &DirectoryEntry{
			name: name,
		},
		size: size,
	}
}

func (f *File) Size() int {
	return f.size
}

func (f *File) PrintList(prefix string) string {
	return prefix + "/" + f.print(f)
}

// directory

type Directory struct {
	*DirectoryEntry
	directory []DirectoryEntryInterface
}

func NewDirectory(name string) *Directory {
	return &Directory{
		DirectoryEntry: &DirectoryEntry{
			name: name,
		},
	}
}

func (d *Directory) Size() int {
	var size int

	for _, entry := range d.directory {
		size += entry.Size()
	}
	return size
}

func (d *Directory) Add(entry DirectoryEntryInterface) {
	d.directory = append(d.directory, entry)
}

func (d *Directory) PrintList(prefix string) string {
	var result string

	result += fmt.Sprintf("%s/%s", prefix, d.print(d))

	for _, entry := range d.directory {
		result += entry.PrintList(fmt.Sprintf("%s/%s", prefix, d.Name()))
	}

	return result
}
