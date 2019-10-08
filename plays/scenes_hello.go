package plays

import (
	"game_fly/core"
	"game_fly/core/prefab"
	"game_fly/core/ui"
	"github.com/hajimehoshi/ebiten"
	"image/color"
)

type ScenesHello struct {
	core.Sprite
}

func NewScenesHello() (scenesHello *ScenesHello) {
	scenesHello = &ScenesHello{}
	scenesHello.W, scenesHello.H = 300, 300
	return
}

func (s *ScenesHello) OnLoad() {
	t := &ui.Text{
		TimeLong: 30,
		FontSize: 16,
		TextX:    100,
		TextY:    150,
		Color:    color.NRGBA{255, 0, 0, 255},
		Text:     "9999",
		TextType: 0,
	}
	t.Create("text")
	prefab.AddPrefab(t, "text")

	t2 := &ui.Text{
		TimeLong: 600,
		FontSize: 16,
		TextX:    200,
		TextY:    150,
		Color:    color.NRGBA{255, 255, 255, 255},
		Text:     "hello world \n haha",
		TextType: 1,
	}
	t2.Create("text")
	prefab.AddPrefab(t2, "say")
}

func (s *ScenesHello) Update(screen *ebiten.Image) (err error) {


	return nil
}
