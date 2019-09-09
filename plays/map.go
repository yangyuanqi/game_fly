package plays

import (
	"fmt"
	"game_fly/core"
	"game_fly/core/sprite"
	"github.com/hajimehoshi/ebiten"
	"strconv"
)

type Map struct {
	core.Sprite
	GameNode Game
	RoleRode Role
}

func (d *Map) Create() (m *Map) {
	d.Sprite.Create()
	d.Name = "map"
	m = d
	m.W, m.H = 320, 480
	return
}

func (d *Map) OnLoad() {
	MapImgW, MapImgH := MapImg.Size()
	d.ScaleW, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", float64(d.W)/float64(MapImgW)), 64)
	d.ScaleH, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", float64(d.H)/float64(MapImgH)), 64)
}

func (d *Map) Start() {
	d.GameNode = *core.GetScene("game").(*Game)
	d.RoleRode = *sprite.GetSprite("game", "role").(*Role)
}

func (d *Map) Update(screen *ebiten.Image) (err error) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(d.ScaleW, d.ScaleH)
	screen.DrawImage(MapImg, op)
	return nil
}
