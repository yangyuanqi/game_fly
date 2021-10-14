package plays

import (
	"game_fly/core"
	"game_fly/plays/assets/images"
)

var FlyGameSceneVar *FlyGameScene

type FlyGameScene struct {
	core.Scene
}

func init() {
	FlyGameSceneVar = &FlyGameScene{}

	FlyGameSceneVar.RootNode.Img["bullet"] = core.Byte2Image(images.Myb_1)
	FlyGameSceneVar.RootNode.Img["t1png"] = core.Byte2Image(images.T1png)
	FlyGameSceneVar.RootNode.Img["t2png"] = core.Byte2Image(images.T2png)
	FlyGameSceneVar.RootNode.Img["itempng"] = core.Byte2Image(images.Itempng)

	FlyGameSceneVar.RootNode.AddGameObject(NewMap())
	FlyGameSceneVar.RootNode.AddGameObject(NewRole())
	FlyGameSceneVar.RootNode.AddGameObject(NewTest())
	FlyGameSceneVar.RootNode.OnLoad()
}
