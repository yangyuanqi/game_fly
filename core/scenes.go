package core

import (
	"github.com/hajimehoshi/ebiten"
)

type Scene interface {
	OnLoad()
	Update(screen *ebiten.Image) (err error)
}

var (
	Scenes    = make(map[string]interface{})
	ScenesIng string
)

func RegisterScene(scenes Scene, name string) {
	Scenes[name] = scenes
}

func GetScene(name string) (scene interface{}) {
	return Scenes[name]
}

func SetScenesIng(name string) {
	ScenesIng = name
	Scenes[name].(Scene).OnLoad()
}
