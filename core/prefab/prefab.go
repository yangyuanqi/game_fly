package prefab

// todo 预制体分组

//var Prefab = make(map[string]component.SpriteComponent)
//外层为组名称，内部map为多个预制体数据
//type prefabData map[string]component.SpriteComponent

//var Prefab = map[string]map[string]component.SpriteComponent{}

//设置预制体
//
//func AddPrefab(prefab data.SpriteComponent, groupName string) {
//	//data.PrefabL.Lock()
//	flg := false
//	for _, v := range sprite.Sprites {
//		if v.GroupName == groupName {
//			flg = true
//			//v.Prefab[prefab.GetId()] = prefab
//			//v.Prefab.Store(prefab.GetId(), prefab)
//
//			v.Prefab[prefab.GetId()] = prefab
//
//			break
//		}
//	}
//
//	if flg == false {
//		newPrefab := make(map[string]data.SpriteComponent)
//		//m := <-newPrefab
//
//		newPrefab[prefab.GetId()] = prefab
//
//		//newPrefab <- m
//		//var newPrefab sync.Map
//
//		//newPrefab.Store(prefab.GetId(), prefab)
//
//		sprite.Sprites = append(sprite.Sprites, data.SpriteGroup{GroupName: groupName, Prefab: newPrefab})
//
//	}
//	//data.PrefabL.Unlock()
//	prefab.OnLoad()
//	prefab.Start()
//}
//
////根据名称获取组里面的预制体
//func GetPrefabAll(groupName string) (prefab map[string]data.SpriteComponent) {
//	for _, v := range sprite.Sprites {
//		if v.GroupName == groupName {
//			return v.Prefab
//		}
//	}
//	return nil
//}
//
////根据name获取当前预制体长度
//func forChildPrefabLen(child []data.Sprite) (l int) {
//	var l1 int
//	if len(child) > 0 {
//		for _, v := range child {
//			l1 = len(v.GetPrefab())
//			l2 := forChildPrefabLen(v.GetChild())
//			l1 = l1 + l2
//		}
//	}
//	return l1
//}
//
//func PrefabLen(scenesName string) (l int) {
//	scenes := core.GetScene(scenesName)
//	l1 := len(scenes.GetPrefab())
//	l2 := forChildPrefabLen(scenes.GetChild())
//	return l1 + l2
//	//for _, v := range sprite.Sprites {
//	//	var i int
//	//	//data.PrefabL.Lock()
//	//	for range v.Prefab {
//	//		i++
//	//	}
//	//	//data.PrefabL.Unlock()
//	//	len += i
//	//}
//	return
//}
//
////根据name与id删除预制体
//func DelPrefab(id string) {
//	//go func(groupName, id string) {
//	for _, v := range sprite.Sprites {
//		delete(v.Prefab, id)
//
//		//if v.GroupName == groupName {
//		//data.PrefabL.Lock()
//		//delete(v.Prefab, id)
//		//data.PrefabL.Unlock()
//		//v.Prefab.Delete(id)
//		//}
//	}
//	//}(groupName, id)
//}
