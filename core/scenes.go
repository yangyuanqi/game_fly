package core

import (
	"github.com/hajimehoshi/ebiten"
)

type Scene interface {
	OnLoad()
	Update(screen *ebiten.Image)
}

var (
	Scenes    = make(map[string]Scene)
	ScenesIng string
)

func RegisterScene(scenes Scene, name string) {
	scenes.OnLoad()
	Scenes[name] = scenes
}

func GetScene(name string) (scene Scene) {
	return Scenes[name]
}

func SetScenesIng(name string) {
	ScenesIng = name
}
