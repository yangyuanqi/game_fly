package plays

import (
	"bytes"
	"github.com/SolarLune/resolv/resolv"
	"github.com/hajimehoshi/ebiten"
	"image"
	"log"
	"math"
	"test/play/plays/assets/images"
)

type Bullet struct {
	CC
	stepX float64 //子弹x轴偏移角度
}

var BulletImg *ebiten.Image
var BulletImg2 *ebiten.Image


func init() {
	img, _, err := image.Decode(bytes.NewReader(images.Myb_1))
	if err != nil {
		log.Fatal(err)
	}
	BulletImg, _ = ebiten.NewImageFromImage(img, ebiten.FilterDefault)

	img, _, err = image.Decode(bytes.NewReader(images.Epb_1))
	if err != nil {
		log.Fatal(err)
	}

}

func NewBullet(game *Game, p P,stepX float64) (bullet *Bullet) {
	x, y, w, h := p.GetPosition()
	bullet = &Bullet{}
	bullet.RootNode = game
	bullet.ScaleH = 0.2
	bullet.ScaleW = 0.2

	rw, rh := BulletImg.Size()
	bullet.W = int(math.Round(float64(rw) * bullet.ScaleW))
	bullet.H = int(math.Round(float64(rh) * bullet.ScaleH))

	bullet.X = x + math.Ceil(float64(w-bullet.W)/2)
	bullet.Y = y + math.Ceil(float64(h)/2)

	bullet.stepX = stepX
	bullet.Collide = resolv.NewRectangle(int32(bullet.X), int32(bullet.Y), int32(bullet.W), int32(bullet.H))
	return
}

func (b *Bullet) Onload() {

}

func (b *Bullet) Update(points string) (err error) {
	if points == "hero" {
		b.Move(b.stepX, -5)
	} else if points == "enemy" {
		b.Move(b.stepX,5)
	}
	if ebiten.IsDrawingSkipped() {
		return nil
	}

	b.Draw()
	return
}

func (b *Bullet) Draw() {
	opts2 := &ebiten.DrawImageOptions{}
	opts2.GeoM.Scale(b.ScaleW, b.ScaleH)
	opts2.GeoM.Translate(b.X, b.Y)
	b.RootNode.screen.DrawImage(BulletImg, opts2)
}

func (b *Bullet) GetPosition() (x, y float64, w, h int) {
	return b.X, b.Y, b.W, b.H
}
