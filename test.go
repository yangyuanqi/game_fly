package main

import "fmt"

type Abc struct {
	data map[string]string
}

func main() {
	var Prefab []Abc
	Prefab = append(Prefab, Abc{data: map[string]string{"a": "b"}})
	for k, v := range Prefab {
		if k==0{
			v.data["c"] = "d"
		}
	}

	fmt.Println(Prefab)
}
