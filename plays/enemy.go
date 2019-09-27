package plays

import (
	"game_fly/core"
	"game_fly/core/sprite"
	"github.com/SolarLune/resolv/resolv"
	"github.com/hajimehoshi/ebiten"
	"math"
)

type Enemy struct {
	core.Sprite
	SouShang int
	Move     func(e *Enemy)
}

func NewEnemy(img *ebiten.Image) (enemy *Enemy) {
	e := &Enemy{}
	e.Sprite.Create("enemy")
	e.Material = img

	e.SetScale(0.3, 0.3)
	w, h := e.Material.Size()
	e.SetWH(math.Round(float64(w)*e.ScaleW), math.Round(float64(h)*e.ScaleH))
	e.SetXY(0, 0)

	e.Collision = resolv.NewRectangle(e.GetPositionInt32())
	return e
}

//加载资源
func (e *Enemy) OnLoad() {

}

func (e *Enemy) Start() {

}

//更新
func (e *Enemy) Update(screen *ebiten.Image) (err error) {
	e.Collision.SetXY(int32(e.X), int32(e.Y))
	if e.SouShang >= 50 {
		sprite.GetSprite("game", "role").(*Role).Fraction++
		e.Destroy = true
	}
	e.Move(e)
	e.Draw(screen)
	return nil
}

//渲染
func (e *Enemy) Draw(screen *ebiten.Image) {
	opts2 := &ebiten.DrawImageOptions{}
	opts2.GeoM.Scale(e.ScaleW, e.ScaleH)
	opts2.GeoM.Translate(e.X, e.Y)
	screen.DrawImage(e.Material, opts2)
}
