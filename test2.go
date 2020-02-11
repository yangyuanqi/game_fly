package main

import (
	"game_fly/core/component"
	"image"
)

type Enemy struct {
	component.GameObject
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
