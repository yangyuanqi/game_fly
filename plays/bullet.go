package plays

import (
	"game_fly/core"
	"game_fly/core/data"
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
	b.Sprite.Create("bullet")
	//获取节点
	b.roleNode = sprite.GetSprite("game", "role").(*Role)
	b.gameNode = *core.GetScene("game").(*Game)
	//贴图
	b.Material = img
	//初始化属性
	b.SetScale(0.2, 0.2)
	rw, rh := BulletImg.Size()
	b.SetWH(math.Round(float64(rw)*b.ScaleW), math.Round(float64(rh)*b.ScaleH))
	b.SetXY(b.roleNode.X+math.Ceil(float64(b.roleNode.W-b.W)/2), b.roleNode.Y)
	//添加碰撞
	b.Collision = resolv.NewRectangle(b.GetPositionInt32())
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

	roleBullet := prefab.GetPrefabAll("roleBullet")
	enemy := prefab.GetPrefabAll("enemy")
	for _, v := range roleBullet {
		for _, v2 := range enemy {
			colliding := v.(data.SpriteComponent).GetResolv().(*resolv.Rectangle).IsColliding(v2.GetResolv().(*resolv.Rectangle))
			if colliding {
				v2.(*Enemy).SouShang++
				prefab.DelPrefab("roleBullet", v.(data.SpriteComponent).GetId())
			}
		}
	}

	if ebiten.IsDrawingSkipped() {
		return nil
	}

	opts2 := &ebiten.DrawImageOptions{}
	opts2.GeoM.Scale(b.ScaleW, b.ScaleH)
	opts2.GeoM.Translate(b.X, b.Y)
	screen.DrawImage(b.Material, opts2)
	return
}
