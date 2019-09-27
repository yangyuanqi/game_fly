package core

import (
	"fmt"
	"strconv"
)

func GetXYOfDegree(x, y float64, degree float64, speed float64) (mx, my float64) {
	if (degree < 90 && degree > 0) || (degree > 270 && degree < 360) {
		if degree < 90 {
			x += speed
		} else {
			x -= speed
		}
		y -= 1
		return x, y
	} else {
		if degree < 180 {
			x += speed
		} else {
			x -= speed
		}
		y += 1
		return x, y
	}
	return
}

func ScaliEq(a1, a2 float64) (s float64) {
	if a1 > a2 {
		s, _ = strconv.ParseFloat(fmt.Sprintf("%.4f", a2/a1), 64)
	} else {
		s, _ = strconv.ParseFloat(fmt.Sprintf("%.4f", a1/a2), 64)
	}
	return
}
