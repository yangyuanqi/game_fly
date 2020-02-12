package component

import "github.com/325Gerbils/go-vector"

type Transform struct {
	Position             vector.Vector
	RotationX, RotationY float64
	ScaleX, ScaleY       float64
}
