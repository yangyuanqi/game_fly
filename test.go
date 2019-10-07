package main

import (
	"fmt"
	"math"
)

func main() {
	for i := 0; i < 480; i++ {
		fmt.Println(math.Sin(float64(i)/5))
	}
}
