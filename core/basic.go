package core

import (
	"game_fly/core/component"
	"github.com/SolarLune/resolv/resolv"
	"github.com/hajimehoshi/ebiten"
	"strconv"
)

type Sprite struct {
	Id        string
	Name      string
	X, Y      float64
	W, H      int
	ScaleW    float64
	ScaleH    float64
	Visible   bool          //隐藏
	Img       *ebiten.Image //子弹贴图
	Component []component.SpriteComponent
	Collision *resolv.Rectangle
}

func (s *Sprite) OnLoad() {

}

//初始化系统属性
func (s *Sprite) Create() {
	ID++
	id := "bullet" + strconv.Itoa(ID)
	s.Id = id
}

func (s *Sprite) SetXY(x, y float64) {
	s.X = x
	s.Y = y
}

func (s *Sprite) SetWH(w, h int) {
	s.W = w
	s.H = h
}

func (s *Sprite) SetScale(sw, sh float64) {
	s.ScaleW = sw
	s.ScaleH = sh
}

func (s *Sprite) GetName() (name string) {
	return s.Name
}

func (s *Sprite) GetPosition() (x, y, w, h int32) {
	return int32(s.X), int32(s.Y), int32(s.W), int32(s.H)
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

func (s *Sprite) AddComponent(component component.SpriteComponent) {
	s.Component = append(s.Component, component)
}

func (s *Sprite) GetComponent() (components []component.SpriteComponent) {
	return s.Component
}
