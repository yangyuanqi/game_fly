package component

import (
	"github.com/325Gerbils/go-vector"
	"github.com/SolarLune/resolv/resolv"
	"github.com/hajimehoshi/ebiten"
	"golang.org/x/image/colornames"
)

type Collider struct {
	Shape     *resolv.Rectangle
	Transform Transform
	//RotationX, RotationY float64
	img       *ebiten.Image
	IsTrigger bool //碰撞边界触发
}

func NewBoxCollider(xy vector.Vector, w, h float64) (collider *Collider) {
	coll := &Collider{
		Shape:resolv.NewRectangle(int32(xy.X), int32(xy.Y), int32(w), int32(h)),
		Transform: Transform{
			ScaleX: 1,
			ScaleY: 1,
			Position:xy,
		},
	}

	coll.img, _ = ebiten.NewImage(int(w), int(h), ebiten.FilterDefault)
	coll.img.Fill(colornames.Green)
	return coll
}

func (coll *Collider) DrawCollider(screen *ebiten.Image) {
	p := &ebiten.DrawImageOptions{}
	px, py := coll.Transform.Position.X, coll.Transform.Position.Y
	p.GeoM.Translate(px, py)
	//设置碰撞体的位置
	coll.Shape.SetXY(int32(px), int32(py))
	//screen.DrawImage(coll.img, p)
}

func (coll *Collider) Move(moveXY vector.Vector) {
	coll.Transform.Position.Add(moveXY)
}

func (coll *Collider) GetSize() (sw, sh float64) {
	w, h := coll.img.Size()
	sw = float64(w) * coll.Transform.ScaleX
	sh = float64(h) * coll.Transform.ScaleY
	return
}
