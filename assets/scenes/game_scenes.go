package scenes

import (
	"game_fly/assets/res/images"
	"game_fly/assets/scripts"
	"game_fly/core"
	"game_fly/core/component"
	"github.com/hajimehoshi/ebiten/v2"
	"time"
)

var (
	My1  *ebiten.Image
	Map2 *ebiten.Image
	Myb1 *ebiten.Image
)

type FlyScene struct {
	core.Scene
}

func (f *FlyScene) Init() {
	f.Name = "flyScene"
	My1 = core.Byte2Image(images.My_1)
	Map2 = core.Byte2Image(images.M2_jpg)
	Myb1 = core.Byte2Image(images.Myb_1)

	bgNode := &core.GameObject{Name: "background"}
	bgNode.AddComponent(
		&component.Node{Scale: component.Vec2{1, 1.05}},
		&component.Sprite{SpriteFrame: Map2},
		&scripts.Background{GameObject: bgNode},
	)
	bgNode2 := &core.GameObject{Name: "background2"}
	bgNode2.AddComponent(
		&component.Node{
			Position: component.Vec3{0, -800, 0},
			Scale:    component.Vec2{1, 1.05},
		},
		&component.Sprite{SpriteFrame: Map2},
		&scripts.Background{GameObject: bgNode2},
	)

	roleNode := &core.GameObject{Name: "role"}
	roleNode.AddComponent(
		&component.Node{
			Position: component.Vec3{100, 500, 0},
			Scale:    component.Vec2{1, 1},
		},
		&component.Sprite{SpriteFrame: My1},
		&scripts.Role{GameObject: roleNode},
	)
	f.AddChildren(bgNode, bgNode2, roleNode)

	core.SetTicker(time.Millisecond*200, func() {
		bulletNode := &core.GameObject{Name: "bullet"}
		bulletNode.AddComponent(
			&component.Node{
				Position: component.Vec3{roleNode.Node().Position.X + 55, roleNode.Node().Position.Y - 50, 0},
				Scale:    component.Vec2{1, 1},
			},
			&component.Sprite{SpriteFrame: Myb1},
			&scripts.Bullet{GameObject: bulletNode},
		)
		f.AddChildren(bulletNode)
	}, 0)
}

func (f *FlyScene) Update(dt float64) error {
	return nil
}

func (f *FlyScene) Draw(screen *ebiten.Image) {

}
