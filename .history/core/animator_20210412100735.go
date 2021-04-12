package core

import "github.com/hajimehoshi/ebiten"

type Animator struct {
	Sprite
	Imgs   []*ebiten.Image
	ImgPng *ebiten.Image
}

func (a *Animator) Img(img []*ebiten.Image) {

}

func (a *Animator) Update(screen *ebiten.Image) (err error) {

	return nil
}
