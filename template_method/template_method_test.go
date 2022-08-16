package template_method

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCharDisplay(t *testing.T) {
	charDisplay := CharDisplay{
		AbstractDisplay: AbstractDisplay{},
		content:         'A',
	}

	assert.Equal(t, "<<AAAAA>>", charDisplay.Display(charDisplay))
}

func TestStringDisplay(t *testing.T) {
	stringDisplay := StringDisplay{
		AbstractDisplay: AbstractDisplay{},
		content:         "Hello!",
	}

	expected := `+----+
|Hello!|
|Hello!|
|Hello!|
|Hello!|
|Hello!|
+----+
`

	assert.Equal(t, expected, stringDisplay.Display(stringDisplay))
}
