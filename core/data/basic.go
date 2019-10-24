package data

import (
	"github.com/SolarLune/resolv/resolv"
	"github.com/hajimehoshi/ebiten"
	_ "image/jpeg"
	_ "image/png"
	"math"
	"strconv"
)

type Base struct {
	Id   string //唯一标识
	Name string //变量名称
	resolv.Rectangle
	//geo.Rect
	ScaleW    float64
	ScaleH    float64
	SkewX     float64
	SkewY     float64
	Destroy   bool              //销毁
	Material  *ebiten.Image     //材质
	//Component []SpriteComponent //组件
	FSM       []int             //状态
	//预制体
	//Prefab []SpriteComponent
}

type Sprite struct {
	Base
	Timing    int
	Visible   bool //隐藏
	Collision *resolv.Rectangle
	Child     []Sprite
}

func (s *Base) SetScale(sw, sh float64) {
	w, h := s.Material.Size()
	s.SetWH(float64(sw*float64(w)), float64(sh*float64(h)))
	s.ScaleW = float64(sw)
	s.ScaleH = float64(sh)
}

func (s *Base) SetWH(w, h float64) {
	s.W = int32(w)
	s.H = int32(h)
}

func (s *Base) GetName() (name string) {
	return s.Name
}

func (s *Base) GetId() (id string) {
	return s.Id
}

func (s *Base) SetFSM(key, value int) {
	if len(s.FSM) <= 0 {
		var f = make([]int, 1024) // todo 最大1024个状态
		f[key] = value
		s.FSM = f
	}
	s.FSM[key] = value
}

func (b *Base) GetPosition() (x, y, w, h float64) {
	return float64(b.X), float64(b.Y), float64(b.W), float64(b.H)
}

func (b *Base) GetPositionInt32() (x, y, w, h int32) {
	return int32(b.X), int32(b.Y), int32(b.W), int32(b.H)
}

func (s *Base) GetFSM(key int) (value int) {
	if len(s.FSM) > 0 {
		return s.FSM[key]
	}
	return 0
}

func (s *Sprite) OnLoad() {

}

func (s *Sprite) Start() {

}

func (s *Sprite) GetChild() (child []Sprite) {
	return s.Child
}

func (s *Sprite) Update(screen *ebiten.Image) (err error) {

	opts := &ebiten.DrawImageOptions{}
	//opts.GeoM.Reset()
	opts.GeoM.Skew(s.SkewX, s.SkewY)
	opts.GeoM.Scale(s.ScaleW, s.ScaleH)
	opts.GeoM.Translate(float64(s.X), float64(s.Y))
	screen.DrawImage(s.Material, opts)
	return nil
}

func (s *Sprite) UpdateTiming() {
	s.Timing++
	if s.Timing > 1000 {
		s.Timing = 0
	}
}

//初始化系统属性
func (s *Sprite) Create(name string) {
	ID++
	id := strconv.Itoa(ID)
	s.Id = id
	s.Name = name
}

func (s *Sprite) GetResolv() (collision interface{}) {
	collision = s.Collision
	return
}

func (s *Sprite) SetMaterial(img *ebiten.Image) {
	s.Material = img
	w, h := s.Material.Size()
	s.SetWH(math.Round(float64(w)), math.Round(float64(h)))
	s.SetScale(1, 1)
}

//同步碰撞
func (s *Sprite) UpdateResolv() {
	//if s.Collision != nil {
	//	s.Collision.SetXY(int32(s.X), int32(s.Y))
	//	s.Collision.W = int32(s.W)
	//	s.Collision.H = int32(s.H)
	//}
}

//func (s *Sprite) AddComponent(component SpriteComponent) {
//	s.Component = append(s.Component, component)
//	component.OnLoad()
//	component.Start()
//}
//
//func (s *Sprite) GetComponent() (components []SpriteComponent) {
//	return s.Component
//}

func (s *Sprite) GetVisible() (b bool) {
	return s.Visible
}
func (s *Sprite) SetVisible(b bool) () {
	s.Visible = b
}
func (s *Sprite) GetDestroy() (b bool) {
	return s.Destroy
}
func (s *Sprite) SetDestroy(b bool) () {
	s.Visible = b
}

func (s *Sprite) SetSkewXY(skewX, skewY float64) {
	s.SkewX = skewX
	s.SkewY = skewY
}

//func (s *Sprite) AddPrefab(sprite SpriteComponent) {
//	s.Prefab = append(s.Prefab, sprite)
//}

//func (s *Sprite) DelPrefab(id string) {
//	for k, v := range s.Prefab {
//		if v.GetId() == id {
//			s.Prefab = append(s.Prefab[0:k], s.Prefab[k+1:]...)
//		}
//	}
//}

//func (s *Sprite) forDelPrefabRoot(sprite []Sprite, id string) {
//	for _, v := range sprite {
//		for _, v := range v.GetPrefab() {
//			if v.GetId() == id {
//				v.DelPrefab(id)
//			}
//		}
//		s.forDelPrefabRoot(v.GetChild(), id)
//	}
//}

//func (s *Sprite) DelPrefabRoot(id string) {
//	scenes := core.GetScene(core.ScenesIng)
//	for _, v := range scenes.GetPrefab() {
//		if v.GetId() == id {
//			v.DelPrefab(id)
//		}
//	}
//	s.forDelPrefabRoot(scenes.GetChild(), id)
//}

//func (s *Sprite) GetPrefab() (prefab []SpriteComponent) {
//	return s.Prefab
//}
