package core

import "github.com/hajimehoshi/ebiten"

type Animator struct {
	Sprite
	Interval int
	Imgs     []*ebiten.Image
	ImgPng   *ebiten.Image
}

var count int

func (a *Animator) Update(screen *ebiten.Image) (err error) {
	count++
	if a.Interval%5 == 0 {

	}
	return nil
}
