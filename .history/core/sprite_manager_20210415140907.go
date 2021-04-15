package core

import (
	"game_fly/plays/conf"

	"github.com/hajimehoshi/ebiten"
)

var Nodes *Node

func init() {
	Nodes = &Node{}
	Nodes.Img = make(map[string]*ebiten.Image)
}

type GameObject interface {
	OnLoad()
	Update(screen *ebiten.Image) (err error)
	AttrDel() bool
}
type Node struct {
	objectGames []GameObject
	//预制体
	prefab     []GameObject
	WinW, WinH int32
	Img        map[string]*ebiten.Image
}

func (n *Node) GameObjectLen() int {
	return len(n.objectGames) + len(n.prefab)
}

func (n *Node) AddGameObject(gameObject GameObject) {
	n.objectGames = append(n.objectGames, gameObject)
}

func (n *Node) AddPrefab(gameObject GameObject) {
	n.prefab = append(n.prefab, gameObject)
}

func (n *Node) OnLoad() {
	n.WinW, n.WinH = int32(conf.GetConfInt("scenes_width")), int32(conf.GetConfInt("scenes_height"))
	for _, v := range n.objectGames {
		v.OnLoad()
	}
}

func (c *Node) Update(screen *ebiten.Image) (err error) {

	var newPrefab []GameObject
	for _, v := range c.prefab {
		//删除对象
		if v.AttrDel() == false {
			newPrefab = append(newPrefab, v)
		}
	}
	c.prefab = newPrefab

	var objectGamesUpdate []GameObject
	objectGamesUpdate = append(c.objectGames, c.prefab...)
	for _, v := range objectGamesUpdate {
		v.Update(screen)

		// l := len(core.Game.GameObjects)
		// for i := k + 1; i < l; i++ {
		// 	if v != core.Game.GameObjects[i] {
		// 		//碰撞边界,
		// 		//碰撞
		// 		if v.GetCollider().Shape.IsColliding(core.Game.GameObjects[i].GetCollider().Shape) {
		// 			if v.GetCollider().IsTrigger || core.Game.GameObjects[i].GetCollider().IsTrigger {
		// 				w, h := v.GetCollider().GetSize()
		// 				w2, h2 := core.Game.GameObjects[i].GetCollider().GetSize()
		// 				if w-w2 < core.Game.GameObjects[i].GetCollider().Transform.Position.X || h-h2 < core.Game.GameObjects[i].GetCollider().Transform.Position.Y || core.Game.GameObjects[i].GetCollider().Transform.Position.X < 0 || core.Game.GameObjects[i].GetCollider().Transform.Position.Y < 0 {
		// 					v.OnTriggerEnter(core.Game.GameObjects[i])
		// 					core.Game.GameObjects[i].OnTriggerEnter(v)
		// 					//fmt.Println("碰撞1",time.Now())
		// 				}
		// 			} else {
		// 				v.OnCollisionEnter(core.Game.GameObjects[i])
		// 				core.Game.GameObjects[i].OnCollisionEnter(v)
		// 				//fmt.Println("碰撞2",time.Now())
		// 			}
		// 		}
		// 	}
		// }
	}
	return nil
}

// func inArray(name string, names []string) bool {
// 	for _, v := range names {
// 		if name == v {
// 			return true
// 		}
// 	}
// 	return false
// }
