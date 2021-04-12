package plays

import (
	"game_fly/core"
	"game_fly/plays/assets/images"
)

func LodImg() {
	core.Nodes.Img["bullet"] = core.Byte2Image(images.Myb_1)
	core.Nodes.Img["t1png"] = core.Byte2Image(images.T1png)
	core.Nodes.Img["t2png"] = core.Byte2Image(images.T2png)
	core.Nodes.Img["itempng"] = core.Byte2Image(images.Itempng)
}
