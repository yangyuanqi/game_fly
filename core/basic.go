package core

import (
	"math"
)

type Sprite struct {
	X, Y      float64
	W, H      int
	ScaleW    float64
	ScaleH    float64
	MX, MY    float64 //运动规律，当实例化后怎么运动会依靠此参数
	Visible   bool    //隐藏
	Component []Component
}

func (s *Sprite) OnLoad() {

}

func (s *Sprite) Move() {
	if s.MX > 0 {
		s.X += s.MX
		//s.Collide.X = int32(s.X)
	} else if s.MX < 0 {
		s.X -= math.Abs(s.MX)
		//s.Collide.X = int32(s.X)
	}

	if s.MY > 0 {
		s.Y += s.MY
		//s.Collide.Y = int32(s.Y)
	} else if s.MY < 0 {
		s.Y -= math.Abs(s.MY)
		//s.Collide.Y = int32(s.Y)
	}
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

func (s *Sprite) AddSystem(sys Component) {
	s.Component = append(s.Component, sys)
	return
}
