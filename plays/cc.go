package plays

import (
	"github.com/SolarLune/resolv/resolv"
	"math"
)

type Sprite struct {
	X, Y         float64
	W, H         int
	ScaleW       float64
	ScaleH       float64
	RootNode     *Game             //根节点
	Collide      *resolv.Rectangle //碰撞绑定
	CollideSpace *resolv.Space     //像地形较复杂
}

func (s *Sprite) Move(offsetX, offsetY float64) {
	if offsetX > 0 {
		s.X += offsetX
		s.Collide.X = int32(s.X)
	} else if offsetX < 0 {
		s.X -= math.Abs(offsetX)
		s.Collide.X = int32(s.X)
	}

	if offsetY > 0 {
		s.Y += offsetY
		s.Collide.Y = int32(s.Y)
	} else if offsetY < 0 {
		s.Y -= math.Abs(offsetY)
		s.Collide.Y = int32(s.Y)
	}
}

func (s *Sprite) GetPosition() (x, y float64, w, h int) {
	return s.X, s.Y, s.W, s.H
}
