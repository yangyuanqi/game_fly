package core

import (
	"core"

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

func (a *Animator) NewAnimator() *Animator {
	a := &Animator{}
	a.Sprite.Create("bullet")
	//b.SetMaterial(core.Byte2Image(images.Myb_1))
	a.SetMaterial(core.Nodes.Img["bullet"])
	a.SetScale(1, 1)
	return a
}

var count float64

func (a *Animator) Update(screen *ebiten.Image) (err error) {
	// if a.IsStop == false {
	// 	count++
	// 	fmt.Println("count:", count)
	// 	index := int(count/a.Interval) - 1
	// 	if int(count)%int(a.Interval) == 0 && index < len(a.Imgs) {

	// 		fmt.Println(int(index))
	// 		a.SetMaterial(a.Imgs[index])
	// 	}
	// }

	// if int(count)%int(a.Interval) == 0 && int(count/a.Interval)-1 > len(a.Imgs) {
	// 	a.IsStop = true
	// 	count = 0
	// }
	// a.SetMaterial(Nodes.Img["t1png"])
	a.Update(screen)

	return nil
}
