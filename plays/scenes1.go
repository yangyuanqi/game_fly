package plays

import (
	"fmt"
	"game_fly/core"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/inpututil"
)

type Scenes1 struct {
	core.BaseComponent
	Hero  *Hero
	Stone *Stone
}

func (s *Scenes1) OnLoad() {
	//core.RegisterComponent(&Scenes1{})
	/*	core.SetTicker(5*time.Second, func() {
			core.SetScenesIng("scen2")
		}, 1)*/
	hero := &Hero{}
	core.RegisterComponent(hero)
	s.Hero = hero

	stone := &Stone{}
	core.RegisterComponent(stone)
	s.Stone = stone
}

func (s *Scenes1) Update(screen *ebiten.Image) {

	msg := fmt.Sprintf(`fps:%.2f`, ebiten.CurrentFPS())
	ebitenutil.DebugPrint(screen, msg)

	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		if v := inpututil.KeyPressDuration(ebiten.KeyUp); 0 < v {
			if v == 1 || v%1 == 0 {
				s.Hero.Y -= 2
			}
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		if v := inpututil.KeyPressDuration(ebiten.KeyDown); 0 < v {
			if v == 1 || v%1 == 0 {
				s.Hero.Y += 2
			}
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		if v := inpututil.KeyPressDuration(ebiten.KeyLeft); 0 < v {
			if v == 1 || v%1 == 0 {
				s.Hero.X -= 2
			}
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		if v := inpututil.KeyPressDuration(ebiten.KeyRight); 0 < v {
			if v == 1 || v%1 == 0 {
				s.Hero.X += 2
			}
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		if v := inpututil.KeyPressDuration(ebiten.KeySpace); 0 < v {
			if v == 1 {
				s.Hero.Jump = 1
			}
		}
	}

	for _, v := range core.GetComponents() {
		v.Update(screen)
	}

	b := s.Hero.CollideRect(s.Stone.Rect)
	if b {
		fmt.Println(s.Hero.Rect)
	}

}
