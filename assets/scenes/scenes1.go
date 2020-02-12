package scenes

import (
	"game_fly/assets/scripts"
	"game_fly/core/component"
	"github.com/325Gerbils/go-vector"
	"github.com/hajimehoshi/ebiten"
)

var Game *component.Scene

func NewScene() {
	Game = &component.Scene{}
	//在场景中添加游戏对象
	enemy := scripts.EnemyGameObject()
	Game.AddGameObject(enemy)

	role := scripts.RoleGameObject(vector.New(100,100))
	Game.AddGameObject(role)
}

func Update(screen *ebiten.Image) (err error) {
	for k, v := range Game.GameObjects {
		if v.GetGameObject().Del == true {
			Game.GameObjects = append(Game.GameObjects[0:k], Game.GameObjects[k+1:]...)
			continue
		}

		if v.GetCollider().Shape != nil {
			v.GetCollider().DrawCollider(screen)
		}
		v.Update(screen)

		for _, v2 := range Game.GameObjects {
			if v != v2 {
				if v.GetCollider().Shape.IsColliding(v2.GetCollider().Shape) {
					v.OnCollisionEnter(v2)
					v2.OnCollisionEnter(v)
				}
			}

		}
	}

	//检测碰撞

	return
}
