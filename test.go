package main

import "fmt"

func main() {
	defer fmt.Println("主协程马上结束")
	ch2 := make(chan string)
	go func() {
		defer fmt.Println("子协程马上结束")
		/*		for i := 0; i < 5; i++ {
					fmt.Println("i = ", i)
				}*/
		ch2 <- "子协程结束"
	}()

	<-ch2

	//fmt.Println(str)
	//fmt.Println("主协程开始")
}
