package strategy

import (
	"fmt"
	"testing"
)

func TestStrategy(t *testing.T) {
	alice := NewPlayer("Alice", &WinningStrategy{})

	// 本当は違うStrategyだけど、勉強用なので妥協
	// Strategyを本来ならここで簡単に切り替えることができる
	bob := NewPlayer("Bob", &WinningStrategy{})

	for i := 0; i < 100; i++ {
		aliceHand := alice.NextHand()
		bobHand := bob.NextHand()

		if aliceHand.IsStrongerThan(bobHand) {
			fmt.Printf("Winner: %s\n", alice.name)
			alice.Win()
			bob.Lose()
		} else if aliceHand.IsWeakerThan(bobHand) {
			fmt.Printf("Winner: %s\n", bob.name)
			alice.Lose()
			bob.Win()
		} else {
			fmt.Println("Even..")
		}
	}
}
