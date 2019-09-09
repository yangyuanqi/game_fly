package plays

import (
	"bytes"
	"game_fly/core"
	"game_fly/core/sprite"
	"game_fly/plays/assets/images"
	"github.com/hajimehoshi/ebiten"
	"image"
	_ "image/jpeg"
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
	//geo.Rect
	Screen *ebiten.Image //屏幕
	Map    *Map
}

func NewGame() (game *Game) {
	game = &Game{}
	game.W, game.H = 320, 480
	return
}

func init() {

}

func (g *Game) OnLoad() {
	//背景
	img1, _, err := image.Decode(bytes.NewReader(images.Map_jpg))
	if err != nil {
		log.Fatal(1, err)
	}
	MapImg, _ = ebiten.NewImageFromImage(img1, ebiten.FilterDefault)

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

	//m := &Map{}
	//m.Create()
	//sprite.AddSprite(m, "game")
	//
	role := &Role{}
	role.Create()
	sprite.AddSprite(role, "game")

	enemy := &Enemy{}
	sprite.AddSprite(enemy.Create(EnemyImg), "enemy")

	enemy2 := &Enemy{}
	enemy2 = enemy2.Create(EnemyImg)
	enemy2.SetXY(200, 150)
	sprite.AddSprite(enemy2, "enemy")

	enemy3 := &Enemy{}
	enemy3 = enemy3.Create(EnemyImg)
	enemy3.SetXY(150, 150)
	sprite.AddSprite(enemy3, "enemy")

}

func (g *Game) Update(screen *ebiten.Image) (err error) {
	g.Screen = screen
	return nil
}
