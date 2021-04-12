package core

import "github.com/hajimehoshi/ebiten"

type Animator struct {
	Sprite
	Interval int
	Imgs     []*ebiten.Image
	ImgPng   *ebiten.Image
}

func (a *Animator) Update(screen *ebiten.Image) (err error) {
	if len(a.Imgs) > 0 {

	}
	return nil
}
