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
	IsStop   bool
}

var count float64

func (a *Animator) Update(screen *ebiten.Image) (err error) {
	if a.IsStop == false {
		count++
		fmt.Println("count:", count)
		if int(count)%int(a.Interval) == 0 && int(count/a.Interval)-1 < len(a.Imgs) {

			fmt.Println(int(count/a.Interval) - 1)
			a.SetMaterial(a.Imgs[int(count/a.Interval)])
		}
	}

	if int(count)%int(a.Interval) == 0 && int(count/a.Interval) > len(a.Imgs) {
		a.IsStop = true
		count = 0
	}

	return nil
}
