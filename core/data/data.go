package data

import (
	"game_fly/core/component"
	"github.com/hajimehoshi/ebiten"
)

type SpriteGroup struct {
	GroupName string
	Sprite    []component.SpriteComponent //有顺序要求
	Prefab    map[string]component.SpriteComponent
	Ui        []UiType
}

type UiType interface {
	OnLoad()
	Start()
	Update(screen *ebiten.Image) (err error)
	GetName() (name string)
}