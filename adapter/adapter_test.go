package adapter

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAdapter(t *testing.T) {
	var printer Print = PrintBanner{Banner{content: "ABCDE"}}
	assert.Equal(t, "(ABCDE)", printer.PrintWeak())
	assert.Equal(t, "*ABCDE*", printer.PrintStrong())
}
