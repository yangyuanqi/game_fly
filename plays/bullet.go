package plays

import (
	"game_fly/core"
	"game_fly/core/prefab"
	"game_fly/core/sprite"
	"github.com/SolarLune/resolv/resolv"
	"github.com/hajimehoshi/ebiten"
	"math"
)

type Bullet struct {
	core.Sprite
	roleNode *Role
	gameNode Game
	Move     func(bullet *Bullet)
}

func NewBullet(img *ebiten.Image) (bullet *Bullet) {
	b := &Bullet{}
	b.Sprite.Create()
	//获取节点
	b.roleNode = sprite.GetSprite("game", "role").(*Role)
	b.gameNode = *core.GetScene("game").(*Game)
	//贴图
	b.Material = img
	//初始化属性
	b.SetScale(0.2, 0.2)
	rw, rh := BulletImg.Size()
	b.SetWH(int(math.Round(float64(rw)*b.ScaleW)), int(math.Round(float64(rh)*b.ScaleH)))
	b.SetXY(b.roleNode.X+math.Ceil(float64(b.roleNode.W-b.W)/2), b.roleNode.Y)
	//添加碰撞
	b.Collision = resolv.NewRectangle(b.GetPosition())
	return b
}

func (b *Bullet) OnLoad() {

}

func (b *Bullet) Start() {

}

func (b *Bullet) Update(screen *ebiten.Image) (err error) {
	//到达边界删除b
	if b.Y > float64(b.gameNode.H) || b.Y < 0 {
		prefab.DelPrefab("roleBullet", b.Id)
	}
	b.Move(b)

	if ebiten.IsDrawingSkipped() {
		return nil
	}

	opts2 := &ebiten.DrawImageOptions{}
	opts2.GeoM.Scale(b.ScaleW, b.ScaleH)
	opts2.GeoM.Translate(b.X, b.Y)
	screen.DrawImage(b.Material, opts2)
	return
}
