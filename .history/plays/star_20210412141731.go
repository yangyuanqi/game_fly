package plays

import (
	"game_fly/core"	"image"

	"github.com/hajimehoshi/ebiten"
)


type Star struct {
	core.Sprite
	count int
}

func (s *Star)NewStar()*Star{
	var a = &core.Animator{
		Sprite: core.Sprite{
			Base: core.Base{
				X:        100,
				Y:        100,
				Material: core.Nodes.Img["t1png"],
			},
		},
		Interval: 1,
		Order:    "loop",
		Imgs: []*ebiten.Image{
			core.Nodes.Img["itempng"].SubImage(image.Rect(0, 205, 60, 270)).(*ebiten.Image),
			core.Nodes.Img["itempng"].SubImage(image.Rect(0, 205, 60, 270)).(*ebiten.Image),
			core.Nodes.Img["itempng"].SubImage(image.Rect(0, 205, 60, 270)).(*ebiten.Image),
			core.Nodes.Img["itempng"].SubImage(image.Rect(0, 205, 60, 270)).(*ebiten.Image),
			core.Nodes.Img["itempng"].SubImage(image.Rect(136, 197, 183, 159)).(*ebiten.Image),
			core.Nodes.Img["itempng"].SubImage(image.Rect(136, 197, 183, 159)).(*ebiten.Image),
			core.Nodes.Img["itempng"].SubImage(image.Rect(136, 197, 183, 159)).(*ebiten.Image),
			core.Nodes.Img["itempng"].SubImage(image.Rect(136, 197, 183, 159)).(*ebiten.Image),
			core.Nodes.Img["itempng"].SubImage(image.Rect(0, 267, 57, 330)).(*ebiten.Image),
			core.Nodes.Img["itempng"].SubImage(image.Rect(0, 267, 57, 330)).(*ebiten.Image),
			core.Nodes.Img["itempng"].SubImage(image.Rect(0, 267, 57, 330)).(*ebiten.Image),
			core.Nodes.Img["itempng"].SubImage(image.Rect(0, 267, 57, 330)).(*ebiten.Image),
			core.Nodes.Img["itempng"].SubImage(image.Rect(83, 165, 125, 225)).(*ebiten.Image),
			core.Nodes.Img["itempng"].SubImage(image.Rect(83, 165, 125, 225)).(*ebiten.Image),
			core.Nodes.Img["itempng"].SubImage(image.Rect(83, 165, 125, 225)).(*ebiten.Image),
			core.Nodes.Img["itempng"].SubImage(image.Rect(83, 165, 125, 225)).(*ebiten.Image),
			core.Nodes.Img["itempng"].SubImage(image.Rect(337, 190, 410, 249)).(*ebiten.Image),
			core.Nodes.Img["itempng"].SubImage(image.Rect(337, 190, 410, 249)).(*ebiten.Image),
			core.Nodes.Img["itempng"].SubImage(image.Rect(337, 190, 410, 249)).(*ebiten.Image),
			core.Nodes.Img["itempng"].SubImage(image.Rect(337, 190, 410, 249)).(*ebiten.Image),
			core.Nodes.Img["itempng"].SubImage(image.Rect(332, 185, 371, 249)).(*ebiten.Image),
			core.Nodes.Img["itempng"].SubImage(image.Rect(332, 185, 371, 249)).(*ebiten.Image),
			core.Nodes.Img["itempng"].SubImage(image.Rect(332, 185, 371, 249)).(*ebiten.Image),
			core.Nodes.Img["itempng"].SubImage(image.Rect(332, 185, 371, 249)).(*ebiten.Image),
			core.Nodes.Img["itempng"].SubImage(image.Rect(84, 102, 126, 165)).(*ebiten.Image),
			core.Nodes.Img["itempng"].SubImage(image.Rect(84, 102, 126, 165)).(*ebiten.Image),
			core.Nodes.Img["itempng"].SubImage(image.Rect(84, 102, 126, 165)).(*ebiten.Image),
			core.Nodes.Img["itempng"].SubImage(image.Rect(84, 102, 126, 165)).(*ebiten.Image),
			core.Nodes.Img["itempng"].SubImage(image.Rect(127, 136, 188, 197)).(*ebiten.Image),
			core.Nodes.Img["itempng"].SubImage(image.Rect(127, 136, 188, 197)).(*ebiten.Image),
			core.Nodes.Img["itempng"].SubImage(image.Rect(127, 136, 188, 197)).(*ebiten.Image),
			core.Nodes.Img["itempng"].SubImage(image.Rect(127, 136, 188, 197)).(*ebiten.Image),
		},
	}
	a.Move = func(a *core.Animator) {
		a.Y += 1
	}
	core.Nodes.AddPrefab(a)
}