package core

import (
	"github.com/hajimehoshi/ebiten"
)

type Animator struct {
	Sprite
	Interval float64
	Order    string
	Imgs     []*ebiten.Image
	ImgPng   *ebiten.Image
	IsStop   bool
	Count    float64
	Move     func(a *Animator)
}

var count float64

func (a *Animator) Update(screen *ebiten.Image) (err error) {
	a.Move(a)
	if a.IsStop == false {
		index := int(a.Count / a.Interval)
		if int(a.Count)%int(a.Interval) == 0 && index < len(a.Imgs) {
			a.SetMaterial(a.Imgs[index])
		}
		a.Count++
		if a.Material != nil {
			a.Sprite.Update(screen)
		}
	}

	if int(a.Count)%int(a.Interval) == 0 && int(a.Count/a.Interval) >= len(a.Imgs) {
		a.Count = 0
		if a.Order == "line" {
			a.IsStop = true
			a.Destroy()
		}
		if a.Order == "loop" {

		}
	}

	return nil
}
