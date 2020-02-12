package component

import "github.com/hajimehoshi/ebiten"

type Renderer interface {
	start()
}

type Render interface {
	Update(screen *ebiten.Image) (err error)
	GetCollider() (collider *Collider)
	OnCollisionEnter(order Render)
	//Destroy()
	GetGameObject()(gameObject *GameObject)
}

type Scene struct {
	GameObjects []Render
}

func (s *Scene) AddGameObject(gameObject Render) {
	s.GameObjects = append(s.GameObjects, gameObject)
}
