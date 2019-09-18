package prefab

import (
	"game_fly/core/component"
	"game_fly/core/data"
	"game_fly/core/sprite"
)

// todo 预制体分组

//var Prefab = make(map[string]component.SpriteComponent)
//外层为组名称，内部map为多个预制体数据
//type prefabData map[string]component.SpriteComponent

//var Prefab = map[string]map[string]component.SpriteComponent{}

//设置预制体
func AddPrefab(prefab component.SpriteComponent, groupName string) {
	prefab.OnLoad()

	flg := false
	for _, v := range sprite.Sprites {
		if v.GroupName == groupName {
			flg = true
			v.Prefab[prefab.GetId()] = prefab
		}
	}

	if flg == false {
		newPrefab := make(map[string]component.SpriteComponent)
		newPrefab[prefab.GetId()] = prefab
		sprite.Sprites = append(sprite.Sprites, data.SpriteGroup{GroupName: groupName, Prefab: newPrefab})
	}
}

//根据名称获取组里面的预制体
func GetPrefabAll(groupName string) (prefab map[string]component.SpriteComponent) {
	for _, v := range sprite.Sprites {
		if v.GroupName == groupName {
			return v.Prefab
		}
	}
	return nil
}

//根据name获取当前预制体长度
func PrefabLen() (l int) {
	for _, v := range sprite.Sprites {
		l += len(v.Prefab)
	}
	return
}

//根据name与id删除预制体
func DelPrefab(groupName, id string) {
	for _, v := range sprite.Sprites {
		if v.GroupName == groupName {
			delete(v.Prefab, id)
		}
	}
}
