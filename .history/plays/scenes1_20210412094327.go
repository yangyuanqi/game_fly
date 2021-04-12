package plays

import (
	"game_fly/core"
	"game_fly/plays/assets/images"

	"github.com/hajimehoshi/ebiten"
)

var Scenes1 map[string]*ebiten.Image

func Loding() {
	Scenes1["bullet"] = core.Byte2Image(images.Myb_1)
}
