package decorator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDecorator(t *testing.T) {
	display1 := NewStringDisplay("Hello")
	assert.Equal(t, "Hello", display1.Show(display1))

	display2 := NewSideBorder('#', display1)
	assert.Equal(t, "#Hello#", display2.Show(display2))

	display3 := NewFullBorder(display2)
	assert.Equal(
		t,
		`+-------+
|#Hello#|
+-------+
`,
		display3.Show(display3),
	)

}
