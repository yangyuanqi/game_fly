package core

import "github.com/hajimehoshi/ebiten"

type Animator struct {
	Sprite
	Interval int
	Imgs     []*ebiten.Image
	ImgPng   *ebiten.Image
}

func (a *Animator) Update(screen *ebiten.Image) (err error) {
	a.Interval++
	if a.Interval%5 == 0 {

	}
	return nil
}
