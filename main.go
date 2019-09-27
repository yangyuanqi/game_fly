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
	Run(320, 480, 1, "灰机大战")
	//ebiten.Run(game.Update, 300, 300, 1, "灰机大战")
}

func Run(width, height int, scale float64, title string) {
	conf.CC = conf.Conf{}
	conf.CC.SetScenes(width, height, scale, title)
	//场景切换示例
	//core.SetTimer(time.Second*5, func() {
	//注册场景
	//core.RegisterScene(plays.NewLogin(), "login")
	//core.RegisterScene(plays.NewScenesHello(), "hello")
	//core.SetScenesIng("login")
	//})
	core.RegisterScene(plays.NewLogin(), "login")
	core.SetScenesIng("login")
	core.GetComponentStart()
	ebiten.Run(Update, width, height, scale, title)
}

func Update(screen *ebiten.Image) (err error) {
	core.RootNode = screen
	//ui,与背景，要先渲染，没有顺序要求的功能可以使用goroutine
	//渲染场景
	core.GetScene(core.ScenesIng).(core.Scene).Update(screen)
	//渲染精灵，组件
	core.GetComponentUpdate(screen)

	msg := fmt.Sprintf(`fps:%.2f  componentNum:%d`, ebiten.CurrentFPS(), prefab.PrefabLen())
	ebitenutil.DebugPrint(screen, msg)
	return
}
