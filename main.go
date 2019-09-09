package main

import (
	"fmt"
	"game_fly/core"
	"game_fly/core/prefab"
	"game_fly/plays"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

// todo camera

func main() {
	game := Core{}
	game.OnLoad()
	game.Start()

	ebiten.Run(game.Update, 320, 480, 1, "灰机大战")
}

type Core struct {
}

func (c *Core) OnLoad() {
	//场景切换示例
	//core.SetTimer(time.Second*5, func() {
	core.RegisterScene(plays.NewGame(), "game")
	core.SetScenesIng("game")
	//})
	core.RegisterScene(plays.NewLogin(), "login")
	//core.SetScenesIng("login")
}

func (c *Core) Start() {
	core.GetComponentStart()
}

func (c *Core) Update(screen *ebiten.Image) (err error) {
	core.GetScene(core.ScenesIng).(core.Scene).Update(screen)
	core.GetComponentUpdate(screen)
	msg := fmt.Sprintf(`fps:%.2f  componentNum:%d`, ebiten.CurrentFPS(), prefab.PrefabLen())
	ebitenutil.DebugPrint(screen, msg)
	return
}
