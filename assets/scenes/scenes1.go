package scenes

import (
	"game_fly/assets/scripts"
	"game_fly/core"
	"game_fly/core/component"
	"github.com/325Gerbils/go-vector"
	"github.com/hajimehoshi/ebiten"
)

//var Game *component.Scene

func NewScene() {
	core.Game = &component.Scene{}
	//在场景中添加游戏对象
	background := scripts.BackgroundObject()
	core.Game.AddGameObject(background)

	enemy := scripts.EnemyGameObject()
	core.Game.AddGameObject(enemy)

	role := scripts.RoleGameObject(vector.New(100, 100))
	core.Game.AddGameObject(role)
}

func Update(screen *ebiten.Image) (err error) {
	var prefabNames []string
	for k, v := range core.Game.GameObjects {
		//删除对象
		if v.GetGameObject().Del == true {
			core.Game.GameObjects = append(core.Game.GameObjects[:k], core.Game.GameObjects[k+1:]...)
		}

		l := len(core.Game.GameObjects)
		for i := k + 1; i < l; i++ {
			if v != core.Game.GameObjects[i] {
				//碰撞边界,
				//碰撞
				if v.GetCollider().Shape.IsColliding(core.Game.GameObjects[i].GetCollider().Shape) {
					if v.GetCollider().IsTrigger || core.Game.GameObjects[i].GetCollider().IsTrigger {
						w, h := v.GetCollider().GetSize()
						w2, h2 := core.Game.GameObjects[i].GetCollider().GetSize()
						if w-w2 < core.Game.GameObjects[i].GetCollider().Transform.Position.X || h-h2 < core.Game.GameObjects[i].GetCollider().Transform.Position.Y || core.Game.GameObjects[i].GetCollider().Transform.Position.X < 0 || core.Game.GameObjects[i].GetCollider().Transform.Position.Y < 0 {
							v.OnTriggerEnter(core.Game.GameObjects[i])
							core.Game.GameObjects[i].OnTriggerEnter(v)
							//fmt.Println("碰撞1",time.Now())
						}
					} else {
						v.OnCollisionEnter(core.Game.GameObjects[i])
						core.Game.GameObjects[i].OnCollisionEnter(v)
						//fmt.Println("碰撞2",time.Now())
					}
				}
			}
		}

		//渲染碰撞
		if v.GetCollider().Shape != nil {
			v.GetCollider().DrawCollider(screen)
		}
		if v.GetGameObject().IsPrefab == false {
			v.Update(screen)
		} else {
			if inArray(v.GetGameObject().Name, prefabNames) == false {
				prefabNames = append(prefabNames, v.GetGameObject().Name)
				v.Update(screen)
			}
		}
	}
	return
}

func inArray(name string, names []string) bool {
	for _, v := range names {
		if name == v {
			return true
		}
	}
	return false
}
