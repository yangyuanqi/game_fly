package plays

import (
	"game_fly/core"
	"github.com/hajimehoshi/ebiten"
	"math"
)

type Bullet struct {
	core.Sprite
	stepX    float64 //子弹x轴偏移角度
	RootNode *Game
	Img      *ebiten.Image //子弹贴图
}

func NewBullet(rootNode *Game, p core.P, img *ebiten.Image) (bullet *Bullet) {
	x, y, w, h := p.GetPosition()
	bullet = &Bullet{}
	bullet.RootNode = rootNode
	bullet.Img = img

	bullet.SetScale(0.2, 0.2)

	rw, rh := BulletImg.Size()
	bullet.SetWH(int(math.Round(float64(rw)*bullet.ScaleW)), int(math.Round(float64(rh)*bullet.ScaleH)))
	bullet.SetXY(x+math.Ceil(float64(w-bullet.W)/2), y+math.Ceil(float64(h)/2))
	return
}

func (b *Bullet) Update() (err error) {
	b.Move()

	if ebiten.IsDrawingSkipped() {
		return nil
	}

	opts2 := &ebiten.DrawImageOptions{}
	opts2.GeoM.Scale(b.ScaleW, b.ScaleH)
	opts2.GeoM.Translate(b.X, b.Y)
	b.RootNode.Screen.DrawImage(b.Img, opts2)
	return
}
