package core

import (
	"game_fly/core/component"
	"game_fly/core/prefab"
	"github.com/hajimehoshi/ebiten"
)

/*type SpriteComponent interface {
	OnLoad()
	Update(screen *ebiten.Image) (err error)
}

var (
	SpriteComponents = make(map[string]SpriteComponent)
	Prefab           = make(map[string]SpriteComponent)
)

func AddComponent(sprite SpriteComponent, name string) {
	sprite.OnLoad()
	SpriteComponents[name] = sprite
}

func GetComponent(name string) (component SpriteComponent) {
	return SpriteComponents[name]
}

func ComponentLen() (l int) {
	return len(SpriteComponents)
}*/

func GetComponentUpdate(screen *ebiten.Image) {
	for _, v := range component.SpriteComponents {
		v.Update(screen)
	}
	for _, v := range prefab.Prefab {
		v.Update(screen)
	}
}
