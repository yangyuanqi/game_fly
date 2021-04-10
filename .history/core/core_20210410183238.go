package core

import "github.com/hajimehoshi/ebiten"

type ObjectGame interface {
	OnLoad()
	Update()
}

type Core struct {
	objectGame ObjectGame
}

func (c *Core) OnLoad() {

}

func (c *Core) Update(screen *ebiten.Image) (err error) {
	return nil
}
