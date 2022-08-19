package bridge

import "fmt"

// 機能のクラス階層

type DefaultDisplay struct {
	impl DisplayImpl
}

func NewDisplay(impl DisplayImpl) *DefaultDisplay {
	return &DefaultDisplay{impl: impl}
}

func (d DefaultDisplay) Open() string {
	return d.impl.RawOpen()
}

func (d DefaultDisplay) Print() string {
	return d.impl.RawPrint()
}

func (d DefaultDisplay) Close() string {
	return d.impl.RawClose()
}

func (d DefaultDisplay) Display() string {
	var result string

	result += fmt.Sprintf("%s\n", d.Open())
	result += fmt.Sprintf("%s\n", d.Print())
	result += fmt.Sprintf("%s\n", d.Close())

	return result
}

type CountDisplay struct {
	DefaultDisplay
}

func (d CountDisplay) MultiDisplay(times int) string {
	var result string

	result += fmt.Sprintf("%s\n", d.Open())
	for i := 0; i < times; i++ {
		result += fmt.Sprintf("%s\n", d.Print())
	}
	result += fmt.Sprintf("%s\n", d.Close())

	return result
}

// 実装のクラス階層

type DisplayImpl interface {
	RawOpen() string
	RawPrint() string
	RawClose() string
}

type StringDisplayImpl struct {
	content string
	width   int
}

func NewStringDisplayImpl(content string) StringDisplayImpl {
	return StringDisplayImpl{
		content: content,
		width:   len(content),
	}
}

func (impl StringDisplayImpl) RawOpen() string {
	return impl.printLine()
}

func (impl StringDisplayImpl) RawPrint() string {
	return fmt.Sprintf("|%s|", impl.content)
}

func (impl StringDisplayImpl) RawClose() string {
	return impl.printLine()
}

func (impl StringDisplayImpl) printLine() string {
	var result string
	result += "+"
	for i := 0; i < impl.width; i++ {
		result += "-"
	}
	result += "+"

	return result
}
