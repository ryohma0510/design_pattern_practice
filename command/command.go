package command

import "fmt"

type command interface {
	Execute() string
}

type MacroCommand struct {
	commands []command
}

func NewMacroCommand() *MacroCommand {
	return &MacroCommand{}
}

func (mc *MacroCommand) Add(newCommand command) {
	mc.commands = append(mc.commands, newCommand)
}

func (mc *MacroCommand) Execute() string {
	var resultStr string

	for _, c := range mc.commands {
		resultStr += c.Execute()
	}

	return resultStr
}

type Position struct {
	X, Y int
}

type DrawCommand struct {
	Position
}

func (d DrawCommand) Execute() string {
	return fmt.Sprintf("%d.%d\n", d.Y, d.Y)
}
