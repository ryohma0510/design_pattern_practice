package bridge

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBridge(t *testing.T) {
	display1 := DefaultDisplay{impl: NewStringDisplayImpl("Hello")}
	display2 := CountDisplay{DefaultDisplay{NewStringDisplayImpl("Good Morning")}}

	assert.Equal(
		t,
		`+-----+
|Hello|
+-----+
`,
		display1.Display(),
	)
	assert.Equal(
		t,
		`+------------+
|Good Morning|
+------------+
`,
		display2.Display(),
	)
	assert.Equal(
		t,
		`+------------+
|Good Morning|
|Good Morning|
|Good Morning|
+------------+
`,
		display2.MultiDisplay(3),
	)
}
