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

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

// todo camera
//var Load2 bool

func main() {
	Loding()
	//go func() {
	//plays.Scenes = &plays.ScenesHello{}
	//core.RegisterScene(plays.Scenes, "hello")
	//core.LoadScene("hello")
	//time.Sleep(time.Second * 1)
	//Load2 = true
	//}()
	//scenes.NewScene()
	core.Nodes.AddGameObject(plays.NewMap())
	core.Nodes.AddGameObject(plays.NewRole())
	core.Nodes.OnLoad()
	ebiten.Run(Update, conf.GetConfInt("scenes_width"), conf.GetConfInt("scenes_height"), 1, "灰机大战")
}

func Update(screen *ebiten.Image) (err error) {

	//scenes1 := scenes.NewScene()
	//fmt.Println(scenes1)
	core.Nodes.Update(screen)

	msg := fmt.Sprintf(`fps:%.2f  GameObject:%d`, ebiten.CurrentFPS(), core.Nodes.GameObjectLen())
	ebitenutil.DebugPrint(screen, msg)
	//for _, v := range scenes1.GameObjects {
	//v.Script
	//}

	//if Load2 {
	//	core.GetScene("hello").Update(screen)
	//	msg := fmt.Sprintf(`fps:%.2f  componentNum:%d`, ebiten.CurrentFPS())
	//	ebitenutil.DebugPrint(screen, msg)
	//} else {
	//	msg := fmt.Sprintf(`fps:%.2f  componentNum:%d %s`, ebiten.CurrentFPS(), "Loading...")
	//	ebitenutil.DebugPrint(screen, msg)
	//}

	return
}
