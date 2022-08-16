package template_method

import "fmt"

type Displayer interface {
	open() string
	print() string
	close() string
}

type AbstractDisplay struct{}

func (d AbstractDisplay) Display(displayer Displayer) string {
	result := displayer.open()
	for i := 0; i < 5; i++ {
		result += displayer.print()
	}
	result += displayer.close()

	return result
}

type CharDisplay struct {
	AbstractDisplay
	content rune
}

func (CharDisplay) open() string {
	return "<<"
}

func (d CharDisplay) print() string {
	return string(d.content)
}

func (CharDisplay) close() string {
	return ">>"
}

type StringDisplay struct {
	AbstractDisplay
	content string
}

func (StringDisplay) open() string {
	return "+----+\n"
}

func (s StringDisplay) print() string {
	return fmt.Sprintf("|%s|\n", s.content)
}

func (StringDisplay) close() string {
	return "+----+\n"
}
