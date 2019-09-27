package core

import (
	"game_fly/core/data"
	"github.com/SolarLune/resolv/resolv"
	"github.com/hajimehoshi/ebiten"
	"strconv"
)

type Sprite struct {
	Id   string //唯一标识
	Name string //变量名称
	Rect
	//geo.Rect
	ScaleW   float64
	ScaleH   float64
	Visible  bool          //隐藏
	Material *ebiten.Image //材质
	Destroy  bool          //销毁
	//Screen    *ebiten.Image
	Component []data.SpriteComponent //组件
	Collision *resolv.Rectangle
}

func (s *Sprite) OnLoad() {

}

//初始化系统属性
func (s *Sprite) Create(name string) {
	ID++
	id := strconv.Itoa(ID)
	s.Id = id
	s.Name = name
}

func (s *Sprite) SetScale(sw, sh float64) {
	s.ScaleW = sw
	s.ScaleH = sh
}

func (s *Sprite) GetName() (name string) {
	return s.Name
}

func (s *Sprite) GetId() (id string) {
	return s.Id
}

func (s *Sprite) GetResolv() (collision interface{}) {
	collision = s.Collision
	return
}

//同步碰撞
func (s *Sprite) UpdateResolv() {
	if s.Collision != nil {
		s.Collision.SetXY(int32(s.X), int32(s.Y))
	}
}

func (s *Sprite) AddComponent(component data.SpriteComponent) {
	s.Component = append(s.Component, component)
}

//func (s *Sprite) SetScreen(screen *ebiten.Image) {
//	if s.Screen == nil {
//		s.Screen = screen
//	}
//}

func (s *Sprite) GetComponent() (components []data.SpriteComponent) {
	return s.Component
}

func (s *Sprite) GetVisible() (b bool) {
	return s.Visible
}
func (s *Sprite) SetVisible(b bool) () {
	s.Visible = b
}
func (s *Sprite) GetDestroy() (b bool) {
	return s.Destroy
}
func (s *Sprite) SetDestroy(b bool) () {
	s.Visible = b
}

type Component struct {
	Sprite
}
