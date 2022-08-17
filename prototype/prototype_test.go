package prototype

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMessageBox_use(t *testing.T) {
	mb := NewMessageBox('!')

	actual := mb.use("Hello!")

	expected := `!!!!!!!!!!
! Hello! !
!!!!!!!!!!
`
	assert.Equal(t, expected, actual)
}

func TestClone(t *testing.T) {
	// パッケージ制作者がやっておくこと
	// それぞれのパターンをクラスにしないでよくなる
	// ここでは2パターンしかないが10パターンとかある時には恩恵がある
	manager := NewManager()
	manager.register("warning box", NewMessageBox('*'))
	manager.register("slash box", NewMessageBox('/'))

	// パッケージ利用者側の操作
	warningBox, err := manager.create("warning box")
	assert.Nil(t, err)

	warningActual := warningBox.use("Hello")
	warningExpected := `*********
* Hello *
*********
`
	assert.Equal(t, warningExpected, warningActual)

	slashBox, err := manager.create("slash box")
	assert.Nil(t, err)

	slashActual := slashBox.use("Hello")
	slashExpected := `/////////
/ Hello /
/////////
`
	assert.Equal(t, slashExpected, slashActual)

}
