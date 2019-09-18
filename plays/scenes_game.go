package plays

import (
	"bytes"
	"game_fly/core"
	"game_fly/core/sprite"
	"game_fly/core/ui"
	"game_fly/plays/assets/images"
	"github.com/hajimehoshi/ebiten"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"log"
)

var (
	BulletImg   *ebiten.Image
	BulletImg02 *ebiten.Image
	EnemyImg    *ebiten.Image
	EndImg      *ebiten.Image
	MapImg      *ebiten.Image
)

type Game struct {
	core.Sprite
	Map *Map
}

func NewGame() (game *Game) {
	game = &Game{}
	game.W, game.H = 320, 480
	return
}

//初始化场景的时候加载所有资源,可模拟资源加载动画，
func (g *Game) OnLoad() {
	//子弹贴图
	img, _, err := image.Decode(bytes.NewReader(images.Myb_1))
	if err != nil {
		log.Fatal(2, err)
	}
	BulletImg, _ = ebiten.NewImageFromImage(img, ebiten.FilterDefault)

	//子弹贴图
	img, _, err = image.Decode(bytes.NewReader(images.Epb_1))
	if err != nil {
		log.Fatal(3, err)
	}
	BulletImg02, _ = ebiten.NewImageFromImage(img, ebiten.FilterDefault)

	//敌人贴图
	img, _, err = image.Decode(bytes.NewReader(images.Ep_1))
	if err != nil {
		log.Fatal(4, err)
	}
	EnemyImg, _ = ebiten.NewImageFromImage(img, ebiten.FilterDefault)
	//结束贴图
	img, _, err = image.Decode(bytes.NewReader(images.Load))
	if err != nil {
		log.Fatal(5, err)
	}
	EndImg, _ = ebiten.NewImageFromImage(img, ebiten.FilterDefault)

	sprite.AddSprite(NewMap(), "game")

	sprite.AddSprite(NewRole(), "game")

	e1 := NewEnemy(EnemyImg)
	e1.SetXY(200, 150)
	sprite.AddSprite(e1, "enemy")

	//enemy2 := &Enemy{}
	//enemy2 = enemy2.Create(EnemyImg)
	//enemy2.SetXY(200, 150)
	//sprite.AddSprite(enemy2, "enemy")
	//
	//enemy3 := &Enemy{}
	//enemy3 = enemy3.Create(EnemyImg)
	//enemy3.SetXY(150, 150)
	//sprite.AddSprite(enemy3, "enemy")

	ui.AddUi(ui.NewNumber(0, 200, 10, 0.5), "ui")
}

func (g *Game) Update(screen *ebiten.Image) (err error) {

	return nil
}
