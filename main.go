package main

import (
	"game_fly/assets/scenes"
	"game_fly/core"
	"github.com/hajimehoshi/ebiten/v2"
	"log"
)

var Game = &core.Game{}

func main() {
	f := &scenes.FlyScene{}
	Game.LoadScene(f)

	ebiten.SetWindowSize(450, 800)
	ebiten.SetWindowTitle("Animation (Ebiten Demo)")
	ebiten.SetWindowResizable(true)
	if err := ebiten.RunGame(Game); err != nil {
		log.Fatal(err)
	}
}
