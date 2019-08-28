package main

import (
	"fmt"
	"game_fly/core"
	"game_fly/plays"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

// todo camera

func main() {
	game := Core{}
	game.OnLoad()
	ebiten.Run(game.Update, 320, 480, 1, "灰机大战")
}

type Core struct {

}

func (c *Core) OnLoad() {

	//场景切换示例
	//core.SetTimer(time.Second*5, func() {
	//	core.RegisterScene(plays.NewGame(), "game")
	//	core.SetScenesIng("game")
	//})

	core.RegisterScene(plays.NewLogin(),"login")
	core.SetScenesIng("login")
}

func (c *Core) Update(screen *ebiten.Image) (err error) {
	core.GetScene(core.ScenesIng).Update(screen)

	msg := fmt.Sprintf(`fps:%.2f`, ebiten.CurrentFPS())
	ebitenutil.DebugPrint(screen, msg)
	return
}
