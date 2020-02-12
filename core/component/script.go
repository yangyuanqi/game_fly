package component

import "github.com/hajimehoshi/ebiten"

type Script interface {
	OnLoad()
	Start()
	OnCollisionEnter()//[碰撞的对象]//碰撞回调
	Update(screen *ebiten.Image) (err error)
}
