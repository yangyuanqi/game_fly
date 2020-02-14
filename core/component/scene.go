package component

import (
	"github.com/hajimehoshi/ebiten"
)

type Renderer interface {
	start()
}

type Render interface {
	Update(screen *ebiten.Image) (err error)
	GetCollider() (collider *Collider)
	OnCollisionEnter(order Render)
	OnTriggerEnter(order Render)
	//Destroy()
	GetGameObject() (gameObject *GameObject)
	GetThis() (this interface{})
}

type Scene struct {
	GameObjects []Render
	//Prefabs     map[string][]Render
}

func (s *Scene) AddGameObject(gameObject Render) {
	s.GameObjects = append(s.GameObjects, gameObject)
}
//
//func (s *Scene) AddPrefab(prefab Render, prefabName string) {
//	flg := false
//	for k, _ := range s.Prefabs {
//		if k == prefabName {
//			flg = true
//			s.Prefabs[prefabName] = append(s.Prefabs[prefabName], prefab)
//			break
//		}
//	}
//	if flg == false {
//		s.Prefabs = make(map[string][]Render)
//		s.Prefabs[prefabName] = append(s.Prefabs[prefabName], prefab)
//	}
//}
//func (p *Scene) GetPrefab(prefabName string) (prefabs []Render) {
//	return p.Prefabs[prefabName]
//}
//func (s *Scene) PrefabsLen() int {
//	var l int
//	for _, v := range s.Prefabs {
//		l += len(v)
//	}
//	return l
//}
