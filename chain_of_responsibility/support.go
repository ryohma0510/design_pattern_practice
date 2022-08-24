package chain_of_responsibility

import "fmt"

type Trouble struct {
	number int
}

func NewTrouble(number int) Trouble {
	return Trouble{number: number}
}

func (t Trouble) Number() int {
	return t.number
}

func (t Trouble) String() string {
	return fmt.Sprintf("[Trouble %d]", t.Number())
}

type SupportInterface interface {
	Support(callerSupport SupportInterface, trouble Trouble) string
	SetNext(support SupportInterface) SupportInterface
	Resolve(trouble Trouble) bool
}

type DefaultSupport struct {
	name string
	next SupportInterface
}

func NewDefaultSupport(name string) *DefaultSupport {
	return &DefaultSupport{name: name}
}

func (s *DefaultSupport) SetNext(next SupportInterface) SupportInterface {
	s.next = next
	return next
}

func (s *DefaultSupport) Support(callerSupport SupportInterface, trouble Trouble) string {
	if callerSupport.Resolve(trouble) {
		return s.done(trouble)
	} else if s.next != nil {
		return s.next.Support(s.next, trouble)
	} else {
		return s.fail(trouble)
	}
}

func (s *DefaultSupport) done(trouble Trouble) string {
	return fmt.Sprintf("%s is resolved by %s.\n", trouble, s.name)
}

func (s *DefaultSupport) fail(trouble Trouble) string {
	return fmt.Sprintf("%s cannot be resolved.\n", trouble)
}

type NoSupport struct {
	DefaultSupport
}

func NewNoSupport(name string) *NoSupport {
	return &NoSupport{DefaultSupport: *NewDefaultSupport(name)}
}

func (s *NoSupport) Resolve(trouble Trouble) bool {
	return false
}

type LimitSupport struct {
	DefaultSupport
	limit int
}

func NewLimitSupport(name string, limit int) *LimitSupport {
	return &LimitSupport{
		DefaultSupport: *NewDefaultSupport(name),
		limit:          limit,
	}
}

func (l *LimitSupport) Resolve(trouble Trouble) bool {
	if trouble.Number() < l.limit {
		return true
	} else {
		return false
	}
}

type OddSupport struct {
	DefaultSupport
}

func NewOddSupport(name string) *OddSupport {
	return &OddSupport{
		DefaultSupport: *NewDefaultSupport(name),
	}
}

func (o *OddSupport) Resolve(trouble Trouble) bool {
	if trouble.Number()%2 == 1 {
		return true
	} else {
		return false
	}
}

type SpecialSupport struct {
	DefaultSupport
	number int
}

func NewSpecialSupport(name string, number int) *SpecialSupport {
	return &SpecialSupport{
		DefaultSupport: *NewDefaultSupport(name),
		number:         number,
	}
}

func (s *SpecialSupport) Resolve(trouble Trouble) bool {
	if trouble.Number() == s.number {
		return true
	} else {
		return false
	}
}
