package core

import (
	"github.com/hajimehoshi/ebiten"
)

func forChild(screen *ebiten.Image, child []SpriteComponent) {
	if len(child) > 0 {
		for _, v := range child {
			v.Update(screen)
			v.UpdateWrap(screen)
			for _, v := range v.GetPrefab() {
				v.Update(screen)
				v.UpdateWrap(screen)
			}
			forChild(screen, v.GetChild())
		}
	}
}

func GetComponentUpdate(screen *ebiten.Image, sprite SpriteComponent) {
	//渲染精灵
	//sprite.Update(screen)
	for _, v := range sprite.GetPrefab() {
		v.Update(screen)
		v.UpdateWrap(screen)
	}
	forChild(screen, sprite.GetChild())

	//for _, v := range sprite.Sprites {
	//	for _, v2 := range v.Sprite {
	//		v2.UpdateResolv()
	//		v2.Update(screen)
	//		for _, v3 := range v2.GetComponent() {
	//			v3.Update(screen)
	//		}
	//	}
	//for _, v2 := range v.Prefab {
	//	v2.UpdateResolv()
	//	//v2.SetScreen(screen)
	//	v2.Update(screen)
	//	if v2.GetVisible() {
	//		delete(v.Prefab, v2.GetId())
	//	}
	//	for _, v3 := range v2.GetComponent() {
	//		v3.Update(screen)
	//	}
	//}
	//Ui
	//for _, v2 := range v.Ui {
	//	v2.Update(screen)
	//}

	//渲染预制体
	//v.Prefab.Range(func(key, value interface{}) bool {
	//
	//	value.(data.SpriteComponent).UpdateResolv()
	//	value.(data.SpriteComponent).Update(screen)
	//	//if value.(component.SpriteComponent).GetVisible() {
	//	//	v.Prefab.Delete(value.(component.SpriteComponent).GetId())
	//	//}
	//	for _, v3 := range value.(data.SpriteComponent).GetComponent() {
	//		v3.Update(screen)
	//	}
	//	return true
	//})

	//}

	//data.PrefabL.Lock()
	//for _, v := range sprite.Sprites {
	//	for _, value := range v.Prefab {
	//		value.UpdateResolv()
	//		value.Update(screen)
	//		if value.GetDestroy() {
	//			prefab.DelPrefab(value.GetId())
	//		}
	//		for _, v3 := range value.GetComponent() {
	//			v3.Update(screen)
	//		}
	//	}
	//}
	//data.PrefabL.Unlock()

}
