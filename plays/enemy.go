package plays

import (
	"game_fly/core"
	"github.com/SolarLune/resolv/resolv"
	"github.com/hajimehoshi/ebiten"
	"math"
)

type Enemy struct {
	core.Sprite
	Img      *ebiten.Image //子弹贴图
	RoleNode Role
}

func NewEnemy(img *ebiten.Image) (enemy *Enemy) {
	e := &Enemy{}
	e.Sprite.Create()
	e.Img = img
	w, h := e.Img.Size()
	e.SetScale(0.3, 0.3)
	e.SetWH(int(math.Round(float64(w)*e.ScaleW)), int(math.Round(float64(h)*e.ScaleH)))
	e.SetXY(0, 0)

	e.Collision = resolv.NewRectangle(e.GetPosition())
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
	e.Draw(screen)
	return nil
}

//渲染
func (e *Enemy) Draw(screen *ebiten.Image) {
	opts2 := &ebiten.DrawImageOptions{}
	opts2.GeoM.Scale(e.ScaleW, e.ScaleH)
	opts2.GeoM.Translate(e.X, e.Y)
	screen.DrawImage(e.Img, opts2)
}
