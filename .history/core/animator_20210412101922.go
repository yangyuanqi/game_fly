package core

import (
	"math"

	"github.com/hajimehoshi/ebiten"
)

type Animator struct {
	Sprite
	Interval float64
	Order    string
	Imgs     []*ebiten.Image
	ImgPng   *ebiten.Image
}

var count float64

func (a *Animator) Update(screen *ebiten.Image) (err error) {
	count++
	if int(math.Ceil(count/a.Interval)) < len(a.Imgs) {
		if count%a.Interval == 0 {

		}
	} else {
		count = 0
	}

	return nil
}
