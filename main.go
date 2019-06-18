package main

import (
	"game_fly/core"
	"game_fly/plays"
	"github.com/hajimehoshi/ebiten"
)

// todo camera

func main() {
	game := Core{}
	game.OnLoad()
	ebiten.Run(game.Update, 640, 640, 1, "灰机大战")
}

type Core struct {
}

func (c *Core) OnLoad() {
	core.RegisterScene(&plays.Scenes1{}, "scen1")
	core.RegisterScene(&plays.Scenes2{}, "scen2")
	core.SetScenesIng("scen1")
}

func (c *Core) Update(screen *ebiten.Image) (err error) {
	core.GetScene(core.ScenesIng).Update(screen)
	return
}