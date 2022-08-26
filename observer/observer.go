package observer

import (
	"fmt"
	"math/rand"
)

type IObserver interface {
	Update(generator INumberGenerator) string
}

type INumberGenerator interface {
	AddObserver(observer IObserver)
	DeleteObserver(observer IObserver)
	NotifyObservers(callerGenerator INumberGenerator) string
	Number() int
	Execute() string
}

// NumberGenerator は共通メソッドを実装する
type NumberGenerator struct {
	observers []IObserver
}

func (g *NumberGenerator) AddObserver(observer IObserver) {
	g.observers = append(g.observers, observer)
}

func (g *NumberGenerator) DeleteObserver(observer IObserver) {
	for i, iObserver := range g.observers {
		if observer == iObserver {
			g.observers = append(g.observers[:i], g.observers[i+1:]...)
			break
		}
	}
}

func (g *NumberGenerator) NotifyObservers(callerGenerator INumberGenerator) string {
	var result string

	for _, observer := range g.observers {
		result += observer.Update(callerGenerator)
	}

	return result
}

// RandomNumberGenerator はINumberGeneratorを実装する
type RandomNumberGenerator struct {
	number int
	*NumberGenerator
}

func NewRandomNumberGenerator() *RandomNumberGenerator {
	return &RandomNumberGenerator{
		NumberGenerator: &NumberGenerator{
			observers: nil,
		},
	}
}

func (r *RandomNumberGenerator) Number() int {
	return r.number
}

func (r *RandomNumberGenerator) Execute() string {
	var result string

	for i := 0; i < 20; i++ {
		r.number = rand.Intn(50)
		result += r.NotifyObservers(r)
	}

	return result
}

type DigitObserver struct{}

func NewDigitObserver() *DigitObserver {
	return &DigitObserver{}
}

func (d DigitObserver) Update(generator INumberGenerator) string {
	return fmt.Sprintf("DigitObserver: %d\n", generator.Number())
}
