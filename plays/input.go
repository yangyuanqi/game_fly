package plays

import (
	"fmt"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
)

type Input struct {
	RootNode *Game
}

func NewInput(game *Game) (input *Input) {
	input = &Input{}
	input.RootNode = game
	return
}

func (i *Input) Update() (err error) {
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		if v := inpututil.KeyPressDuration(ebiten.KeyUp); 0 < v {
			if v == 1 || v%2 == 0 {
				i.RootNode.hero.Y -= 15
				i.RootNode.hero.Collide.Y -= 15
			}
		}
		if i.RootNode.hero.Y <= 0 {
			i.RootNode.hero.Y = 0
			i.RootNode.hero.Collide.Y = 0
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		fmt.Println(i.RootNode.hero.X,i.RootNode.hero.Y)
		if v := inpututil.KeyPressDuration(ebiten.KeyDown); 0 < v {
			if v == 1 || v%2 == 0 {
				i.RootNode.hero.Y += 15
				i.RootNode.hero.Collide.Y += 15
			}
		}
				if i.RootNode.hero.Y >= float64(i.RootNode.H)-float64(i.RootNode.hero.H) {
					i.RootNode.hero.Y = float64(i.RootNode.H) - float64(i.RootNode.hero.H)
					i.RootNode.hero.Collide.Y = int32(i.RootNode.H) - int32(i.RootNode.hero.H)
				}
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		if v := inpututil.KeyPressDuration(ebiten.KeyLeft); 0 < v {
			if v == 1 || v%2 == 0 {
				i.RootNode.hero.X -= 15
				i.RootNode.hero.Collide.X -= 15
			}
		}
		if i.RootNode.hero.X <= 0 {
			i.RootNode.hero.X = 0
			i.RootNode.hero.Collide.X = 0
		}
	}
	// 當「按鍵右」被按下時⋯⋯
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		if v := inpututil.KeyPressDuration(ebiten.KeyRight); 0 < v {
			if v == 1 || v%2 == 0 {
				i.RootNode.hero.X += 15
				i.RootNode.hero.Collide.X += 15
			}
		}
		if i.RootNode.hero.X >= float64(i.RootNode.W)-float64(i.RootNode.hero.W) {
			i.RootNode.hero.X = float64(i.RootNode.W) - float64(i.RootNode.hero.W)
			i.RootNode.hero.Collide.X = int32(i.RootNode.W) - int32(i.RootNode.hero.W)
		}
	}

	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		if v := inpututil.KeyPressDuration(ebiten.KeySpace); 0 < v {
			if v == 1 {
				bullet := NewBullet(i.RootNode, i.RootNode.hero,0)
				i.RootNode.hero.bullet = append(i.RootNode.hero.bullet, bullet)

				bullet.Collide.AddTags("heroBullet")
				i.RootNode.hero.CollideSpaceBullet.Add(bullet.Collide)
			}
		}
		i.RootNode.scenesIng = 1
	}

	return nil
}
