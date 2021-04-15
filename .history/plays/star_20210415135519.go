package plays

import (
	"game_fly/core"
	"image"
	"math/rand"
	"time"

	"github.com/SolarLune/resolv"
	"github.com/hajimehoshi/ebiten"
)

var shape2 *resolv.Rectangle

func NewStar() *core.Animator {
	shape2 = resolv.NewRectangle(m.X, m.Y, m.W, m.H)

	rand.Seed(time.Now().Unix())
	n := rand.Intn(512)

	var a = &core.Animator{
		Sprite: core.Sprite{
			Base: core.Base{
				X: int32(n),
				Y: 0,
				// Material: core.Nodes.Img["itempng"],
			},
		},
		Interval: 4,
		Order:    "loop",
		Imgs: []*ebiten.Image{
			core.Nodes.Img["itempng"].SubImage(image.Rect(0, 205, 60, 270)).(*ebiten.Image),
			core.Nodes.Img["itempng"].SubImage(image.Rect(136, 197, 183, 159)).(*ebiten.Image),
			core.Nodes.Img["itempng"].SubImage(image.Rect(0, 267, 57, 330)).(*ebiten.Image),
			core.Nodes.Img["itempng"].SubImage(image.Rect(83, 165, 125, 225)).(*ebiten.Image),
			core.Nodes.Img["itempng"].SubImage(image.Rect(332, 185, 371, 249)).(*ebiten.Image),
			core.Nodes.Img["itempng"].SubImage(image.Rect(84, 102, 126, 165)).(*ebiten.Image),
			core.Nodes.Img["itempng"].SubImage(image.Rect(127, 136, 188, 197)).(*ebiten.Image),
		},
	}
	return a
}
