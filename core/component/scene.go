package component

type Renderer interface {
	start()
}

type Scene struct {
	Component []*GameObject
}

func (s *Scene) AddGameObject(gameObject *GameObject) {
	s.Component = append(s.Component, gameObject)
}
