package component

import (
	"github.com/325Gerbils/go-vector"
	"log"
	"reflect"
)

var CE = &GameObject{}

type GameObject struct {
	Name      string //名称，预制体有用到
	Del       bool //删除
	Transform Transform //
	Sprite    Sprite //精灵
	Collider  Collider //碰撞对象
	//Script    Script
	IsPrefab bool //是否预制体

	Speed     float64 //速度
	Direction vector.Vector //方向

	Layer int  //层
}

func (gao *GameObject) RegisterComponent(componentVal interface{}) (component interface{}) {
	switch componentVal.(type) {
	case Transform:
		CE.Transform = componentVal.(Transform)
	case Sprite:
		CE.Sprite = componentVal.(Sprite)
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

func (gao *GameObject) OnTriggerEnter(order Render) {

}

func (gao *GameObject) GetThis() (this interface{}) {

	return
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
