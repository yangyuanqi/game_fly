package main

import "fmt"

func main() {
	var s = []int{2,1,0}

	for k, v := range s {
		if k < len(s){
			fmt.Println(k, v, s[:k], s[k+1:])
			s = append(s[:k], s[k+1:]...)
		}


		fmt.Println(s)
	}
}
