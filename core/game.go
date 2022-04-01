package core

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	CurrentScene GameObjectInterface
}

func (g *Game) LoadScene(scene GameObjectInterface) {
	//scene.Init()
	g.CurrentScene = scene
	g.CurrentScene.Init()
}
func (g *Game) childrenInitTree(children []GameObjectInterface) {
	for _, v := range children {
		v.Init()
		if len(v.GetChildren()) > 0 {
			g.childrenInitTree(v.GetChildren())
		}
	}
}

func (g *Game) Update() error {
	g.CurrentScene.Update() //节点
	if len(g.CurrentScene.GetComponent()) > 0 {
		for _, v2 := range g.CurrentScene.GetComponent() {
			v2.Update()
		}
	}

	g.childrenUpdateTree(g.CurrentScene.GetChildren())
	return nil
}
func (g *Game) childrenUpdateTree(children []GameObjectInterface) {
	for _, v := range children {
		v.Update()
		if len(v.GetComponent()) > 0 {
			for _, v2 := range v.GetComponent() {
				v2.Update()
			}
		}
		if len(v.GetChildren()) > 0 {
			g.childrenUpdateTree(v.GetChildren())
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
func (g *Game) childrenDrawTree(screen *ebiten.Image, children []GameObjectInterface) {
	for _, v := range children {
		//渲染节点
		v.Draw(screen)
		//渲染组件
		if len(v.GetComponent()) > 0 {
			for _, v2 := range v.GetComponent() {
				v2.Draw(screen)
			}
		}
		if len(v.GetChildren()) > 0 {
			g.childrenDrawTree(screen, v.GetChildren())
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}
