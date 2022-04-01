package component

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
)

type Script struct {
}

func (s *Script) Update(dt float64) error {
	fmt.Println("script")
	return nil
}

func (s *Script) Draw(screen *ebiten.Image) {

}

//
//func (s *Script) Draw(screen *ebiten.Image) {
//	opts := &ebiten.DrawImageOptions{}
//	//optb.GeoM.Reset()
//	opts.GeoM.Scale(s.Transform.ScaleX, s.Transform.ScaleY)
//	//optb.GeoM.Rotate(-1.1)//旋转
//	//optb.GeoM.Skew(1,1)//倾斜
//	opts.GeoM.Translate(s.Collider.Transform.Position.X, s.Collider.Transform.Position.Y)
//	s.Transform.Position = s.Collider.Transform.Position
//	screen.DrawImage(s.Sprite.Img, opts)
//}
