package composite

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestComposite(t *testing.T) {
	rootDir := NewDirectory("root")
	binDir := NewDirectory("bin")
	tmpDir := NewDirectory("tmp")
	usrDir := NewDirectory("usr")

	rootDir.Add(binDir)
	rootDir.Add(tmpDir)
	rootDir.Add(usrDir)

	viFile := NewFile("vi", 10000)
	binDir.Add(viFile)
	latexFile := NewFile("latex", 20000)
	binDir.Add(latexFile)

	assert.Equal(
		t,
		"/root (30000)\n/root/bin (30000)\n/root/bin/vi (10000)\n/root/bin/latex (20000)\n/root/tmp (0)\n/root/usr (0)\n",
		rootDir.PrintList(""),
	)
}
