package plays

import (
	"github.com/hajimehoshi/ebiten"
)

var (
	BulletImg *ebiten.Image
	EnemyImg  *ebiten.Image
	RoleImg   *ebiten.Image
	MapImg    *ebiten.Image
	StarImg   *ebiten.Image
)

//type Game struct {
//	data.Sprite
//	danyao    map[string]data.SpriteComponent
//	roleNode  *Role
//	StarSpace *resolv.Space
//}

//func NewGame() (game *Game) {
//	game = &Game{}
//	return
//}

//初始化场景的时候加载所有资源,可模拟资源加载动画，
//func (g *Game) OnLoad() {
//	fmt.Println(time.Now().UnixNano())
//	g.W, g.H = int32(conf.GetConfInt("scenes_width")), int32(conf.GetConfInt("scenes_height"))
//
//	BulletImg = core.Byte2Image(images.Myb_1)
//	EnemyImg = core.Byte2Image(images.Ep_1)
//	RoleImg = core.Byte2Image(images.My_1)
//	MapImg = core.Byte2Image(images.M2_jpg)
//	StarImg = core.Byte2Image(images.Xx_png)
//	fmt.Println(time.Now().UnixNano())
//}

//func (g *Game) Start() () {
//	//sprite.AddSprite(&Map{}, "game")
//	//sprite.AddSprite(NewRole(), "game")
//	//
//	//ui.AddUi(ui.NewNumber(0, 200, 10, 15), "ui")
//
//	//g.roleNode = sprite.GetSprite("game", "role").(*Role)
//
//	g.StarSpace = resolv.NewSpace()
//	g.StarSpace.AddTags("star")
//
//}
//
//func (g *Game) Update(screen *ebiten.Image) (err error) {
//	g.UpdateTiming()
//
//	if g.Timing%conf.GetConfInt("enemy_speed") == 0 {
//		//g.autoEnemy()
//	}
//
//	if g.Timing%100 == 0 {
//		g.createStar()
//	}
//
//	if g.StarSpace.IsColliding(g.roleNode) {
//		for _, v := range *g.StarSpace.GetCollidingShapes(g.roleNode) {
//			//prefab.DelPrefab(v.(*Star).GetId())
//			g.StarSpace.Remove(v)
//			g.roleNode.danyao = 1
//		}
//	}
//
//	return nil
//}
//
////func (g *Game) autoEnemy() {
////	rand.Seed(time.Now().UnixNano())
////	x := rand.Intn(260)
////
////	e1 := NewEnemy(EnemyImg)
////	e1.SetXY(int32(x), -60)
////	e1.SetFSM(EnemyMove, 1)
////	prefab.AddPrefab(e1, "enemy")
////}
//
//func (g *Game) createStar() {
//	//rand.Seed(time.Now().UnixNano())
//	//x := rand.Intn(300)
//	//s := NewStar()
//	//s.SetXY(100, -35)
//	//s.SetFSM(EnemyMove, 1)
//	//prefab.AddPrefab(s, "danyao")
//
//	//g.StarSpace.Add(s)
//}
