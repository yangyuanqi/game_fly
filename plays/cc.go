package plays

import (
	"github.com/SolarLune/resolv/resolv"
	"math"
)

type CC struct {
	X, Y         float64
	W, H         int
	ScaleW       float64
	ScaleH       float64
	RootNode     *Game             //根节点
	Collide      *resolv.Rectangle //碰撞绑定
	CollideSpace *resolv.Space     //碰撞空间绑定
}

func (c *CC) Move(offsetX, offsetY float64) {
	if offsetX > 0 {
		c.X += offsetX
		c.Collide.X = int32(c.X)
	} else if offsetX < 0 {
		c.X -= math.Abs(offsetX)
		c.Collide.X = int32(c.X)
	}

	if offsetY > 0 {
		c.Y += offsetY
		c.Collide.Y = int32(c.Y)
	} else if offsetY < 0 {
		c.Y -= math.Abs(offsetY)
		c.Collide.Y = int32(c.Y)
	}
}
