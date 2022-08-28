package command

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCommand(t *testing.T) {
	macroCommand := NewMacroCommand()
	macroCommand.Add(DrawCommand{Position{
		X: 0,
		Y: 0,
	}})
	macroCommand.Add(DrawCommand{Position{
		X: 1,
		Y: 1,
	}})

	assert.Equal(t, "0.0\n1.1\n", macroCommand.Execute())
}
