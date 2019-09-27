package core

import "math"

func MoveTo(tx, ty float64, t string) (x, y float64) {
	switch t {
	case "sin":
		ty = math.Sin(tx)
	}
	return tx, ty
}
