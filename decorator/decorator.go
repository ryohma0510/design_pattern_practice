package decorator

import "fmt"

// Interface

// DisplayInterface は装飾と中身を同一視するためのInterface
type DisplayInterface interface {
	Columns() int
	Rows() int
	RowText(row int) string
	Show(callerDisplay DisplayInterface) string
}

// DefaultDisplay はDisplayInterfaceのShow()が共通実装として使えるようにするための構造体
type DefaultDisplay struct {
	DisplayInterface
}

// Show は呼び出し元の持っているメソッドを使いたいので引数で受け取る
func (d *DefaultDisplay) Show(callerDisplay DisplayInterface) string {
	var result string
	for i := 0; i < callerDisplay.Rows(); i++ {
		result += callerDisplay.RowText(i)
	}

	return result
}

// 中身の文字列

type StringDisplay struct {
	innerStr string
	*DefaultDisplay
}

func NewStringDisplay(str string) *StringDisplay {
	return &StringDisplay{innerStr: str}
}

func (s *StringDisplay) Columns() int {
	return len(s.innerStr)
}

func (s *StringDisplay) Rows() int {
	return 1
}

func (s *StringDisplay) RowText(row int) string {
	if row == 0 {
		return s.innerStr
	} else {
		return ""
	}
}

// 飾り付け

type Border struct {
	*DefaultDisplay
	innerDisplay DisplayInterface
}

// SideBorder は左右両端に装飾するための構造体
type SideBorder struct {
	borderChar rune
	*Border
}

func NewSideBorder(borderChar rune, display DisplayInterface) *SideBorder {
	return &SideBorder{
		borderChar: borderChar,
		Border: &Border{
			innerDisplay: display,
		},
	}
}

func (b *SideBorder) Columns() int {
	return 1 + b.innerDisplay.Columns() + 1
}

func (b *SideBorder) Rows() int {
	return b.innerDisplay.Rows()
}

func (b *SideBorder) RowText(row int) string {
	return string(b.borderChar) + b.innerDisplay.RowText(row) + string(b.borderChar)
}

// FullBorder は囲みで装飾するための構造体
type FullBorder struct {
	*Border
}

func NewFullBorder(display DisplayInterface) *FullBorder {
	return &FullBorder{
		Border: &Border{
			innerDisplay: display,
		},
	}
}

func (b *FullBorder) Columns() int {
	return 1 + b.innerDisplay.Columns() + 1
}

func (b *FullBorder) Rows() int {
	return 1 + b.innerDisplay.Rows() + 1
}

func (b *FullBorder) RowText(row int) string {
	if row == 0 {
		return fmt.Sprintf("+%s+\n", b.makeLine('-', b.innerDisplay.Columns()))
	} else if row == b.innerDisplay.Rows()+1 {
		return fmt.Sprintf("+%s+\n", b.makeLine('-', b.innerDisplay.Columns()))
	} else {
		return fmt.Sprintf("|%s|\n", b.innerDisplay.RowText(row-1))
	}
}

func (b FullBorder) makeLine(char rune, count int) string {
	var result string

	for i := 0; i < count; i++ {
		result += string(char)
	}

	return result
}
