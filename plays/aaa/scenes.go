package aaa

import "game_fly/core/component"

//type Scene interface {
//	OnLoad()
//	Start()
//	Update(screen *ebiten.Image) (err error)
//}

var Game *component.Scene

var (
	Scenes    = make(map[string]SpriteComponent)
	ScenesIng string
)

func RegisterScene(scenes SpriteComponent, name string) {
	Scenes[name] = scenes
	scenes.OnLoad()
	scenes.Start()
	for _, v := range scenes.GetPrefab() {
		v.OnLoad()
		v.Start()
	}
	GetComponentStart(scenes)
}

func GetScene(name string) (scene SpriteComponent) {
	return Scenes[name]
}

func LoadScene(name string) {
	ScenesIng = name
}
