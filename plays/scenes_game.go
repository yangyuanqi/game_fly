package plays

import (
	"bytes"
	"game_fly/core"
	"game_fly/core/component"
	"game_fly/plays/assets/images"
	"github.com/hajimehoshi/ebiten"
	"image"
	"log"
)

var (
	BulletImg   *ebiten.Image
	BulletImg02 *ebiten.Image
	EnemyImg    *ebiten.Image
	EndImg      *ebiten.Image
)

type Game struct {
	core.Sprite
	//geo.Rect
	Screen  *ebiten.Image //屏幕
}

func NewGame() (game *Game) {
	game = &Game{}
	game.W, game.H = 320, 480
	return
}

func (g *Game) OnLoad() {
	//子弹贴图
	img, _, err := image.Decode(bytes.NewReader(images.Myb_1))
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

	component.AddComponent(NewRole(), "role")

}

func (g *Game) Update(screen *ebiten.Image) (err error) {
	g.Screen = screen
	core.GetComponentUpdate(screen)

	//for _, v := range g.SpriteComponent {
	//	v.Update(screen)
	//}

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
