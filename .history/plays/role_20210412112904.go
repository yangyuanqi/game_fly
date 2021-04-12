package plays

import (
	"game_fly/core"
	"game_fly/plays/assets/images"

	"github.com/hajimehoshi/ebiten"
)

type Role struct {
	core.Sprite
	bullet *Bullet
}

func NewRole() *Role {
	return &Role{}

}

func (r *Role) OnLoad() {
	r.SetMaterial(core.Byte2Image(images.My_1))
	r.SetScale(0.3, 0.3)

	r.SetXY(100, 100)
	r.SetXY(1, 1)

	var a = &core.Animator{
		Sprite: core.Sprite{
			Base: core.Base{
				X: 300,
				Y: 300,
			},
		},
		Interval: 5,
		Order:    "line",
		Imgs:     []*ebiten.Image{core.Nodes.Img["t1png"], core.Nodes.Img["t2png"]},
	}
	core.Nodes.AddPrefab(a)
}

func (r *Role) Update(screen *ebiten.Image) (err error) {
	r.X = r.X + int32(r.KeyX()*3)
	r.Y = r.Y + int32(r.KeyY()*3)

	if r.KeySpace() {
		r.bullet = NewBullet()
		r.bullet.SetXY(r.X, r.Y)
		core.Nodes.AddPrefab(r.bullet)
	}
	r.Sprite.Update(screen)
	return
}

//
//import (
//	"game_fly/core/data"
//	"game_fly/plays/conf"
//	"github.com/SolarLune/resolv/resolv"
//	"github.com/hajimehoshi/ebiten"
//)
//
//type Role struct {
//	data.Sprite
//	//RootNode     *Game
//	danyao       int
//	QingXieLeft  bool
//	QingXieRight bool
//}
//
//const (
//	StartAnimation = iota //出场动画
//	Fraction              //分数
//	Left                  //左右移动动画
//	Right
//)
//
//func NewRole() (role *Role) {
//	r := &Role{}
//	r.Sprite.Create("role")
//	r.SetFSM(StartAnimation, 0)
//	return r
//}
//
//func (r *Role) OnLoad() {
//	//角色贴图
//	r.SetMaterial(RoleImg)
//	//game := core.GetScene("game").(*Game)
//	//r.RootNode = game
//	r.SetScale(0.3, 0.3)
//	//r.SetXY(int32(math.Ceil(float64((game.W-r.W)/2))), game.H)
//	r.Collision = resolv.NewRectangle(r.GetPositionInt32())
//
//	//r.AddComponent(NewInput())
//}
//
//func (r *Role) Start() {
//
//}
//
//func (r *Role) Update(screen *ebiten.Image) (err error) {
//	//ui.ReplaceUi("ui", "number1", ui.NewNumber(r.GetFSM(Fraction), 200, 10, 15))
//	r.Drow(screen)
//	return nil
//}
//
//var danyaoTime int = 500
//var QingXieTime float64 = 0
//
//func (r *Role) Drow(screen *ebiten.Image) {
//	r.UpdateTiming()
//	if r.Timing%conf.GetConfInt("role_bullet_speed") == 0 {
//		//r.autoBullet()
//	}
//
//	if r.danyao == 1 {
//		danyaoTime--
//	}
//	if danyaoTime < 0 {
//		r.danyao = 0
//		danyaoTime = 500
//	}
//
//	//opts := &ebiten.DrawImageOptions{}
//	//opts.GeoM.Reset()
//	//opts.GeoM.Scale(r.ScaleW, r.ScaleH)
//
//	//飞机初始动画
//	//if r.GetFSM(StartAnimation) == 0 {
//	//	r.Move(0, -2)
//	//	if r.Y < r.RootNode.H-r.H {
//	//		r.SetFSM(StartAnimation, 1)
//	//	}
//	//}
//
//	//灰机倾斜处理
//	if r.QingXieLeft {
//		QingXieTime--
//		if QingXieTime > -20 {
//			r.SetSkewXY(0, QingXieTime/100)
//		} else {
//			QingXieTime = -20
//		}
//	}
//	if r.QingXieRight {
//		QingXieTime++
//		if QingXieTime < 20 {
//			r.SetSkewXY(0, QingXieTime/100)
//		} else {
//			QingXieTime = 20
//		}
//	}
//
//	if r.QingXieRight == false && r.QingXieLeft == false {
//		if QingXieTime < 0 {
//			QingXieTime++
//			r.SetSkewXY(0, QingXieTime/100)
//		}
//		if QingXieTime > 0 {
//			QingXieTime--
//			r.SetSkewXY(0, QingXieTime/100)
//		}
//	}
//
//	//opts.GeoM.Translate(float64(r.X), float64(r.Y))
//	//screen.DrawImage(r.Material, opts)
//	r.Sprite.Update(screen)
//}
//
////func (r *Role) autoBullet() {
////	if r.danyao == 0 {
////		bullet := NewBullet(BulletImg, r)
////		bullet.SetFSM(BulletType, 1)
////		prefab.AddPrefab(bullet, "roleBullet")
////	}
////	if r.danyao == 1 {
////		bullet := NewBullet(BulletImg, r)
////		bullet.SetFSM(BulletType, 1)
////		prefab.AddPrefab(bullet, "roleBullet")
////		bullet.X = bullet.X + 10
////
////		bullet2 := NewBullet(BulletImg, r)
////		bullet2.SetFSM(BulletType, 1)
////		prefab.AddPrefab(bullet2, "roleBullet")
////		bullet2.X = bullet2.X - 10
////	}
////}
