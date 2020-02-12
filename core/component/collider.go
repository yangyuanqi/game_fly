package component

import (
	"github.com/325Gerbils/go-vector"
	"github.com/SolarLune/resolv/resolv"
	"github.com/hajimehoshi/ebiten"
	"golang.org/x/image/colornames"
)

type Collider struct {
	Shape    *resolv.Rectangle
	Position vector.Vector
	//RotationX, RotationY float64
	img *ebiten.Image
}

func NewBoxCollider(xy vector.Vector, w, h float64) (collider *Collider) {
	coll := &Collider{}
	coll.Shape = resolv.NewRectangle(int32(xy.X), int32(xy.Y), int32(w), int32(h))
	coll.Position = xy

	coll.img, _ = ebiten.NewImage(int(w), int(h), ebiten.FilterDefault)
	coll.img.Fill(colornames.Green)
	return coll
}

func (coll *Collider) DrawCollider(screen *ebiten.Image) {
	p := &ebiten.DrawImageOptions{}
	px, py := coll.Position.X, coll.Position.Y
	p.GeoM.Translate(px, py)
	//设置碰撞体的位置
	coll.Shape.SetXY(int32(px), int32(py))
	screen.DrawImage(coll.img, p)
}

func (coll *Collider) Move(moveXY vector.Vector) {
		coll.Position.Add(moveXY)
}
