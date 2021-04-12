package core

import "github.com/hajimehoshi/ebiten"

type Animator struct {
	Sprite
	Imgs     []*ebiten.Image
	ImgPng   *ebiten.Image
	Interval int
}

func (a *Animator) Update(screen *ebiten.Image) (err error) {
	if len(a.Imgs) > 0 {

	}
	return nil
}
