package ui

import (
	"game_fly/core/data"
	"golang.org/x/image/font"
	"image/color"
)

type Text struct {
	data.Sprite
	uiFont       font.Face
	TextType     int         //弹出类型
	TimeLong     int         //持续时间
	FontSize     float64     //字体大小
	Text         string      //文本
	Color        color.NRGBA //字体颜色
	TextX, TextY int32
}
//
//func (t *Text) OnLoad() {
//	t.SetXY(t.TextX, t.TextY)
//
//	tt, err := truetype.Parse(goregular.TTF)
//	if err != nil {
//		log.Fatal(err)
//	}
//	t.uiFont = truetype.NewFace(tt, &truetype.Options{
//		Size:    t.FontSize,
//		DPI:     72,
//		Hinting: font.HintingFull,
//	})
//
//	if t.TextType == 2 {
//		t.W = int32(len([]rune(t.Text)) * 8)
//		t.Material, _ = ebiten.NewImage(int(t.W), int(t.H), ebiten.FilterDefault)
//		t.Material.Fill(color.White)
//	}
//}
//
//func (t *Text) Start() {
//
//}
//
//func (t *Text) Update(screen *ebiten.Image) (err error) {
//	t.UpdateTiming()
//
//	if t.TextType == 0 {
//		if t.Timing < t.TimeLong {
//			t.Move(0, int32(-t.Timing/10))
//			sy := t.Color.A / (uint8(t.TimeLong) * 3)
//			t.Color.A = uint8(t.Color.A - uint8(t.Timing)*sy)
//			text.Draw(screen, t.Text, t.uiFont, int(t.X), int(t.Y), t.Color)
//		} else {
//			prefab.DelPrefab(t.Id)
//		}
//	}
//
//	if t.TextType == 1 {
//		if t.Timing < t.TimeLong {
//			text.Draw(screen, t.Text, t.uiFont, int(t.X), int(t.Y), t.Color)
//			//if t.Timing > t.TimeLong/3*2 {
//			//sy := float64(t.Color.A / uint8(t.TimeLong/3))
//			//t.Color.A = uint8(float64(t.Color.A) - sy*(float64(t.Timing)-float64(t.TimeLong/3*2)))
//			//text.Draw(screen, t.Text, t.uiFont, int(t.X), int(t.Y), t.Color)
//			//}
//		} else {
//			prefab.DelPrefab(t.Id)
//		}
//	}
//
//	if t.TextType == 2 {
//		op := &ebiten.DrawImageOptions{}
//		op.GeoM.Translate(float64(t.X), float64(t.Y))
//		screen.DrawImage(t.Material, op)
//
//		text.Draw(screen, t.Text, t.uiFont, int(t.X), int(t.Y+20), t.Color)
//	}
//	return
//}
//
//func (t *Text) DelPrefab() {
//	prefab.DelPrefab(t.Id)
//}
