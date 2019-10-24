package plays

import (
	"game_fly/core"
	"game_fly/plays/assets/images"
	"github.com/hajimehoshi/ebiten"
)

var Scenes *ScenesHello

type ScenesHello struct {
	core.RectangleSprite
	r    *ebiten.Image
	role *Role
}

func (s *ScenesHello) OnLoad() {
	s.SetMaterial(core.Byte2Image(images.M2_jpg))
	s.SetScaleWH(1, 1)
	//s.SetXY()

	s.r = core.Byte2Image(images.My_1)

	//bg, _ = ebiten.NewImage(320, 480, ebiten.FilterDefault)

}

func (s *ScenesHello) Start() {
	s.role = &Role{}
	s.Child = append(s.Child, s.role)
}

func (s *ScenesHello) Update(screen *ebiten.Image) (err error) {

	//if s.LR == 1 {
	//	if s.role.X >= 320-50-s.role.W {
	//		s.Move(-2, 0)
	//	}
	//}
	//
	//if s.LR == -1 {
	//	if s.role.X <= 50 {
	//		s.Move(2, 0)
	//	}
	//}

	//s.UpdateTiming()
	//opts := &ebiten.DrawImageOptions{}

	//opts.GeoM.Skew(s.SkewX, s.SkewY)
	//opts.GeoM.Scale(s.ScaleW, s.ScaleH)

	//渲染精灵，组件
	s.RectangleSprite.Update(screen)

	//screen.DrawImage(bg, opts)
	core.GetComponentUpdate(screen, s)

	//	opts.GeoM.Translate(a, b)
	//	s.Material.DrawImage(s.r, opts)

	//if s.UD == 1 {
	//	s.Move(0, 1)
	//}
	//if s.UD == -1 {
	//	s.Move(0, -1)
	//}

	return nil
}
