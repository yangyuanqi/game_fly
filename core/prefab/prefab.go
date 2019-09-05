package prefab

import (
	"game_fly/core/component"
)

var Prefab = make(map[string]component.SpriteComponent)

func AddPrefab(prefab component.SpriteComponent, name string) {
	prefab.OnLoad()
	Prefab[name] = prefab
}

func GetPrefab(name string) (prefab component.SpriteComponent) {
	return Prefab[name]
}

func PrefabLen() (l int) {
	return len(Prefab)
}
func DelPrefab(id string) {
	delete(Prefab, id)
}
