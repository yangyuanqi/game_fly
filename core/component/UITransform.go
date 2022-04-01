package component

import "github.com/hajimehoshi/ebiten/v2"

type ContentSize struct {
	W float64
	H float64
}

type UITransform struct {
	Active      bool
	ContentSize ContentSize
	AnchorPoint Vec2
}

func (u *UITransform) Update() error {
	return nil
}

func (u *UITransform) Draw(screen *ebiten.Image) {

}
