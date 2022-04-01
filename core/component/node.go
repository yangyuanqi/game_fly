package component

import "github.com/hajimehoshi/ebiten/v2"

type Vec3 struct {
	X float64
	Y float64
	Z float64
}
type Vec2 struct {
	X float64
	Y float64
}

type Node struct {
	Active   bool
	Position Vec3
	Rotation Vec3
	Scale    Vec2
	Layer    uint
}

func (n *Node) Update(dt float64) error {
	return nil
}

func (n *Node) Draw(screen *ebiten.Image) {

}
