package facade

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFacade(t *testing.T) {
	pageMaker := PageMaker{}
	assert.Equal(
		t,
		`# Welcome to a's !
This is content
`,
		pageMaker.MakeWelcomePage("a@a.com"),
	)
}
