//通道数据一直避免阻塞
package main

import (
	"fmt"
	"time"
)

var m = make(chan []int)
var m2 = make(chan []int)
var data []int

func main() {
	go func(m chan []int) {
		var mm []int
		for i := 0; i < 999; i++ {
			fmt.Println(1)
			mm := append(mm, i)
			m <- mm
			time.Sleep(time.Second * 1)
		}
	}(m)

	//go func(m chan []int) {
	//	for {
	//		x := <-m
	//		m <- x
	//	}
	//}(m)

	go func(m chan []int, m2 chan []int) {
		var tpl []int
		for {
			select {
			case ok := <-m:
				fmt.Println(ok)
				tpl = ok
			default:
				fmt.Println(tpl)
			}
		}
	}(m, m2)
	for {

	}
}
