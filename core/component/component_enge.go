package component

import (
	"log"
	"reflect"
)

var CE = &GameObject{}

type GameObject struct {
	Transform *Transform
	Sprite    *Sprite
	Script    *Script
}

func (gao *GameObject) RegisterComponent(componentVal interface{}) (component *GameObject) {
	switch componentVal.(type) {
	case *Transform:
		CE.Transform = componentVal.(*Transform)
	case *Sprite:
		CE.Sprite = componentVal.(*Sprite)
	case *Script:
		CE.Script = componentVal.(*Script)
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
