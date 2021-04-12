package core

import (
	"github.com/hajimehoshi/ebiten"
)

type Animator struct {
	Sprite
	Interval float64
	Order    string
	Imgs     []*ebiten.Image
	ImgPng   *ebiten.Image
	IsStop   bool
	Count    float64
	Move     func()
}

var count float64

func (a *Animator) Update(screen *ebiten.Image) (err error) {
	if a.Move() != nil {

	}
	if a.IsStop == false {
		a.Count++
		//fmt.Println("count:", count)
		index := int(a.Count/a.Interval) - 1
		if int(a.Count)%int(a.Interval) == 0 && index < len(a.Imgs) {
			//fmt.Println(int(index))
			a.SetMaterial(a.Imgs[index])
			a.Sprite.Update(screen)
		}
	}

	if int(a.Count)%int(a.Interval) == 0 && int(a.Count/a.Interval)-1 > len(a.Imgs) {
		a.Count = 0
		if a.Order == "line" {
			a.IsStop = true
			a.Destroy()
		}
		if a.Order == "loop" {

		}
	}

	return nil
}
