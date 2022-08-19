package strategy

import "math/rand"

type HandType int

const (
	Guu HandType = iota + 1
	Choki
	Paa
)

type ResultType int

const (
	Even ResultType = iota + 1
	Win
	Lose
)

type Hand struct {
	Value HandType
}

func (h Hand) fight(opponent Hand) ResultType {
	if h.Value == opponent.Value {
		return Even
	} else if (h.Value == Guu) && (opponent.Value == Choki) {
		return Win
	} else if (h.Value == Choki) && (opponent.Value == Paa) {
		return Win
	} else if (h.Value == Paa) && (opponent.Value == Guu) {
		return Win
	} else {
		return Lose
	}
}

func (h Hand) IsStrongerThan(opponent Hand) bool {
	return h.fight(opponent) == Win
}

func (h Hand) IsWeakerThan(opponent Hand) bool {
	return h.fight(opponent) == Lose
}

type Strategy interface {
	NextHand() Hand
	Study(win bool)
}

type WinningStrategy struct {
	won      bool
	prevHand Hand
}

func (s *WinningStrategy) NextHand() Hand {
	if s.won {
		return s.prevHand
	} else {
		randValue := rand.Intn(3-1) + 1
		return Hand{Value: HandType(randValue)}
	}
}

func (s *WinningStrategy) Study(win bool) {
	s.won = win
}
