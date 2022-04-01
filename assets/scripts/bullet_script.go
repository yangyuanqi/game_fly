package scripts

import (
	"game_fly/core"
	"game_fly/core/component"
)

type Bullet struct {
	*core.GameObject
}

func (b *Bullet) Update(dt float64) error {
	b.SetPosition(component.Vec2{b.Node().Position.X, b.Node().Position.Y - 500*dt})
	if b.Node().Position.Y <= -50 {
		b.Destroy()
	}
	return nil
}
