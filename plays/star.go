package plays

import (
	"game_fly/core"
	"github.com/hajimehoshi/ebiten"
	"math"
)

type Star struct {
	core.Sprite
}

func NewStar() (m *Star) {
	m = &Star{}
	m.Sprite.Create("map")
	return
}

func (s *Star) OnLoad() {

}

func (s *Star) Start() {
	s.SetMaterial(StarImg)
	s.SetScale(0.3, 0.3)
	//m.Collision = resolv.NewRectangle(m.GetPositionInt32())
	//sW, sH := conf.GetConfInt("scenes_width"), conf.GetConfInt("scenes_height")
	//m.SetScale(core.ScaliEq(float64(sW), float64(m.W)), core.ScaliEq(float64(sH), float64(m.H)))
}

func (s *Star) Update(screen *ebiten.Image) (err error) {
	s.Move(int32(math.Sin(float64(s.Y)/50)*2), 1)
	//d.Move(0, 1)

	s.Sprite.Update(screen)
	return nil
}
