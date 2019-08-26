package core

import "github.com/hajimehoshi/ebiten"

type SpriteComponent interface {
	OnLoad()
	Update(screen *ebiten.Image) (err error)
}

var (
	SpriteComponents = make(map[string]SpriteComponent)
)

func AddSprite(sprite SpriteComponent, name string) {
	sprite.OnLoad()
	SpriteComponents[name] = sprite
}

func GetComponent(name string) (component SpriteComponent) {
	return SpriteComponents[name]
}

func GetComponentUpdate(screen *ebiten.Image) {
	for _, v := range SpriteComponents {
		v.Update(screen)
	}
}
