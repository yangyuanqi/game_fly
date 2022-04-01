package scripts

import (
	"game_fly/core"
	"game_fly/core/component"
)

type Background struct {
	*core.GameObject
}

func (f *Background) Update(dt float64) error {
	f.SetPosition(component.Vec2{f.Node().Position.X, f.Node().Position.Y + 1})
	if f.Node().Position.Y >= 800 {
		f.Node().Position.Y = -800
	}
	return nil
}
