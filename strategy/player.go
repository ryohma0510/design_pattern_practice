package strategy

import "fmt"

type Player struct {
	name      string
	strategy  Strategy
	winCount  int
	loseCount int
	gameCount int
}

func NewPlayer(name string, strategy Strategy) *Player {
	return &Player{name: name, strategy: strategy}
}

func (p *Player) NextHand() Hand {
	return p.strategy.NextHand()
}

func (p *Player) Win() {
	p.strategy.Study(true)
	p.winCount++
	p.gameCount++
}

func (p *Player) Lose() {
	p.strategy.Study(false)
	p.loseCount++
	p.gameCount++
}

func (p *Player) Even() {
	p.gameCount++
}

func (p *Player) String() string {
	return fmt.Sprintf(
		"[%s:%d games, wincount %d, losecount %d]",
		p.name,
		p.gameCount,
		p.winCount,
		p.loseCount,
	)
}
