package core

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
