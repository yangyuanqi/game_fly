package core

import (
	"fmt"
	"math"
)

func MvCos(v1 float64) (v2 float64) {
	v2 = math.Sin(v1)
	fmt.Println(v1, v2)
	return
}
