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
	"game_fly/core/prefab"
	"game_fly/plays"
	"game_fly/plays/conf"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

// todo camera

func main() {
	core.RegisterScene(plays.NewLogin(), "login")
	core.SetScenesIng("login")
	ebiten.Run(Update,conf.GetConfInt("scenes_width"), conf.GetConfInt("scenes_height"), 1, "灰机大战")
}

//func Run(width, height int, scale float64, title string) {

	//core.RegisterScene(plays.NewLogin(), "login")
	//core.SetScenesIng("login")
	//ebiten.Run(Update, width, height, scale, title)
//}

func Update(screen *ebiten.Image) (err error) {
	//core.RootNode = screen
	//渲染场景
	core.GetScene(core.ScenesIng).(core.Scene).Update(screen)
	//渲染精灵，组件
	core.GetComponentUpdate(screen)

	msg := fmt.Sprintf(`fps:%.2f  componentNum:%d`, ebiten.CurrentFPS(), prefab.PrefabLen())
	ebitenutil.DebugPrint(screen, msg)
	return
}
