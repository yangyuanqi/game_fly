package core

import (
	"game_fly/core/sprite"
)

func GetComponentStart() {
	//渲染精灵
	for _, v := range sprite.Sprites {
		for _, v2 := range v.Sprite {
			v2.Start()
			for _, v3 := range v2.GetComponent() {
				v3.Start()
			}
		}
		//渲染预制体
		//v.Prefab.Range(func(key, value interface{}) bool {
		//	value.(data.SpriteComponent).Start()
		//	for _, v3 := range value.(data.SpriteComponent).GetComponent() {
		//		v3.Start()
		//	}
		//	return true
		//})

		for _, v2 := range v.Prefab {
			v2.Start()
			for _, v3 := range v2.GetComponent() {
				v3.Start()
			}
		}
		//ui
		for _, v2 := range v.Ui {
			v2.Start()
		}
	}
}