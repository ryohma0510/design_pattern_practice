package builder

import "fmt"

type Builder interface {
	makeTitle(title string)
	makeString(str string)
	makeItems(items []string)
	close()
}

type Director struct {
	builder Builder
}

func NewDirector(builder Builder) *Director {
	return &Director{builder: builder}
}

func (d *Director) Construct() {
	d.builder.makeTitle("Greeting")
	d.builder.makeString("朝から昼にかけて")
	d.builder.makeItems([]string{"おはようございます。", "こんにちは。"})
	d.builder.makeString("夜に")
	d.builder.makeItems([]string{"こんばんは。", "おやすみなさい。", "さようなら。"})
	d.builder.close()
}

type TextBuilder struct {
	str string
}

func NewTextBuilder() *TextBuilder {
	return &TextBuilder{}
}

func (t *TextBuilder) makeTitle(title string) {
	t.str += "==\n"
	t.str += fmt.Sprintf("「 %s 」\n", title)
	t.str += "\n"
}

func (t *TextBuilder) makeString(str string) {
	t.str += fmt.Sprintf("■%s\n", str)
	t.str += "\n"
}

func (t *TextBuilder) makeItems(items []string) {
	for _, item := range items {
		t.str += fmt.Sprintf(" *%s\n", item)
	}
	t.str += "\n"
}

func (t *TextBuilder) close() {
	t.str += "==\n"
}

func (t *TextBuilder) getResult() string {
	return t.str
}
