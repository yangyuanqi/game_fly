package plays

import (
	"bytes"
	"fmt"
	"game_fly/core"
	"game_fly/core/ui"
	"game_fly/plays/assets/images"
	"game_fly/plays/assets/images/ui_img"
	"github.com/hajimehoshi/ebiten"
	"image"
	"log"
	"strconv"
)

type Login struct {
	core.Sprite
	Screen *ebiten.Image //屏幕
}

var LoginImg *ebiten.Image

func NewLogin() (login *Login) {
	login = &Login{}
	login.W, login.H = 320, 480
	return
}

func (l *Login) OnLoad() {
	img, _, err := image.Decode(bytes.NewReader(images.Load))
	if err != nil {
		log.Fatal("loginImg:", err)
	}
	LoginImg, _ = ebiten.NewImageFromImage(img, ebiten.FilterDefault)

	LoginImgW, LoginImgH := LoginImg.Size()
	l.ScaleW, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", float64(l.W)/float64(LoginImgW)), 64)
	l.ScaleH, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", float64(l.H)/float64(LoginImgH)), 64)

	img, _, err = image.Decode(bytes.NewReader(ui_img.UI_png))
	if err != nil {
		log.Fatal(err)
	}
	UiImage, _ = ebiten.NewImageFromImage(img, ebiten.FilterDefault)

	button1.SetOnPressed(func(b *ui.Button) {
		core.SetScenesIng("game")
	})
}

var UiImage *ebiten.Image

var button1 = &ui.Button{
	Rect: image.Rect(110, 400, 210, 440),
	Text: "New Game",
}

func (l *Login) Update(screen *ebiten.Image) (err error) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(l.ScaleW, l.ScaleH)
	screen.DrawImage(LoginImg, op)

	button1.Update()
	button1.Draw(screen, UiImage)
	return nil
}
