package proxy

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProxy(t *testing.T) {
	proxy := NewPrinterProxy("Alice")
	assert.Equal(t, "Alice", proxy.PrinterName())

	proxy.SetPrinterName("Bob")
	assert.Equal(t, "Bob", proxy.PrinterName())

	assert.Equal(
		t,
		`Printerのインスタンス(Bob)を生成中 ..... 完了。
=== Bob ===
Hello
`,
		proxy.Print("Hello"),
	)
}
