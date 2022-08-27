package state

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestState(t *testing.T) {
	ctx := NewSampleContext(DayStateInstance())

	var actual string
	for hour := 0; hour < 24; hour++ {
		actual += ctx.SetClock(hour)
		actual += ctx.currentState.DoUse(ctx)
		actual += ctx.currentState.DoAlarm(ctx)
		actual += ctx.currentState.DoPhone(ctx)
	}

	expected := `現在時刻は00:00
[昼間]から[夜間]へ状態が変化しました
call! 非常:夜間の金庫使用
call! 非常ベル(夜間)
record ... 夜間の通話録音
現在時刻は01:00
call! 非常:夜間の金庫使用
call! 非常ベル(夜間)
record ... 夜間の通話録音
現在時刻は02:00
call! 非常:夜間の金庫使用
call! 非常ベル(夜間)
record ... 夜間の通話録音
現在時刻は03:00
call! 非常:夜間の金庫使用
call! 非常ベル(夜間)
record ... 夜間の通話録音
現在時刻は04:00
call! 非常:夜間の金庫使用
call! 非常ベル(夜間)
record ... 夜間の通話録音
現在時刻は05:00
call! 非常:夜間の金庫使用
call! 非常ベル(夜間)
record ... 夜間の通話録音
現在時刻は06:00
call! 非常:夜間の金庫使用
call! 非常ベル(夜間)
record ... 夜間の通話録音
現在時刻は07:00
call! 非常:夜間の金庫使用
call! 非常ベル(夜間)
record ... 夜間の通話録音
現在時刻は08:00
call! 非常:夜間の金庫使用
call! 非常ベル(夜間)
record ... 夜間の通話録音
現在時刻は09:00
[夜間]から[昼間]へ状態が変化しました
record ... 金庫使用(昼間)
call! 非常ベル(昼間)
call! 通常の通話(昼間)
現在時刻は10:00
record ... 金庫使用(昼間)
call! 非常ベル(昼間)
call! 通常の通話(昼間)
現在時刻は11:00
record ... 金庫使用(昼間)
call! 非常ベル(昼間)
call! 通常の通話(昼間)
現在時刻は12:00
record ... 金庫使用(昼間)
call! 非常ベル(昼間)
call! 通常の通話(昼間)
現在時刻は13:00
record ... 金庫使用(昼間)
call! 非常ベル(昼間)
call! 通常の通話(昼間)
現在時刻は14:00
record ... 金庫使用(昼間)
call! 非常ベル(昼間)
call! 通常の通話(昼間)
現在時刻は15:00
record ... 金庫使用(昼間)
call! 非常ベル(昼間)
call! 通常の通話(昼間)
現在時刻は16:00
record ... 金庫使用(昼間)
call! 非常ベル(昼間)
call! 通常の通話(昼間)
現在時刻は17:00
[昼間]から[夜間]へ状態が変化しました
call! 非常:夜間の金庫使用
call! 非常ベル(夜間)
record ... 夜間の通話録音
現在時刻は18:00
call! 非常:夜間の金庫使用
call! 非常ベル(夜間)
record ... 夜間の通話録音
現在時刻は19:00
call! 非常:夜間の金庫使用
call! 非常ベル(夜間)
record ... 夜間の通話録音
現在時刻は20:00
call! 非常:夜間の金庫使用
call! 非常ベル(夜間)
record ... 夜間の通話録音
現在時刻は21:00
call! 非常:夜間の金庫使用
call! 非常ベル(夜間)
record ... 夜間の通話録音
現在時刻は22:00
call! 非常:夜間の金庫使用
call! 非常ベル(夜間)
record ... 夜間の通話録音
現在時刻は23:00
call! 非常:夜間の金庫使用
call! 非常ベル(夜間)
record ... 夜間の通話録音
`

	assert.Equal(t, expected, actual)
}
