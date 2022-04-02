package core

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Component interface {
	Update(dt float64) error
	Draw(screen *ebiten.Image)
}

type GameObjectInterface interface {
	Init()
	Component
	GetChildren() *[]GameObjectInterface
	GetComponent() []Component
	GetDestroy() bool
}
