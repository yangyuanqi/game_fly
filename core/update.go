package core

import (
	"game_fly/core/sprite"
	"github.com/hajimehoshi/ebiten"
)

func GetComponentUpdate(screen *ebiten.Image) {
	//渲染精灵
	for _, v := range sprite.Sprites {
		for _, v2 := range v.Sprite {
			v2.UpdateResolv()
			//v2.SetScreen(screen)
			v2.Update(screen)
			for _, v3 := range v2.GetComponent() {
				v3.Update(screen)
			}
		}
		//渲染预制体
		for _, v2 := range v.Prefab {
			v2.UpdateResolv()
			//v2.SetScreen(screen)
			v2.Update(screen)
			for _, v3 := range v2.GetComponent() {
				v3.Update(screen)
			}
		}
		//Ui
		for _, v2 := range v.Ui {
			v2.Update(screen)
		}
	}
}
