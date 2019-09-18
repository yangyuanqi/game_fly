package plays

import (
	"game_fly/core"
	"game_fly/core/prefab"
	"game_fly/core/sprite"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
)

type Input struct {
	core.Component
}

func NewInput() (input *Input) {
	return &Input{}
}

func (i *Input) OnLoad() {

}

func (i *Input) Start() {

}

func (i *Input) Update(screen *ebiten.Image) (err error) {
	hero := sprite.GetSprite("game", "role").(*Role)
	game := core.GetScene("game").(*Game)
	hero.Left = 0
	hero.Right = 0

	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		if v := inpututil.KeyPressDuration(ebiten.KeyUp); 0 < v {
			if v == 1 || v%1 == 0 {
				hero.Y -= 5
			}
		}
		if hero.Y <= 0 {
			hero.Y = 0
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		if v := inpututil.KeyPressDuration(ebiten.KeyDown); 0 < v {
			if v == 1 || v%1 == 0 {
				hero.Y += 5
			}
		}
		if hero.Y >= float64(game.H)-float64(hero.H) {
			hero.Y = float64(game.H) - float64(hero.H)
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		if v := inpututil.KeyPressDuration(ebiten.KeyLeft); 0 < v {
			if v == 1 || v%1 == 0 {
				hero.X -= 5
			}
			hero.Left = v
		}
		if hero.X <= 0 {
			hero.X = 0
		}
	}

	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		if v := inpututil.KeyPressDuration(ebiten.KeyRight); 0 < v {
			if v == 1 || v%1 == 0 {
				hero.X += 5
			}
			hero.Right = v
		}
		if hero.X >= float64(game.W)-float64(hero.W) {
			hero.X = float64(game.W) - float64(hero.W)
		}
	}

	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		if v := inpututil.KeyPressDuration(ebiten.KeySpace); 0 < v {
			if v == 1 {
				bullet := NewBullet(BulletImg)
				bullet.Move = func(bullet2 *Bullet) {
					bullet2.Y -= 10
				}
				prefab.AddPrefab(bullet, "roleBullet")
			}
		}
	}

	return nil
}
