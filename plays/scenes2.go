package plays

import (
	"fmt"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/inpututil"
)

type Scenes2 struct {
}

func (s *Scenes2) OnLoad() {

}

func (s *Scenes2) Update(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen,"c2")
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		if v := inpututil.KeyPressDuration(ebiten.KeyUp); 0 < v {
			fmt.Println("keyUp")
		}
	}
}
