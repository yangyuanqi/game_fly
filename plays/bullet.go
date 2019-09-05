package plays

import (
	"game_fly/core"
	"game_fly/core/component"
	"game_fly/core/prefab"
	"github.com/hajimehoshi/ebiten"
	"math"
	"strconv"
)

type Bullet struct {
	core.Sprite
	Img      *ebiten.Image //子弹贴图
	roleNode Role
	gameNode Game
	Move     func(bullet *Bullet)
}

func NewBullet(img *ebiten.Image) (bullet *Bullet) {
	bullet = &Bullet{}
	//生成id
	core.ID++
	id := "bullet" + strconv.Itoa(core.ID)
	bullet.Id = id
	//获取节点
	bullet.roleNode = *component.GetComponent("role").(*Role)
	bullet.gameNode = *core.GetScene("game").(*Game)
	//贴图
	bullet.Img = img
	//初始化属性
	bullet.Visible = true
	bullet.SetScale(0.2, 0.2)
	bullet.MY = -10
	rw, rh := BulletImg.Size()
	bullet.SetWH(int(math.Round(float64(rw)*bullet.ScaleW)), int(math.Round(float64(rh)*bullet.ScaleH)))
	bullet.SetXY(bullet.roleNode.X+math.Ceil(float64(bullet.roleNode.W-bullet.W)/2), bullet.roleNode.Y)
	return
}

func (b *Bullet) OnLoad() {

}

func (b *Bullet) Update(screen *ebiten.Image) (err error) {
	//到达边界删除b
	if b.Y > float64(b.gameNode.H) || b.Y < 0 {
		prefab.DelPrefab(b.Id)
	}
	b.Move(b)

	if ebiten.IsDrawingSkipped() {
		return nil
	}

	opts2 := &ebiten.DrawImageOptions{}
	opts2.GeoM.Scale(b.ScaleW, b.ScaleH)
	opts2.GeoM.Translate(b.X, b.Y)
	b.gameNode.Screen.DrawImage(b.Img, opts2)
	return
}
