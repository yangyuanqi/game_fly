package core

import (
	"fmt"

	"github.com/hajimehoshi/ebiten"
)

type Animator struct {
	Sprite
	Interval float64
	Order    string
	Imgs     []*ebiten.Image
	ImgPng   *ebiten.Image
	IsStop   false
}

var count float64

func (a *Animator) Update(screen *ebiten.Image) (err error) {
	count++
	if int(count)%int(a.Interval) == 0 && int(count/a.Interval) < len(a.Imgs) {
		fmt.Println(int(count/a.Interval) - 1)
		a.SetMaterial(a.Imgs[int(count/a.Interval)])
	}
	if int(count)%int(a.Interval) == 0 && int(count/a.Interval) > len(a.Imgs) {
		count = 0
	}

	return nil
}
