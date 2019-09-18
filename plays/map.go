package plays

import (
	"bytes"
	"fmt"
	"game_fly/core"
	"game_fly/core/sprite"
	"game_fly/plays/assets/images"
	"github.com/hajimehoshi/ebiten"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"strconv"
)

type Map struct {
	core.Sprite
	GameNode Game
	RoleRode Role
}

func NewMap() (m *Map) {
	m = &Map{}
	m.Sprite.Create()
	m.Name = "map"
	m.W, m.H = 320, 480
	return
}

func (d *Map) OnLoad() {
	//背景
	img1, _, err := image.Decode(bytes.NewReader(images.M2_jpg))
	if err != nil {
		log.Fatal(1, err)
	}
	MapImg, _ = ebiten.NewImageFromImage(img1, ebiten.FilterDefault)

	MapImgW, MapImgH := MapImg.Size()
	d.ScaleW, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", float64(d.W)/float64(MapImgW)), 64)
	d.ScaleH, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", float64(d.H)/float64(MapImgH)), 64)
	//timer.SetTicker(time.Millisecond*5, func() {
	//	d.Y++
	//	if d.Y > 0 {
	//		d.Y = (-2048) + 480
	//	}
	//}, -1)
	d.X = 0
	d.Y = 0
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
	//fmt.Println(MapImg.Size())
	return nil
}
