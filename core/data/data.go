package data

import (
	"github.com/hajimehoshi/ebiten"
	"sync"
)

var PrefabL sync.Mutex

type SpriteGroup struct {
	GroupName string
	Sprite    []SpriteComponent //有顺序要求
	Prefab map[string]SpriteComponent
	Ui     []UiType
}

type UiType interface {
	OnLoad()
	Start()
	Update(screen *ebiten.Image) (err error)
	GetName() (name string)
}

type SpriteComponent interface {
	OnLoad()
	Start()
	Update(screen *ebiten.Image) (err error)
	GetResolv() (collision interface{})
	UpdateResolv()
	GetId() (id string)
	GetName() (name string)
	//SetScreen(image *ebiten.Image)
	GetComponent() (components []SpriteComponent)
	GetVisible() (b bool)
	SetVisible(b bool)
	GetDestroy() (b bool)
	SetDestroy(b bool)
	GetPosition() (x, y, w, h float64)
}

type Core struct {
	Width, Height int
	Scale         float64
	title         string
}

func (c *Core) GetWH() (w, h int) {
	return c.Width, c.Height
}
