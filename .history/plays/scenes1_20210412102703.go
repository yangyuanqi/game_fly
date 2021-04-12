package plays

import (
	"game_fly/core"
	"game_fly/plays/assets/images"
)

func LodImg() {
	core.Nodes.Img["bullet"] = core.Byte2Image(images.Myb_1)
}
