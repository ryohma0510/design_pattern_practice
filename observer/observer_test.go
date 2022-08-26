package observer

import (
	"fmt"
	"testing"
)

func TestObserver(t *testing.T) {
	generator := NewRandomNumberGenerator()
	observer := NewDigitObserver()

	generator.AddObserver(observer)
	fmt.Println(generator.Execute())
}
