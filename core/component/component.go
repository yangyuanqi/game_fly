package component

import (
	"github.com/hajimehoshi/ebiten"
)

type SpriteComponent interface {
	OnLoad()
	Start()
	Update(screen *ebiten.Image) (err error)
	GetResolv() (collision interface{})
	UpdateResolv()
	GetId() (id string)
	GetName()(name string)
	GetComponent() (components []SpriteComponent)
}

var (
	SpriteComponents = make(map[string]SpriteComponent)
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
}
