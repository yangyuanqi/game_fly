package component

import (
	"log"
	"reflect"
)

var CE = &GameObject{}

type GameObject struct {
	Del       bool
	Transform Transform
	Sprite    Sprite
	Script    Script
	Collider  Collider
}

func (gao *GameObject) RegisterComponent(componentVal interface{}) (component interface{}) {
	switch componentVal.(type) {
	case Transform:
		CE.Transform = componentVal.(Transform)
	case Sprite:
		CE.Sprite = componentVal.(Sprite)
	case Script:
		CE.Script = componentVal.(Script)
	case Collider:
		CE.Collider = componentVal.(Collider)
	default:
		log.Println("未知类型组件")
	}
	return CE
}

func (gao *GameObject) GetComponent(typeName string) (componentVal interface{}) {
	d := reflect.ValueOf(*CE)
	componentVal = d.FieldByName(typeName)
	return
}

func (gao *GameObject) OnCollisionEnter(order Render) {

}

func (gao *GameObject) GetCollider() (collider *Collider) {
	return &gao.Collider
}

func (gao *GameObject) Destroy() {
	gao.Del = true
}

func (gao *GameObject) GetGameObject() (gameObject *GameObject) {
	return gao
}
