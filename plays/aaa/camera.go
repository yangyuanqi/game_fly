package aaa

var CameraGlobal *Camera

type Camera struct {
	MaxW float64
	MaxH float64
	TW   float64
	LW   float64
}

func CreateCamera(maxW, maxH, tW, lW float64) {
	CameraGlobal = &Camera{
		MaxW: maxW,
		MaxH: maxH,
		TW:   tW,
		LW:   lW,
	}
}

