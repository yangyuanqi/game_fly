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
}

func Update(screen *ebiten.Image) (err error) {
	core.Nodes.Update(screen)

	msg := fmt.Sprintf(`fps:%.2f  GameObject:%d`, ebiten.CurrentFPS(), core.Nodes.GameObjectLen())
	ebitenutil.DebugPrint(screen, msg)
	return
}
