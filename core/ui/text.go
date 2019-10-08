package ui

import (
	"game_fly/core"
	"game_fly/core/prefab"
	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/gofont/goregular"
	"image/color"
	"log"
)

type Text struct {
	core.Sprite
	uiFont       font.Face
	TextType     int
	TimeLong     int     //持续时间
	FontSize     float64 //字体大小
	TextX, TextY int32   //位置
	Text         string
	Color        color.NRGBA
}

func (t *Text) OnLoad() {
	if t.TextType == 0 || t.TextType == 1 {
		t.SetXY(t.TextX, t.TextY)
		tt, err := truetype.Parse(goregular.TTF)
		if err != nil {
			log.Fatal(err)
		}
		t.uiFont = truetype.NewFace(tt, &truetype.Options{
			Size:    t.FontSize,
			DPI:     72,
			Hinting: font.HintingFull,
		})
	}

}

func (t *Text) Start() {

}

func (t *Text) Update(screen *ebiten.Image) (err error) {
	t.UpdateTiming()

	if t.TextType == 0 {
		if t.Timing < t.TimeLong {
			t.Move(0, int32(-t.Timing/10))
			sy := t.Color.A / (uint8(t.TimeLong) * 3)
			t.Color.A = uint8(t.Color.A - uint8(t.Timing)*sy)
			text.Draw(screen, t.Text, t.uiFont, int(t.X), int(t.Y), t.Color)
		} else {
			prefab.DelPrefab(t.Id)
		}
	}

	if t.TextType == 1 {
		if t.Timing < t.TimeLong {
			text.Draw(screen, t.Text, t.uiFont, int(t.X), int(t.Y), t.Color)
			if t.Timing > t.TimeLong/3*2 {
				sy := float64(t.Color.A / uint8(t.TimeLong/3))
				t.Color.A = uint8(float64(t.Color.A) - sy*(float64(t.Timing)-float64(t.TimeLong/3*2)))
				text.Draw(screen, t.Text, t.uiFont, int(t.X), int(t.Y), t.Color)
			}
		} else {
			prefab.DelPrefab(t.Id)
		}
	}
	return
}
