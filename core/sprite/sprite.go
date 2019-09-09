package sprite

import (
	"game_fly/core/component"
)

type SpriteGroup struct {
	GroupName string
	Sprite    []component.SpriteComponent//有顺序要求
	Prefab    map[string]component.SpriteComponent
}

var (
	//Sprites = make(map[string][]component.SpriteComponent)
	Sprites []SpriteGroup
)

func AddSprite(sprite component.SpriteComponent, groupName string) {
	sprite.OnLoad()

	flg := false
	for k, v := range Sprites {
		if v.GroupName == groupName {
			flg = true
			Sprites[k].Sprite = append(v.Sprite, sprite)
		}
	}

	if flg == false {
		newSprite := []component.SpriteComponent{sprite}
		Sprites = append(Sprites, SpriteGroup{GroupName: groupName, Sprite: newSprite})
	}

}

func GetSprite(groupName, name string) (sprite component.SpriteComponent) {
	for _, v := range Sprites {
		if v.GroupName == groupName {
			for _, v2 := range v.Sprite {
				if v2.GetName() == name {
					return v2
				}
			}
		}
	}
	return nil
}

func GetSpriteAll(groupName string) (component []component.SpriteComponent) {
	for _, v := range Sprites {
		if v.GroupName == groupName {
			return v.Sprite
		}
	}
	return nil
}
