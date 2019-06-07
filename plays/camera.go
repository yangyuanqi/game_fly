package plays

import (
	"bytes"
	"fmt"
	"game_fly/core"
	"game_fly/plays/assets/images"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"image"
	"log"
	"time"
)

var (
	RoleImg     *ebiten.Image
	BulletImg   *ebiten.Image
	BulletImg02 *ebiten.Image
	EnemyImg    *ebiten.Image
	EndImg      *ebiten.Image
)

type Game struct {
	W, H      int           //屏幕size
	Screen    *ebiten.Image //屏幕
	scenesIng int           //场景0初始化面，1第一关，2第二关，100结束画面
	hero      *Role         //英雄
	Enemy     *Enemy
	Enemys    []*Enemy //敌军
	Input     *Input   //输入
}

func init() {
	//角色贴图
	img, _, err := image.Decode(bytes.NewReader(images.My_1))
	if err != nil {
		log.Fatal(err)
	}
	RoleImg, _ = ebiten.NewImageFromImage(img, ebiten.FilterDefault)
	//子弹贴图
	img, _, err = image.Decode(bytes.NewReader(images.Myb_1))
	if err != nil {
		log.Fatal(err)
	}
	BulletImg, _ = ebiten.NewImageFromImage(img, ebiten.FilterDefault)

	//子弹贴图
	img, _, err = image.Decode(bytes.NewReader(images.Epb_1))
	if err != nil {
		log.Fatal(err)
	}
	BulletImg02, _ = ebiten.NewImageFromImage(img, ebiten.FilterDefault)

	//敌人贴图
	img, _, err = image.Decode(bytes.NewReader(images.Ep_1))
	if err != nil {
		log.Fatal(err)
	}
	EnemyImg, _ = ebiten.NewImageFromImage(img, ebiten.FilterDefault)

	//结束贴图
	img, _, err = image.Decode(bytes.NewReader(images.Load))
	if err != nil {
		log.Fatal(err)
	}
	EndImg, _ = ebiten.NewImageFromImage(img, ebiten.FilterDefault)
}

func NewGame() (game *Game) {
	game = &Game{}
	game.scenesIng = 1
	game.W, game.H = 360, 640
	return
}

func (g *Game) Onload() {
	g.hero = NewRole(g, RoleImg, BulletImg)
	g.hero.Onload()

	g.Input = NewInput(g)

	params := Enemy{}
	params.BulletStatus = 0
	g.Enemy = NewEnemy(g, EnemyImg, BulletImg02, params)
	g.Enemy.Onload()
	//
	core.SetTimer(time.Second*1, func() {
		core.SetTicker(time.Millisecond*500, g.enemy01, 5)
	})
	core.SetTimer(time.Second*1, func() {
		core.SetTicker(time.Millisecond*500, g.enemy02, 5)
	})
	core.SetTimer(time.Second*6, func() {
		core.SetTicker(time.Millisecond*500, g.enemy01, 5)
	})
	core.SetTimer(time.Second*10, func() {
		core.SetTicker(time.Millisecond*500, g.enemy02, 5)
	})
	core.SetTicker(time.Millisecond*500, g.enemy02, 0)
}

func (g *Game) Update(screen *ebiten.Image) (err error) {
	g.Screen = screen
	g.Input.Update()

	switch g.scenesIng {
	case 0:
	case 1:
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
		op := &ebiten.DrawImageOptions{}
		screen.DrawImage(EndImg, op)
	}

	return nil
}

//第一波敌人
func (g *Game) enemy01() {
	//敌人
	var enemy *Enemy
	enemy = NewEnemy(g, EnemyImg, BulletImg02, Enemy{BulletStatus: 1})
	enemy.Onload()
	enemy.MY = 4
	enemy.MX = 3
	g.Enemys = append(g.Enemys, enemy)
}
func (g *Game) enemy02() {
	//敌人
	var enemy *Enemy
	enemy = NewEnemy(g, EnemyImg, BulletImg02, Enemy{BulletStatus: 1})
	enemy.Onload()
	enemy.X = float64(g.W - g.Enemy.W)
	enemy.Y = 0
	enemy.MY = 4
	enemy.MX = -3
	g.Enemys = append(g.Enemys, enemy)
}
