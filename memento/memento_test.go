package memento

import (
	"fmt"
	"testing"
)

func TestMemento(t *testing.T) {
	gamer := NewGamer(100)
	memento := gamer.CreateMemento()

	for i := 0; i < 10; i++ {
		fmt.Printf("==== %d\n", i)
		fmt.Printf("現状: %s\n", gamer)

		gamer.Bet()

		fmt.Printf("所持金は%d円になりました\n", gamer.Money())

		if gamer.Money() > memento.Money() {
			fmt.Println("だいぶ増えたので現在の状態を保存しておこう")
			memento = gamer.CreateMemento()
		} else if gamer.Money() < memento.Money()/2 {
			fmt.Println("だいぶ減ったので以前の状態に復帰しよう")
			gamer.RestoreMemento(memento)
		}
	}
}
