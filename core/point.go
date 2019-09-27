package core

type Rect struct {
	X, Y float64
	W, H float64
}

func (r *Rect) Move(dx, dy float64) {
	r.X += dx
	r.Y += dy
}

func (r *Rect) Point() (x, y float64) {
	return r.X, r.Y
}

func (r *Rect) GetPosition() (x, y, w, h float64) {
	return r.X, r.Y, r.W, r.H
}

func (r *Rect) GetPositionInt32() (x, y, w, h int32) {
	return int32(r.X), int32(r.Y), int32(r.W), int32(r.H)
}

func (s *Rect) SetXY(x, y float64) {
	s.X = x
	s.Y = y
}

func (s *Rect) SetWH(w, h float64) {
	s.W = w
	s.H = h
}

func (s *Rect) SetXYWH(x, y, w, h float64) {
	s.X = x
	s.Y = y
	s.W = w
	s.H = h

}
