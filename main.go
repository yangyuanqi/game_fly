package main

import (
	"game_fly/assets/scenes"
	"game_fly/core"
	"github.com/hajimehoshi/ebiten/v2"
)

var Game = &core.Game{}

func main() {
	ebiten.SetWindowSize(450, 800)
	ebiten.SetWindowTitle("Animation (Ebiten Demo)")
	f := &scenes.FlyScene{}
	Game.LoadScene(f)
	ebiten.RunGame(Game)
}
