package component

import (
	"github.com/325Gerbils/go-vector"
	"github.com/hajimehoshi/ebiten/v2"
)

type Transform struct {
	Position             vector.Vector
	RotationX, RotationY float64
	ScaleX, ScaleY       float64
}

func (t *Transform) Update() error {
	return nil
}

func (t *Transform) Draw(screen *ebiten.Image) {

}
