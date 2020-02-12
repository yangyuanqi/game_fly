package core

import (
	"game_fly/core/component"
)

//input,move,update,start,collider...
var Renderer *Render

func init(){
	Renderer = &Render{}
}

type Render struct {
	Render *component.Scene
}

//func (rd *Render) AddScene(scene *component.Scene) {
//	Renderer = append(Renderer, scene)
//}

//func (rd *Render) Update(screen *ebiten.Image) (err error) {
//	for _, v := range rd.Render.GameObjects {
//			//v2.(component.GameObject).GetComponent("Script").(component.Script).Update(screen)
//			//fmt.Println(v.(component.GameObject))
//	}
//
//	//检测碰撞
//
//	return
//}

//
//func (rd *Render) OnCollisionEnter() {
//
//}
