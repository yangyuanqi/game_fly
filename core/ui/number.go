//计数
package ui

import (
	"bytes"
	"game_fly/core"
	"game_fly/core/ui/ui_img"
	"github.com/hajimehoshi/ebiten"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"strconv"
)

type Number struct {
	//data.UiType
	core.Sprite
	n    int
	strN []byte
	w    int
	sum  int
	Zoom float64
}

var NumberImg *ebiten.Image

func NewNumber(n int, x, y, zoom float64) (num *Number) {
	num = &Number{}
	num.Name = "number1"
	num.Zoom = zoom
	num.n = n
	num.w = 17
	num.sum = 10
	num.strN = []byte(strconv.Itoa(n))
	num.X = x //第一个数字的位置
	num.Y = y
	return
}


func (n *Number) OnLoad() {
	img, _, err := image.Decode(bytes.NewReader(ui_img.Shuzi_png))
	if err != nil {
		log.Fatal("number", err)
	}
	NumberImg, _ = ebiten.NewImageFromImage(img, ebiten.FilterDefault)
}

func (n *Number) Start() {

}

func (n *Number) Update(screen *ebiten.Image) (err error) {
	opts := &ebiten.DrawImageOptions{}

	for k, v := range n.strN {
		num, _ := strconv.Atoi(string(v))
		opts.GeoM.Reset()
		opts.GeoM.Scale(n.Zoom, n.Zoom)
		if k == 0 {
			opts.GeoM.Translate(n.X, n.Y)
		} else {
			opts.GeoM.Translate(n.X+float64(k*17)*n.Zoom, n.Y)
		}

		screen.DrawImage(NumberImg.SubImage(image.Rect(num*n.w, 0, num*n.w+17, 24)).(*ebiten.Image), opts)
	}
	return nil
}
