package main

import (
	"game_fly/core/component"
	"image"
)

func main() {
	//


	script := &component.Sprite{}
	script.Color = image.Black

	scene := component.Scene{}
	enemy := component.Register(script)



	scene.AddGameObject(enemy)


	//fmt.Println(component.CE)
	//fmt.Println(component.GetComponent("Sprite"))
}
