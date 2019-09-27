package plays

import (
	"bytes"
	"game_fly/core"
	"game_fly/core/ui"
	"game_fly/plays/assets/images"
	"game_fly/plays/assets/images/ui_img2"
	"game_fly/plays/conf"
	"github.com/hajimehoshi/ebiten"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"log"
)

type Login struct {
	core.Sprite
}

var LoginImg *ebiten.Image

func NewLogin() (login *Login) {
	login = &Login{}
	return
}

func (l *Login) OnLoad() {
	img, _, err := image.Decode(bytes.NewReader(images.M1_jpg))
	if err != nil {
		log.Fatal("loginImg:", err)
	}
	LoginImg, _ = ebiten.NewImageFromImage(img, ebiten.FilterDefault)

	LoginImgW, LoginImgH := LoginImg.Size()
	sw, sh := conf.CC.GetScenesWH()
	l.SetScale(core.ScaliEq(sw, float64(LoginImgW)),core.ScaliEq(sh, float64(LoginImgH)))

	img, _, err = image.Decode(bytes.NewReader(ui_img2.UI_png))
	if err != nil {
		log.Fatal("login2", err)
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
