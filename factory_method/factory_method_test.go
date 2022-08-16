package factory_method

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIDCardFactory(t *testing.T) {
	factory := &IDCardFactory{}

	card1 := factory.Create(factory, "John")
	card2 := factory.Create(factory, "Bob")
	card3 := factory.Create(factory, "Ken")

	assert.Equal(t, "Johnのカードを使います", card1.use())
	assert.Equal(t, "Bobのカードを使います", card2.use())
	assert.Equal(t, "Kenのカードを使います", card3.use())
}
