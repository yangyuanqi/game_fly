package plays

import (
	"fmt"
	"game_fly/core"
	"github.com/Bredgren/geo"
	"github.com/hajimehoshi/ebiten"
	"image/color"
	"time"
)

type Hero struct {
	core.BaseComponent
	Jump int //1up,2down,0stop
}

var (
	img *ebiten.Image
)

func (h *Hero) OnLoad() {
	h.X = 0
	h.Y = 610
	h.W = 30
	h.H = 30

	img, _ = ebiten.NewImage(int(h.W), int(h.H), ebiten.FilterDefault)
	img.Fill(color.White)
}

var start =  time.Now()
func (h *Hero) Update(screen *ebiten.Image) {
	dt := time.Now().Sub(start)
	fmt.Println(dt)
	duration := 1 * time.Second
	t := geo.Clamp(dt.Seconds()/duration.Seconds(), 0, 1)
	if h.Jump == 1 {
		pos := geo.EaseVec(geo.VecXY(0, h.Y), geo.VecXY(0, h.Y-1), t, geo.EaseOutBack)
		h.Y = pos.Y
		fmt.Println(h.Y)
		if h.Y <= 640-30-200 {
			h.Jump = 2
		}
	} else if h.Jump == 2 {
		pos := geo.EaseVec(geo.VecXY(0, h.Y), geo.VecXY(0, h.Y+1), t, geo.EaseInBack)
		h.Y = pos.Y
		if h.Y >= 640-30 {
			h.Jump = 0
		}
	}

	opts2 := &ebiten.DrawImageOptions{}
	//opts2.GeoM.Scale(b.ScaleW, b.ScaleH)
	opts2.GeoM.Translate(h.X, h.Y)
	screen.DrawImage(img, opts2)
}
