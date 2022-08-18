package builder

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHTMLBuilder(t *testing.T) {
	textBuilder := NewTextBuilder()
	director := NewDirector(textBuilder)
	director.Construct()

	assert.Equal(
		t,
		`==
「 Greeting 」

■朝から昼にかけて

 *おはようございます。
 *こんにちは。

■夜に

 *こんばんは。
 *おやすみなさい。
 *さようなら。

==
`,
		textBuilder.getResult(),
	)
}
