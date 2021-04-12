package core

import "github.com/hajimehoshi/ebiten"

type Animator struct {
	Sprite
	Interval int
	Order    string
	Imgs     []*ebiten.Image
	ImgPng   *ebiten.Image
}

var count int

func (a *Animator) Update(screen *ebiten.Image) (err error) {
	count++
	if 
	
	if count%a.Interval == 0 {

	}
	return nil
}
