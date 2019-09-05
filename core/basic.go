package core

import "game_fly/core/component"

type Sprite struct {
	Id              string
	X, Y            float64
	W, H            int
	ScaleW          float64
	ScaleH          float64
	MX, MY          float64 //运动规律，当实例化后怎么运动会依靠此参数
	Visible         bool    //隐藏
	SpriteComponent []component.SpriteComponent
	Prefab          []component.SpriteComponent
}

func (s *Sprite) OnLoad() {

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

func (s *Sprite) GetPosition() (x, y float64, w, h int) {
	return s.X, s.Y, s.W, s.H
}

func (s *Sprite) GetPrefab() (prefab []component.SpriteComponent) {
	return s.Prefab
}
