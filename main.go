package main

import (
	"game_fly/plays"
	"github.com/hajimehoshi/ebiten"
)



func main() {
	game := plays.NewGame()
	game.Onload()
	ebiten.Run(game.Update, 360, 640, 1, "灰机大战")
}
