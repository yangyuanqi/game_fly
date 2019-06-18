package core

import (
	"github.com/Bredgren/geo"
	"math"
)

type BaseComponent struct {
	geo.Rect
	ScaleW    float64
	ScaleH    float64
	MX, MY    float64 //运动规律，当实例化后怎么运动会依靠此参数
	Visible   bool    //隐藏
}

func (s *BaseComponent) Move() {
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

func (s *BaseComponent) SetXY(x, y float64) {
	s.X = x
	s.Y = y
}

func (s *BaseComponent) SetWH(w, h float64) {
	s.W = w
	s.H = h
}

func (s *BaseComponent) SetScale(sw, sh float64) {
	s.ScaleW = sw
	s.ScaleH = sh
}

func (s *BaseComponent) GetPosition() (x, y float64, w, h float64) {
	return s.X, s.Y, s.W, s.H
}

func (s *BaseComponent) AddSystem(sys Component) {
	//s.Component = append(s.Component, sys)
	return
}
