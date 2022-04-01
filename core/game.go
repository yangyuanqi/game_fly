package core

import (
	"github.com/hajimehoshi/ebiten/v2"
	"time"
)

type Game struct {
	CurrentScene GameObjectInterface
	dt           float64
}

func (g *Game) LoadScene(scene GameObjectInterface) {
	//scene.Init()
	g.CurrentScene = scene
	g.CurrentScene.Init()
	//g.dt = float64(time.Now().UnixMilli())
}
func (g *Game) childrenInitTree(children *[]GameObjectInterface) {
	for _, v := range *children {
		v.Init()
		if len(*v.GetChildren()) > 0 {
			g.childrenInitTree(v.GetChildren())
		}
	}
}

func (g *Game) Update() error {
	//间隔计时
	dt := (float64(time.Now().UnixMilli()) - g.dt) / 1000

	// 根节点
	g.CurrentScene.Update(dt)
	// 渲染根节点的组件
	if len(g.CurrentScene.GetComponent()) > 0 {
		for _, v2 := range g.CurrentScene.GetComponent() {
			v2.Update(dt)
		}
	}
	//渲染子节点与子节点的组件
	g.childrenUpdateTree(g.CurrentScene.GetChildren(), dt)

	//间隔计时
	g.dt = float64(time.Now().UnixMilli())
	return nil
}
func (g *Game) childrenUpdateTree(children *[]GameObjectInterface, dt float64) {
	//销毁
	var newChildren []GameObjectInterface
	for _, v := range *children {
		if v.GetDestroy() != true {
			//*children = append((*children)[:k], (*children)[k+1:]...)
			newChildren = append(newChildren, v)
			//children = append(children[:k], children[k+1:]...)
		}
	}
	*children = newChildren
	//children = &newChildren
	//fmt.Println(children)
	//
	for _, v := range *children {
		v.Update(dt)
		if len(v.GetComponent()) > 0 {
			for _, v2 := range v.GetComponent() {
				v2.Update(dt)
			}
		}
		if len(*v.GetChildren()) > 0 {
			g.childrenUpdateTree(v.GetChildren(), dt)
		}
	}
}
func (g *Game) Draw(screen *ebiten.Image) {
	g.CurrentScene.Draw(screen)
	if len(g.CurrentScene.GetComponent()) > 0 {
		for _, v2 := range g.CurrentScene.GetComponent() {
			v2.Draw(screen)
		}
	}
	g.childrenDrawTree(screen, g.CurrentScene.GetChildren())
}
func (g *Game) childrenDrawTree(screen *ebiten.Image, children *[]GameObjectInterface) {
	for _, v := range *children {
		//渲染节点
		v.Draw(screen)
		//渲染组件
		if len(v.GetComponent()) > 0 {
			for _, v2 := range v.GetComponent() {
				v2.Draw(screen)
			}
		}
		if len(*v.GetChildren()) > 0 {
			g.childrenDrawTree(screen, v.GetChildren())
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}
