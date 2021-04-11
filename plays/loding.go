package plays

import (
	"game_fly/core"
	"game_fly/plays/assets/images"

	"github.com/hajimehoshi/ebiten"
)

var Scenes map[string]*ebiten.Image

func Loding(){
	Scenes["bullet"] = core.Byte2Image(images.Myb_1)
}
