package core

import "github.com/hajimehoshi/ebiten"

type Animator struct {
	Sprite
	Img    []*ebiten.Image
	ImgPng *ebiten.Image
}

func (a *Animator) Img() {

}

func (a *Animator) Update(screen *ebiten.Image) (err error) {

	return nil
}
