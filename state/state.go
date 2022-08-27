package state

import "fmt"

// State は状態に依存した振る舞いをするメソッドの集まり
type State interface {
	DoClock(ctx Context, hour int) string
	DoUse(ctx Context) string
	DoAlarm(ctx Context) string
	DoPhone(ctx Context) string
	fmt.Stringer
}

// singleton
var dayStateInstance *dayState

type dayState struct{}

func DayStateInstance() *dayState {
	if dayStateInstance == nil {
		dayStateInstance = &dayState{}
	}

	return dayStateInstance
}

func (d dayState) DoClock(ctx Context, hour int) string {
	if hour < 9 || hour >= 17 {
		return ctx.ChangeState(NightStateInstance())
	} else {
		return ""
	}
}

func (d dayState) DoUse(ctx Context) string {
	return ctx.RecordLog("金庫使用(昼間)")
}

func (d dayState) DoAlarm(ctx Context) string {
	return ctx.CallSecurityCenter("非常ベル(昼間)")
}

func (d dayState) DoPhone(ctx Context) string {
	return ctx.CallSecurityCenter("通常の通話(昼間)")
}

func (d dayState) String() string {
	return "[昼間]"
}

// singleton
var nightStateInstance *nightState

type nightState struct{}

func NightStateInstance() *nightState {
	if nightStateInstance == nil {
		nightStateInstance = &nightState{}
	}

	return nightStateInstance
}

func (n nightState) DoClock(ctx Context, hour int) string {
	if hour >= 9 && hour < 17 {
		return ctx.ChangeState(DayStateInstance())
	} else {
		return ""
	}

}

func (n nightState) DoUse(ctx Context) string {
	return ctx.CallSecurityCenter("非常:夜間の金庫使用")
}

func (n nightState) DoAlarm(ctx Context) string {
	return ctx.CallSecurityCenter("非常ベル(夜間)")
}

func (n nightState) DoPhone(ctx Context) string {
	return ctx.RecordLog("夜間の通話録音")
}

func (n nightState) String() string {
	return "[夜間]"
}

// Context は現在の状態を保持し、Stateパターンの利用に必要なインターフェースを定める
type Context interface {
	SetClock(hour int) string
	ChangeState(state State) string
	CallSecurityCenter(log string) string
	RecordLog(log string) string
}

type defaultContext struct {
	currentState State
}

type SampleContext struct {
	*defaultContext
}

func NewSampleContext(currentState State) *SampleContext {
	return &SampleContext{
		defaultContext: &defaultContext{
			currentState: currentState,
		},
	}
}

func (s *SampleContext) SetClock(hour int) string {
	resultString := fmt.Sprintf("現在時刻は%02d:00\n", hour)

	// 現在の状態に依存した処理として扱うことができている
	resultString += s.currentState.DoClock(s, hour)
	return resultString
}

func (s *SampleContext) ChangeState(state State) string {
	resultStr := fmt.Sprintf("%sから%sへ状態が変化しました\n", s.currentState, state)
	s.currentState = state

	return resultStr
}

func (s *SampleContext) CallSecurityCenter(log string) string {
	return fmt.Sprintf("call! %s\n", log)
}

func (s *SampleContext) RecordLog(log string) string {
	return fmt.Sprintf("record ... %s\n", log)
}
