package core

import (
	"github.com/hajimehoshi/ebiten"
)

var (
	node []GameObject
)

type Core struct {
	objectGame []GameObject
}

func (c *Core) OnLoad() {

}

func (c *Core) Update(screen *ebiten.Image) (err error) {
	var prefabNames []string
	var newGameObjects []GameObjects
	for k, v := range c.Game.GameObjects {
		//删除对象
		if v.GetGameObject().Del != true {
			newGameObjects = append(newGameObjects, v)
		}
	}

	for k, v := range c.Game.newGameObjects {
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
	return nil
}

func inArray(name string, names []string) bool {
	for _, v := range names {
		if name == v {
			return true
		}
	}
	return false
}
