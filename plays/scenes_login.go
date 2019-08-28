package plays

import (
	"bytes"
	"fmt"
	"game_fly/core"
	"game_fly/plays/assets/images"
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
}

func (l *Login) Update(screen *ebiten.Image) (err error) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(l.ScaleW, l.ScaleH)
	screen.DrawImage(LoginImg, op)
	return nil
}
