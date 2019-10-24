package data

import (
	"github.com/hajimehoshi/ebiten"
	"sync"
)

var PrefabL sync.Mutex
var ID int

//type SpriteGroup struct {
//	GroupName string
//	Sprite    []SpriteComponent //有顺序要求
//	Prefab    map[string]SpriteComponent
//	Ui        []UiType
//}

type UiType interface {
	OnLoad()
	Start()
	Update(screen *ebiten.Image) (err error)
	GetName() (name string)
}



type Core struct {
	Width, Height int
	Scale         float64
	title         string
}

func (c *Core) GetWH() (w, h int) {
	return c.Width, c.Height
}

type Prefab struct {
	Sprite
}
