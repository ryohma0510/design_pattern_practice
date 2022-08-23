package visitor

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestVisitor(t *testing.T) {
	rootDir := NewDirectory("root")
	binDir := NewDirectory("bin")
	tmpDir := NewDirectory("tmp")
	usrDir := NewDirectory("usr")

	rootDir.Add(binDir)
	rootDir.Add(tmpDir)
	rootDir.Add(usrDir)

	binDir.Add(NewFile("vi", 10000))
	binDir.Add(NewFile("latex", 20000))
	assert.Equal(
		t,
		`/root (30000)
/root/bin (30000)
/root/bin/vi (10000)
/root/bin/latex (20000)
/root/tmp (0)
/root/usr (0)
`,
		rootDir.Accept(&ListVisitor{}),
	)
}
