package scripts

import (
	"game_fly/core"
	"game_fly/core/component"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Role struct {
	*core.GameObject
}

func (r *Role) Update(dt float64) error {
	if d := inpututil.KeyPressDuration(ebiten.KeyUp); d > 0 {
		r.SetPosition(component.Vec2{r.Node().Position.X, r.Node().Position.Y - 500*dt})
	}
	if d := inpututil.KeyPressDuration(ebiten.KeyDown); d > 0 {
		r.SetPosition(component.Vec2{r.Node().Position.X, r.Node().Position.Y + 500*dt})
	}
	if d := inpututil.KeyPressDuration(ebiten.KeyLeft); d > 0 {
		r.SetPosition(component.Vec2{r.Node().Position.X - 500*dt, r.Node().Position.Y})
	}
	if d := inpututil.KeyPressDuration(ebiten.KeyRight); d > 0 {
		r.SetPosition(component.Vec2{r.Node().Position.X + 500*dt, r.Node().Position.Y})
	}
	return nil
}

func repeatingKeyPressed(key ebiten.Key) bool {
	const (
		delay    = 30
		interval = 3
	)
	d := inpututil.KeyPressDuration(key)
	if d == 1 {
		return true
	}
	if d >= delay && (d-delay)%interval == 0 {
		return true
	}
	return false
}
