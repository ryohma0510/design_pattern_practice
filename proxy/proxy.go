package proxy

import "fmt"

type Printable interface {
	SetPrinterName(name string)
	PrinterName() string
	Print(str string) string
}

// Printer は重いやつ
type Printer struct {
	name string
}

func (p *Printer) PrinterName() string {
	return p.name
}

func (p *Printer) SetPrinterName(name string) {
	p.name = name
}

// NewPrinter はインスタンス生成が重たい想定
func NewPrinter(name string) (*Printer, string) {
	p := &Printer{name: name}

	return p, p.heavyJob(fmt.Sprintf("Printerのインスタンス(%s)を生成中", name))
}

func (p *Printer) heavyJob(msg string) string {
	return fmt.Sprintf("%s ..... 完了。\n", msg)
}

func (p *Printer) Print(str string) string {
	return fmt.Sprintf("=== %s ===\n%s\n", p.name, str)
}

// PrinterProxy はProxy役
// 必要になった時に初めてPrinterインスタンスを生成する
type PrinterProxy struct {
	name string
	real *Printer
}

func NewPrinterProxy(name string) *PrinterProxy {
	return &PrinterProxy{name: name}
}

func (p *PrinterProxy) SetPrinterName(name string) {
	if p.real != nil {
		p.real.SetPrinterName(name)
	}

	p.name = name
}

func (p *PrinterProxy) PrinterName() string {
	return p.name
}

func (p *PrinterProxy) Print(str string) string {
	var resultStr string

	resultStr += p.realize()
	resultStr += p.real.Print(str)
	return resultStr
}

func (p *PrinterProxy) realize() string {
	if p.real == nil {
		newPrinter, str := NewPrinter(p.name)
		p.real = newPrinter

		return str
	}

	return ""
}
