package main

import (
	"game_fly/plays"
	"github.com/hajimehoshi/ebiten"
)

// todo camera

func main() {
	game := plays.NewGame()
	game.Onload()
	ebiten.Run(game.Update, game.W, game.H, 1, "灰机大战")
}
