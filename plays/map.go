package plays

import (
	"bytes"
	"game_fly/core"
	"game_fly/core/sprite"
	"game_fly/plays/assets/images"
	"game_fly/plays/conf"
	"github.com/hajimehoshi/ebiten"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"log"
)

type Map struct {
	core.Sprite
	GameNode Game
	RoleRode Role
}

func NewMap() (m *Map) {
	m = &Map{}
	m.Sprite.Create("map")
	return
}

func (m *Map) OnLoad() {
	//背景
	img1, _, err := image.Decode(bytes.NewReader(images.M2_jpg))
	if err != nil {
		log.Fatal(1, err)
	}
	MapImg, _ = ebiten.NewImageFromImage(img1, ebiten.FilterDefault)

	mapImgW, mapImgH := MapImg.Size()
	sW, sH := conf.CC.GetScenesWH()
	m.SetScale(core.ScaliEq(sW, float64(mapImgW)), core.ScaliEq(sH, float64(mapImgH)))
	m.SetXY(0, 0)
}

func (d *Map) Start() {
	d.GameNode = *core.GetScene("game").(*Game)
	d.RoleRode = *sprite.GetSprite("game", "role").(*Role)
}

func (d *Map) Update(screen *ebiten.Image) (err error) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(d.X, d.Y)
	op.GeoM.Scale(d.ScaleW, d.ScaleH)
	screen.DrawImage(MapImg, op)
	return nil
}
