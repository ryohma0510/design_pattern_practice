package singleton

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSingleton(t *testing.T) {
	i1 := GetInstance()
	i2 := GetInstance()

	// pointerで比較するので等値性を検証している
	assert.Equal(t, i1, i2)
}
