package plays

import (
	"game_fly/core"
	"game_fly/core/prefab"
	"game_fly/core/sprite"
	"game_fly/plays/conf"
	"github.com/SolarLune/resolv/resolv"
	"github.com/hajimehoshi/ebiten"
)

type Enemy struct {
	core.Sprite
	SouShang int
}

const (
	SouShang = iota
	EnemyMove
)

func NewEnemy(img *ebiten.Image) (enemy *Enemy) {
	e := &Enemy{}
	e.Sprite.Create("enemy")
	return e
}

//加载资源
func (e *Enemy) OnLoad() {
	e.SetMaterial(EnemyImg)
	e.SetScale(0.3, 0.3)
	e.Collision = resolv.NewRectangle(e.GetPositionInt32())
}

func (e *Enemy) Start() {

}

//更新
func (e *Enemy) Update(screen *ebiten.Image) (err error) {
	if ebiten.IsDrawingSkipped() {
		return nil
	}

	if e.GetFSM(SouShang) >= 1 {
		role := sprite.GetSprite("game", "role").(*Role)
		role.SetFSM(Fraction, role.GetFSM(Fraction)+1)
		//prefab.DelPrefab("enemy", e.GetId())
		e.Destroy = true
	}

	sh := conf.GetConfInt("scenes_height")
	if e.Y > int32(sh) {
		prefab.DelPrefab(e.Id)
	}

	enemyMove := e.GetFSM(EnemyMove)
	if enemyMove == 1 {
		e.Move(0, 1)
		//if e.Y > 50 {
		//e.Move(0, math.Sin(e.Y))
		//}

	}

	e.Draw(screen)
	return nil
}

//渲染
func (e *Enemy) Draw(screen *ebiten.Image) {

	e.UpdateTiming()
	if e.Timing%conf.GetConfInt("enemy_bullet_speed") == 0 {
		e.autoBullet()
	}

	opts2 := &ebiten.DrawImageOptions{}
	opts2.GeoM.Scale(e.ScaleW, e.ScaleH)
	opts2.GeoM.Translate(float64(e.X), float64(e.Y))
	screen.DrawImage(e.Material, opts2)
}

func (e *Enemy) autoBullet() {
	bullet := NewBullet(BulletImg, e)
	bullet.SetFSM(BulletType, 2)
	prefab.AddPrefab(bullet, "enemyBullet")
}
