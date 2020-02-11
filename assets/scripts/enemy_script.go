package scripts

import (
	"game_fly/assets/prefabs"
	"game_fly/core/component"
	"image"
)

type Enemy struct {
	component.GameObject
	prefab prefabs.Bullet
}

func NewEnemy() (enemy *Enemy) {
	enemy = &Enemy{}
	script := &component.Sprite{}
	script.Color = image.Black

	enemy.RegisterComponent(script)
	return
}

func (e *Enemy) Start() {

}