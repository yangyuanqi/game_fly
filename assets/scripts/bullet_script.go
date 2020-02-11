package scripts

import (
	"game_fly/assets/prefabs"
	"game_fly/core/component"
)

type Bullet struct {
	component.GameObject
	prefab prefabs.Bullet
}