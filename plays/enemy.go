package plays

import (
	"game_fly/core"
	"github.com/hajimehoshi/ebiten"
	"math"
	"time"
)

type Enemy struct {
	core.Sprite
	GroupBillet  []*Bullet
	BulletStatus int //给与敌人什么子弹
	RootNode     *Game
	EnemyImg     *ebiten.Image
	BulletImg    *ebiten.Image
}

//初始化
func NewEnemy(rootNode *Game, enemyImg, bulletImg *ebiten.Image, params Enemy) (enemy *Enemy) {
	//基础属性
	enemy = &Enemy{}
	enemy.RootNode = rootNode
	enemy.EnemyImg = enemyImg
	enemy.BulletImg = bulletImg

	w, h := enemyImg.Size()
	enemy.SetScale(0.3, 0.3)
	enemy.SetWH(int(math.Round(float64(w)*enemy.ScaleW)), int(math.Round(float64(h)*enemy.ScaleH)))
	enemy.SetXY(0, 0)

	//自定义属性
	enemy.BulletStatus = params.BulletStatus
	return
}

//加载资源
func (e *Enemy) Onload() {
	core.SetTicker(time.Millisecond*200, e.bullet01, 0)
}

//更新
func (e *Enemy) Update() (err error) {

	for k, v := range e.GroupBillet {
		//到达边界删除
		if v.Y >= float64(e.RootNode.H) {
			if k < len(e.GroupBillet) {
				e.GroupBillet = append(e.GroupBillet[:k], e.GroupBillet[k+1:]...)
			}
		}
		if core.CheckCollision(v, e.RootNode.hero) {
			e.RootNode.scenesIng = 100
		}
		v.Update()
	}

	e.Draw()
	return nil
}

//渲染
func (e *Enemy) Draw() {
	e.Move()

	opts2 := &ebiten.DrawImageOptions{}
	opts2.GeoM.Scale(e.ScaleW, e.ScaleH)
	opts2.GeoM.Translate(e.X, e.Y)
	e.RootNode.Screen.DrawImage(e.EnemyImg, opts2)
}

func (e *Enemy) bullet01() {
	b := NewBullet(e.RootNode, e, e.BulletImg)
	b.MY = 10
	e.GroupBillet = append(e.GroupBillet, b)
}
