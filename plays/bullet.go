package plays

import (
	"game_fly/core"
	"github.com/hajimehoshi/ebiten"
	"math"
)

type Bullet struct {
	core.Sprite
	RootNode *Game
	Img      *ebiten.Image //子弹贴图
}

func NewBullet(rootNode *Game, p core.P, img *ebiten.Image) (bullet *Bullet) {
	x, y, w, _ := p.GetPosition()
	bullet = &Bullet{}
	bullet.RootNode = rootNode
	bullet.Img = img
	bullet.Visible = true

	bullet.SetScale(0.2, 0.2)

	rw, rh := BulletImg.Size()
	bullet.SetWH(int(math.Round(float64(rw)*bullet.ScaleW)), int(math.Round(float64(rh)*bullet.ScaleH)))
	bullet.SetXY(x+math.Ceil(float64(w-bullet.W)/2), y)
	return
}

func (b *Bullet) Update() (err error) {
	if b.Visible == false {
		return nil
	}
	//到达边界删除
	if b.Y > float64(b.RootNode.H) || b.Y < 0 {
		b.Visible = false
	}

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
