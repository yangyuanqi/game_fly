package core

import "github.com/hajimehoshi/ebiten"

type Component interface {
	OnLoad()
	Update(screen *ebiten.Image)
}

var Components []Component

func RegisterComponent(component Component) {
	component.OnLoad()
	Components = append(Components, component)
}

func GetComponents() (components []Component) {
	return Components
}
