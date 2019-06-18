package plays

import (
	"game_fly/core"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
)

type Input struct {
	core.Sprite
	RootNode *Game
}

func NewInput(rootNode *Game) (input *Input) {
	input = &Input{}
	input.RootNode = rootNode
	return
}

func (i *Input) Update() (err error) {
	i.RootNode.hero.Left = 0
	i.RootNode.hero.Right = 0

	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		if v := inpututil.KeyPressDuration(ebiten.KeyUp); 0 < v {
			if v == 1 || v%1 == 0 {
				i.RootNode.hero.Y -= 5
			}
		}
		if i.RootNode.hero.Y <= 0 {
			i.RootNode.hero.Y = 0
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		if v := inpututil.KeyPressDuration(ebiten.KeyDown); 0 < v {
			if v == 1 || v%1 == 0 {
				i.RootNode.hero.Y += 5
			}
		}
		if i.RootNode.hero.Y >= float64(i.RootNode.H)-float64(i.RootNode.hero.H) {
			i.RootNode.hero.Y = float64(i.RootNode.H) - float64(i.RootNode.hero.H)
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		if v := inpututil.KeyPressDuration(ebiten.KeyLeft); 0 < v {
			if v == 1 || v%1 == 0 {
				i.RootNode.hero.X -= 5
			}
			i.RootNode.hero.Left = v
		}
		if i.RootNode.hero.X <= 0 {
			i.RootNode.hero.X = 0
		}
	}
	// 當「按鍵右」被按下時⋯⋯
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		if v := inpututil.KeyPressDuration(ebiten.KeyRight); 0 < v {
			if v == 1 || v%1 == 0 {
				i.RootNode.hero.X += 5
			}
			i.RootNode.hero.Right = v
		}
		if i.RootNode.hero.X >= float64(i.RootNode.W)-float64(i.RootNode.hero.W) {
			i.RootNode.hero.X = float64(i.RootNode.W) - float64(i.RootNode.hero.W)
		}
	}

	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		if v := inpututil.KeyPressDuration(ebiten.KeySpace); 0 < v {
			if v == 1 {
				bullet := NewBullet(i.RootNode, i.RootNode.hero, BulletImg)
				bullet.MY = -10
				i.RootNode.hero.GroupBullet = append(i.RootNode.hero.GroupBullet, bullet)
			}
		}
		i.RootNode.scenesIng = 1
	}

	return nil
}
