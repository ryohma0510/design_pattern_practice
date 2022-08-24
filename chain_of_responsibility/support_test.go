package chain_of_responsibility

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestResponsibility(t *testing.T) {
	aliceSupport := NewNoSupport("Alice")
	bobSupport := NewLimitSupport("Bob", 100)
	charlieSupport := NewSpecialSupport("Charlie", 429)
	dianaSupport := NewLimitSupport("Diana", 200)
	elmoSupport := NewOddSupport("Elmo")
	fredSupport := NewLimitSupport("Fred", 300)

	aliceSupport.
		SetNext(bobSupport).
		SetNext(charlieSupport).
		SetNext(dianaSupport).
		SetNext(elmoSupport).
		SetNext(fredSupport)

	var actual string
	for i := 0; i < 500; i += 33 {
		actual += aliceSupport.Support(aliceSupport, NewTrouble(i))
	}

	assert.Equal(
		t,
		`[Trouble 0] is resolved by Bob.
[Trouble 33] is resolved by Bob.
[Trouble 66] is resolved by Bob.
[Trouble 99] is resolved by Bob.
[Trouble 132] is resolved by Diana.
[Trouble 165] is resolved by Diana.
[Trouble 198] is resolved by Diana.
[Trouble 231] is resolved by Elmo.
[Trouble 264] is resolved by Fred.
[Trouble 297] is resolved by Elmo.
[Trouble 330] cannot be resolved.
[Trouble 363] is resolved by Elmo.
[Trouble 396] cannot be resolved.
[Trouble 429] is resolved by Charlie.
[Trouble 462] cannot be resolved.
[Trouble 495] is resolved by Elmo.
`,
		actual,
	)
}
