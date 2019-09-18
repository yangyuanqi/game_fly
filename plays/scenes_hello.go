package plays

import (
	"game_fly/core"
	"game_fly/core/timer"
	"game_fly/core/ui"
	"github.com/hajimehoshi/ebiten"
	"time"
)

type ScenesHello struct {
	core.Sprite
}

func NewScenesHello() (scenesHello *ScenesHello) {
	scenesHello = &ScenesHello{}
	scenesHello.W, scenesHello.H = 300, 300
	return
}

var NumI int
var Num *ui.Number

func (s *ScenesHello) OnLoad() {
	//core.RegisterScene(NewGame(), "game")
	//core.SetScenesIng("game") //跳转到game场景
	timer.SetTicker(time.Millisecond*100, func() {
		ui.ReplaceUi("ui", "number1", ui.NewNumber(NumI, 100, 100, 0.5))
		NumI++
	}, -1)
}

func (s *ScenesHello) Update(screen *ebiten.Image) (err error) {

	return nil
}
