package scenes

import (
	"game_fly/assets/res/images"
	"game_fly/assets/scripts"
	"game_fly/core"
	"game_fly/core/component"
	"github.com/hajimehoshi/ebiten/v2"
)

var (
	My1  *ebiten.Image
	Map2 *ebiten.Image
)

type FlyScene struct {
	core.Scene
}

func (f *FlyScene) Init() {
	f.Name = "flyScene"
	My1 = core.Byte2Image(images.My_1)
	Map2 = core.Byte2Image(images.M2_jpg)

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
	roleNode.AddComponent(&component.Node{}, &component.Sprite{SpriteFrame: My1})
	f.AddChildren(bgNode, bgNode2, roleNode)
}

func (f *FlyScene) Update() error {
	return nil
}

func (f *FlyScene) Draw(screen *ebiten.Image) {

}
