package main

import "fmt"

func main() {
	var s = []int{2, 1}

	for k, v := range s {
		if k < len(s)-1{
			fmt.Println(k, v, s[:k], s[k+1:])
			s = append(s[:k], s[k+1:]...)
		}


		fmt.Println(s)
	}
}
