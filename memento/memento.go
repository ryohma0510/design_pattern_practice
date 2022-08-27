package memento

import (
	"fmt"
	"math/rand"
	"strings"
)

// Memento は記録用の構造体
// Originator役の内部情報を持っているが、その情報を2つのインターフェースを使って公開範囲を狭めている
// しかしGolangにはJavaのようにクラス間でのアクセス制限ができないので、Javaの手本通りには実装できない
type Memento struct {
	money  int
	fruits []string
}

func NewMemento(money int) *Memento {
	return &Memento{money: money}
}

func (m *Memento) Money() int {
	return m.money
}

func (m *Memento) Fruits() []string {
	return m.fruits
}

// setterは外部に公開せずOriginator役からのみ触れるようにする
func (m *Memento) addFruit(fruit string) {
	m.fruits = append(m.fruits, fruit)
}

// Gamer はOriginator役の構造体
// 前の状態のMemento役を渡されるとその状態に戻る処理を行なう
type Gamer struct {
	money  int
	fruits []string
}

func (g *Gamer) Money() int {
	return g.money
}

func NewGamer(money int) *Gamer {
	return &Gamer{money: money}
}

var fruitsNames = []string{"りんご", "ぶどう", "ばなな", "みかん"}

const DeliciousPrefix = "美味しい"

func (g *Gamer) Bet() string {
	dice := rand.Intn(6) + 1
	switch dice {
	case 1:
		g.money += 100
		return "所持金が増えました"
	case 2:
		g.money /= 2
		return "所持金が半分になりました"
	case 6:
		fruit := g.randomFruit()
		g.fruits = append(g.fruits, fruit)
		return fmt.Sprintf("フルーツ(%s)をもらいました", fruit)
	default:
		return "何も起こりませんでした"
	}
}

func (g Gamer) randomFruit() string {
	prefix := ""
	if rand.Intn(2) == 1 {
		prefix += DeliciousPrefix
	}

	return prefix + fruitsNames[rand.Intn(len(fruitsNames))]
}

// CreateMemento はスナップショットをとる
// 状態を一時保存したくなったらこれを使うだけでいい
func (g *Gamer) CreateMemento() *Memento {
	memento := NewMemento(g.money)

	for _, fruit := range g.fruits {
		if strings.Contains(fruit, DeliciousPrefix) {
			// 美味しいフルーツだけ保存
			memento.addFruit(fruit)
		}
	}

	return memento
}

// RestoreMemento は渡されたMementoの状態に戻る
func (g *Gamer) RestoreMemento(memento *Memento) {
	g.money = memento.Money()
	g.fruits = memento.Fruits()
}

func (g *Gamer) String() string {
	return fmt.Sprintf("[money = %d, fruits = %v]", g.money, g.fruits)
}
