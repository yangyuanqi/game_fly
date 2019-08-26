package plays

import (
	"bytes"
	"game_fly/core"
	"game_fly/plays/assets/images"
	"github.com/hajimehoshi/ebiten"
	"image"
	"log"
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
	game.W, game.H = 320, 480
	return
}

func (g *Game) OnLoad() {
	g.hero = NewRole(g, RoleImg, BulletImg)
	core.AddSprite(g.hero, "role")

	core.AddSprite(NewInput(g), "input")

	//
	//g.hero = NewRole(g, RoleImg, BulletImg)
	//g.hero.OnLoad()
	//g.Input = NewInput(g)
	//
	//params := Enemy{}
	//params.BulletStatus = 0
	//g.Enemy = NewEnemy(g, EnemyImg, BulletImg02, params)
	//g.Enemy.Onload()
	//
	//core.SetTimer(time.Second*1, func() {
	//	core.SetTicker(time.Millisecond*500, g.enemy01, 5)
	//})
	//core.SetTimer(time.Second*1, func() {
	//	core.SetTicker(time.Millisecond*500, g.enemy02, 5)
	//})
	//core.SetTimer(time.Second*6, func() {
	//	core.SetTicker(time.Millisecond*500, g.enemy01, 5)
	//})
	//core.SetTimer(time.Second*10, func() {
	//	core.SetTicker(time.Millisecond*500, g.enemy02, 5)
	//})
	//core.SetTicker(time.Millisecond*500, g.enemy02, 0)
}

func (g *Game) Update(screen *ebiten.Image) (err error) {
	g.Screen = screen
	core.GetComponentUpdate(screen)

	//g.Input.Update()
	//
	//switch g.scenesIng {
	//case 0:
	//case 1:
	//	if ebiten.IsDrawingSkipped() {
	//		return nil
	//	}
	//
	//	//渲染敌人
	//	for _, v := range g.Enemys {
	//		v.Update()
	//	}
	//
	//	g.hero.Update()
	//
	//	msg := fmt.Sprintf(`fps:%.2f`, ebiten.CurrentFPS())
	//	ebitenutil.DebugPrint(screen, msg)
	//
	//	fmt.Println(len(g.hero.GroupBullet))
	//case 100:
	//	op := &ebiten.DrawImageOptions{}
	//	screen.DrawImage(EndImg, op)
	//}
	//
	////销毁
	//for k, v := range g.Enemys {
	//	if v.Visible == false {
	//		if k < len(g.Enemys) {
	//			g.Enemys = append(g.Enemys[:k], g.Enemys[k+1:]...)
	//		}
	//	}
	//}
	//for k, v := range g.hero.GroupBullet {
	//	if v.Visible == false {
	//		if k < len(g.hero.GroupBullet) {
	//			g.hero.GroupBullet = append(g.hero.GroupBullet[:k], g.hero.GroupBullet[k+1:]...)
	//		}
	//	}
	//}
	//
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
