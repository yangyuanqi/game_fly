//计数
package ui

import (
	"game_fly/core"
	"game_fly/core/ui/ui_img"
	"github.com/hajimehoshi/ebiten"
	"image"
	"strconv"
)

type Number struct {
	core.Sprite
	n    int
	strN []byte
	sum  int
	wh   float64
}

//20*20 10
func NewNumber(n int, x, y, wh float64) (num *Number) {

	num = &Number{}
	num.Name = "number1"
	num.n = n
	num.sum = 10
	num.strN = []byte(strconv.Itoa(n))
	num.X = int32(x) //第一个数字的位置
	num.Y = int32(y)

	num.SetMaterial(core.Byte2Image(ui_img.Member_png))
	w, h := num.Material.Size()

	num.ScaleW = wh * 10 / float64(w)
	num.ScaleH = wh / float64(h)
	num.wh = wh

	return
}

func (n *Number) OnLoad() {

}

func (n *Number) Start() {

}

func (n *Number) Update(screen *ebiten.Image) (err error) {
	opts := &ebiten.DrawImageOptions{}
	for k, v := range n.strN {
		num, _ := strconv.Atoi(string(v))
		opts.GeoM.Reset()
		opts.GeoM.Scale(n.ScaleW, n.ScaleH)
		if k == 0 {
			opts.GeoM.Translate(float64(n.X), float64(n.Y))
		} else {
			opts.GeoM.Translate(float64(n.X)+float64(k)*n.wh, float64(n.Y))
		}

		screen.DrawImage(n.Material.SubImage(image.Rect(num*int(n.W/10), 0, num*int(n.W/10)+int(n.W/10), 128)).(*ebiten.Image), opts)
	}
	return nil
}
