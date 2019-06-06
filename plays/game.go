package plays

import (
	"fmt"
	"github.com/SolarLune/resolv/resolv"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

var (
	Count int
)

type Game struct {
	W, H       int           //屏幕size
	screen     *ebiten.Image //屏幕
	scenesIng  int           //场景0初始化面，1第一关，2第二关，100结束画面
	hero       *Role         //英雄
	Enemys     []*Enemy      //敌军
	Input      *Input        //输入
	EnemySpace *resolv.Space //敌人碰撞空间
	Screen02   *Screen02
}

func NewGame() (game *Game) {
	game = &Game{}
	game.scenesIng = 1
	game.W, game.H = 360, 640
	game.EnemySpace = resolv.NewSpace()
	return
}

func (g *Game) Onload() {
	g.hero = NewRole(g)

	g.Input = NewInput(g)

	enemy := NewEnemy(g)
	enemy.Onload()

	g.Screen02 = NewScreen02(g)
	g.Screen02.Onload()

}

func (g *Game) Update(screen *ebiten.Image) (err error) {
	g.screen = screen
	Count ++
	g.Input.Update()
fmt.Println(len(g.Enemys))
	switch g.scenesIng {
	case 0:
	case 1:

		if Count == 1000 {
			Count = 0
		}

		//敌人
		if Count%50 == 0 {
			enemy := NewEnemy(g)
			g.Enemys = append(g.Enemys, enemy)

			g.EnemySpace.Add(enemy.Collide)
			g.EnemySpace.AddTags("enemy")
		}

		if ebiten.IsDrawingSkipped() {
			return nil
		}

		//渲染敌人
		for k, v := range g.Enemys {
			if v.Y > float64(g.H)+float64(g.Enemys[0].H) {
				if k < len(g.Enemys)-1 {
					g.Enemys = append(g.Enemys[:k], g.Enemys[k+1:]...)
				}
			}

			v.Update()
		}

		g.hero.Update()

		msg := fmt.Sprintf(`fps:%.2f`, ebiten.CurrentFPS())
		ebitenutil.DebugPrint(screen, msg)
	case 100:
		g.Screen02.Update()
	}

	return nil
}
