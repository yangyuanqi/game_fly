package plays

import (
	"github.com/SolarLune/resolv/resolv"
	"github.com/hajimehoshi/ebiten"
	"image/color"
	"math/rand"
	"time"
)

type Enemy struct {
	Billet []*Bullet
	CC
}

var billetCount int
var enemyImg *ebiten.Image
//初始化
func NewEnemy(game *Game) (enemy *Enemy) {
	enemy = &Enemy{}
	enemy.RootNode = game
	enemy.W = 20
	enemy.H = 20
	rand.Seed(time.Now().UnixNano())
	enemy.X = float64(rand.Intn(enemy.RootNode.W - enemy.W))
	enemy.Y = 0
	enemy.Collide = resolv.NewRectangle(int32(enemy.X), int32(enemy.Y), int32(enemy.W), int32(enemy.H))
	enemy.CollideSpace = resolv.NewSpace()
	return
}

//加载资源
func (e *Enemy) Onload() {
	enemyImg, _ = ebiten.NewImage(e.W, e.H, ebiten.FilterNearest)
	enemyImg.Fill(color.White)

	bullet := NewBullet(e.RootNode, e, 0)
	bullet.Onload()

}

//更新
func (e *Enemy) Update() (err error) {
	billetCount++
	if billetCount == 1000 {
		billetCount = 0
	}

	//每2个单位生成一个子弹
	if billetCount%90 == 0 {
		b := NewBullet(e.RootNode, e, 0)
		e.Billet = append(e.Billet, b)

		e.Collide.AddTags("enemyBullet")
		e.RootNode.EnemySpace.Add(e.Collide)
	}

	if len(e.Billet) > 0 {
		for k, v := range e.Billet {
			v.Update("enemy")
			//到达边界删除
			if v.Y >= float64(e.RootNode.H) {
				e.Billet = append(e.Billet[:k], e.Billet[k+1:]...)
			}
		}
	}

	e.Draw()
	return nil
}

//渲染
func (e *Enemy) Draw() {
	e.Move(0, 2)

	opts2 := &ebiten.DrawImageOptions{}
	opts2.GeoM.Translate(e.X, e.Y)
	e.RootNode.screen.DrawImage(enemyImg, opts2)
}

func (e *Enemy) GetPosition() (x, y float64, w, h int) {
	return e.X, e.Y, e.W, e.H
}
