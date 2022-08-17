package prototype

import (
	"fmt"
)

type Product interface {
	use(s string) string
	createClone() Product
}

type Manager struct {
	showcase map[string]Product
}

func NewManager() Manager {
	return Manager{showcase: map[string]Product{}}
}

func (m *Manager) register(protoName string, proto Product) {
	m.showcase[protoName] = proto
}

func (m *Manager) create(protoName string) (Product, error) {
	p, ok := m.showcase[protoName]
	if !ok {
		return nil, fmt.Errorf("missing key %s", protoName)
	}

	return p.createClone(), nil
}

type MessageBox struct {
	decoChar rune
}

func NewMessageBox(decoChar rune) MessageBox {
	return MessageBox{
		decoChar: decoChar,
	}
}

func (m MessageBox) use(s string) string {
	sLen := len(s)

	var decoCharAtTopAndBottom string
	for i := 0; i < sLen+4; i++ {
		decoCharAtTopAndBottom += string(m.decoChar)
	}

	result := fmt.Sprintf("%s\n", decoCharAtTopAndBottom)
	result += fmt.Sprintf("%s %s %s\n", string(m.decoChar), s, string(m.decoChar))
	result += fmt.Sprintf("%s\n", decoCharAtTopAndBottom)

	return result
}

func (m MessageBox) createClone() Product {
	return NewMessageBox(m.decoChar)
}
