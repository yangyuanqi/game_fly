package plays

import (
	"bytes"
	"fmt"
	"game_fly/core"
	"game_fly/core/ui"
	"game_fly/plays/assets/images"
	"game_fly/plays/assets/images/ui_img2"
	"github.com/hajimehoshi/ebiten"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"strconv"
)

type Login struct {
	core.Sprite
}

var LoginImg *ebiten.Image

func NewLogin() (login *Login) {
	login = &Login{}
	login.W, login.H = 320, 480
	return
}

func (l *Login) OnLoad() {
	img, _, err := image.Decode(bytes.NewReader(images.M1_jpg))
	if err != nil {
		log.Fatal("loginImg:", err)
	}
	LoginImg, _ = ebiten.NewImageFromImage(img, ebiten.FilterDefault)

	LoginImgW, LoginImgH := LoginImg.Size()
	l.ScaleW, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", float64(l.W)/float64(LoginImgW)), 64)
	l.ScaleH, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", float64(l.H)/float64(LoginImgH)), 64)

	img, _, err = image.Decode(bytes.NewReader(ui_img2.UI_png))
	if err != nil {
		log.Fatal("login2",err)
	}
	UiImage, _ = ebiten.NewImageFromImage(img, ebiten.FilterDefault)

	button1.SetOnPressed(func(b *ui.Button) {
		//初始化后，切换到下一个场景,可配置加载场景动画
		core.RegisterScene(NewGame(), "game")
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
