//数据结构，场景，预制体，精灵，基础结构，组件,ui
//组件存在基础机构里面
/*
	-------onload
	-------start
	-------update

	->加载资源，初始化
    ->添加组件
*/

package main

import (
	"fmt"
	"game_fly/core"
	"game_fly/plays"
	"game_fly/plays/conf"

	"github.com/SolarLune/resolv/resolv"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

var shape1 *resolv.Rectangle
var shape2 *resolv.Rectangle

func main() {
	plays.LodImg()
	core.Nodes.AddGameObject(plays.NewMap())
	core.Nodes.AddGameObject(plays.NewRole())
	core.Nodes.OnLoad()
	ebiten.Run(Update, conf.GetConfInt("scenes_width"), conf.GetConfInt("scenes_height"), 1, "灰机大战")

	shape1 = resolv.NewRectangle(10, 10, 16, 16)

	// Create another rectangle, as well.
	shape2 = resolv.NewRectangle(11, 100, 16, 16)
}

func Update(screen *ebiten.Image) (err error) {
	// Later on, in the game's update loop...

	// Let's say we were trying to move the shape to the right by 2 pixels. We'll create a delta X movement variable
	// that stores a value of 2.
	dx := 2

	// Here, we check to see if there's a collision should shape1 try to move to the right by the delta movement. The Resolve()
	// functions return a Collision object that has information about whether the attempted movement resulted in a collision
	// or not.
	resolution := resolv.Resolve(shape1, shape2, int32(dx), 0)

	if resolution.Colliding() {

		// If there was a collision, then shape1 couldn't move fully to the right. It came into contact with shape2,
		// and the variable "resolution" now holds a Collision struct with helpful information, like how far to move to be touching.

		// Here we just move the shape over to the right by the distance reported by the Collision struct so it'll come into contact
		// with shape2.
		shape1.X += resolution.ResolveX

	} else {

		// If there wasn't a collision, shape1 should be able to move fully to the right, so we move it.
		shape1.X += int32(dx)

	}

	// We can also do collision testing only pretty simply:

	colliding := shape1.IsColliding(shape2)

	if colliding {
		fmt.Println("WHOA! shape1 and shape2 are colliding.")
	}

	//core.Nodes.Update(screen)

	msg := fmt.Sprintf(`fps:%.2f  GameObject:%d`, ebiten.CurrentFPS(), core.Nodes.GameObjectLen())
	ebitenutil.DebugPrint(screen, msg)
	return
}
