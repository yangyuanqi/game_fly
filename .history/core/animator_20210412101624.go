package core

import (
	"math"

	"github.com/hajimehoshi/ebiten"
)

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
	if math.Ceil(count/a.Interval) > len(a.Imgs) {

	}

	if count%a.Interval == 0 {

	}
	return nil
}
