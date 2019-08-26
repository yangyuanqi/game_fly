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
	ebiten.Run(game.Update, 320, 480, 1, "灰机大战")
}

type Core struct {
}

func (c *Core) OnLoad() {
	core.RegisterScene(plays.NewGame(), "game")
	core.SetScenesIng("game")
}

func (c *Core) Update(screen *ebiten.Image) (err error) {
	core.GetScene(core.ScenesIng).Update(screen)
	return
}
