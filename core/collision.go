package core

import "math"

type P interface {
	GetPosition() (x, y float64, w, h int)
}

func CheckCollision(p1, p2 P) (b bool) {
	x1, y1, w1, h1 := p1.GetPosition()
	x2, y2, w2, h2 := p2.GetPosition()

	centerX1 := math.Round(x1 + float64(w1)/2)
	centerY1 := math.Round(y1 + float64(h1)/2)

	centerX2 := math.Round(x2 + float64(w2)/2)
	centerY2 := math.Round(y2 + float64(h2)/2)

	if math.Abs(centerX1-centerX2) < float64(w1+w2)/2 && math.Abs(centerY1-centerY2) < float64(h1+h2)/2 {
		return true
	}

	return false
}
