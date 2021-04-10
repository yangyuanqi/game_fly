package sprite

//
//var (
//	//Sprites = make(map[string][]component.SpriteComponent)
//	Sprites []data.SpriteGroup
//)
//
//func AddSprite(sprite data.SpriteComponent, groupName string) {
//	flg := false
//	for k, v := range Sprites {
//		if v.GroupName == groupName {
//			flg = true
//			Sprites[k].Sprite = append(v.Sprite, sprite)
//			break
//		}
//	}
//
//	if flg == false {
//		newSprite := []data.SpriteComponent{sprite}
//		Sprites = append(Sprites, data.SpriteGroup{GroupName: groupName, Sprite: newSprite})
//	}
//	sprite.OnLoad()
//	sprite.Start()
//}

//func GetSprite(groupName, name string) (sprite data.SpriteComponent) {
//	for _, v := range Sprites {
//		if v.GroupName == groupName {
//			for _, v2 := range v.Sprite {
//				if v2.GetName() == name {
//					return v2
//				}
//			}
//		}
//	}
//	return nil
//}

//func GetSpriteAll(groupName string) (component []data.SpriteComponent) {
//	for _, v := range Sprites {
//		if v.GroupName == groupName {
//			return v.Sprite
//		}
//	}
//	return nil
//}
