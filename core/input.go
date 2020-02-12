package core

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
)

func GetAxisRaw(raw string) (r int) {
	if raw == "Y" {
		if ebiten.IsKeyPressed(ebiten.KeyUp) {
			if v := inpututil.KeyPressDuration(ebiten.KeyUp); v > 0 {
				return -1
			}
		}

		if ebiten.IsKeyPressed(ebiten.KeyDown) {
			if v := inpututil.KeyPressDuration(ebiten.KeyDown); 0 < v {
				return 1
			}
		}
	}

	if raw == "X" {
		if ebiten.IsKeyPressed(ebiten.KeyLeft) {
			if v := inpututil.KeyPressDuration(ebiten.KeyLeft); 0 < v {
				return -1
			}
		}

		if ebiten.IsKeyPressed(ebiten.KeyRight) {
			if v := inpututil.KeyPressDuration(ebiten.KeyRight); 0 < v {
				return 1
			}
		}
	}
	return 0
}
