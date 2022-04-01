package core

import (
	"bytes"
	"fmt"
	"game_fly/core/component"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"log"
)

//var PrefabL sync.Mutex

//var ID int

type GameObject struct {
	Name      string //名称
	destroy   bool   //销毁
	Component []Component
	Children  []GameObjectInterface
}

//添加组件
func (g *GameObject) AddComponent(components ...Component) {
	g.Component = append(g.Component, components...)
}

// 添加子对象
func (g *GameObject) AddChildren(gameObject ...GameObjectInterface) {
	g.Children = append(g.Children, gameObject...)
}

// 设置父对象
func (g *GameObject) SetParent() {

}

func (g *GameObject) Update(dt float64) error {
	return nil
}

func (g *GameObject) Draw(screen *ebiten.Image) {
	options := &ebiten.DrawImageOptions{}
	options.GeoM.Scale(g.Node().Scale.X, g.Node().Scale.Y)
	options.GeoM.Translate(g.Node().Position.X, g.Node().Position.Y)
	screen.DrawImage(g.Sprite().SpriteFrame, options)

	msg := fmt.Sprintf(`TPS: %0.2f
FPS: %0.2f
Num of sprites: %d
`, ebiten.CurrentTPS(), ebiten.CurrentFPS(), len(g.Children))
	ebitenutil.DebugPrint(screen, msg)
}

func (g *GameObject) Destroy() {
	g.destroy = true
}

func (g *GameObject) GetDestroy() bool {
	return g.destroy
}
func (s *GameObject) GetChildren() *[]GameObjectInterface {
	return &s.Children
}

//func (g *GameObject) SetChildren(newChildren []GameObjectInterface) {
//	g.Children = newChildren
//}

//func (g *GameObject) FilterChildren() {
//	for k, v := range g.Children {
//		if v.GetDestroy() == true {
//			g.Children = append(g.Children[:k], g.Children[:k+1]...)
//		}
//	}
//}

func (s *GameObject) GetComponent() []Component {
	return s.Component
}

func (s *GameObject) Init() {
}
func (s *GameObject) SetPosition(vec2 component.Vec2) {
	s.Node().Position.X = vec2.X
	s.Node().Position.Y = vec2.Y
}
func (s *GameObject) Node() (node *component.Node) {
	for _, v := range s.Component {
		if _, ok := v.(*component.Node); ok {
			return v.(*component.Node)
		}
	}
	return nil
}

func (s *GameObject) Sprite() (node *component.Sprite) {
	for _, v := range s.Component {
		if _, ok := v.(*component.Sprite); ok {
			return v.(*component.Sprite)
		}
	}
	return nil
}

//func (s *GameObject) GetChildren() []component.GameObjectInterface {
//	return s.Children
//}

type Scene struct {
	GameObject
}

//func (s *Base) SetWH(w, h float64) {
//	s.W = int32(w)
//	s.H = int32(h)
//}
//func (s *Base) SetXY(x, y int32) {
//	s.X = x
//	s.Y = y
//}
//
//func (s *Base) GetName() (name string) {
//	return s.Name
//}
//
//func (s *Base) GetId() (id string) {
//	return s.Id
//}
//
//func (s *Base) SetFSM(key, value int) {
//	if len(s.FSM) <= 0 {
//		var f = make([]int, 1024) // todo 最大1024个状态
//		f[key] = value
//		s.FSM = f
//	}
//	s.FSM[key] = value
//}
//
//func (b *Base) GetPosition() (x, y, w, h float64) {
//	return float64(b.X), float64(b.Y), float64(b.W), float64(b.H)
//}
//
//func (b *Base) GetPositionInt32() (x, y, w, h int32) {
//	return int32(b.X), int32(b.Y), int32(b.W), int32(b.H)
//}
//
//func (s *Base) GetFSM(key int) (value int) {
//	if len(s.FSM) > 0 {
//		return s.FSM[key]
//	}
//	return 0
//}

//func (s Sprite) OnLoad() {
//
//}
//
//func (s *Sprite) Start() {
//
//}

//
//func (s *Sprite) SetScale(sw, sh float64) {
//	w, h := s.Material.Size()
//	s.SetWH(float64(sw*float64(w)), float64(sh*float64(h)))
//	s.ScaleW = float64(sw)
//	s.ScaleH = float64(sh)
//}
//
//func (s *Sprite) GetChild() (child []Sprite) {
//	return s.Child
//}
//
//func (s *Sprite) Update(screen *ebiten.Image) (err error) {
//	s.Move(s)
//	opts := &ebiten.DrawImageOptions{}
//	//opts.GeoM.Reset()
//	opts.GeoM.Skew(s.SkewX, s.SkewY)                //歪斜
//	opts.GeoM.Scale(s.ScaleW, s.ScaleH)             //比例
//	opts.GeoM.Translate(float64(s.X), float64(s.Y)) //位置
//	//opts.GeoM.Rotate()//旋转
//	screen.DrawImage(s.Material, opts)
//	return nil
//}
//
//func (s *Sprite) UpdateTiming() {
//	s.Timing++
//	if s.Timing > 1000 {
//		s.Timing = 0
//	}
//}
//
////初始化系统属性
//func (s *Sprite) Create(name string) {
//	ID++
//	id := strconv.Itoa(ID)
//	s.Id = id
//	s.Name = name
//}
//
//func (s *Sprite) GetResolv() (collision interface{}) {
//	collision = s.Collision
//	return
//}
//
//func (s *Sprite) SetMaterial(img *ebiten.Image) {
//	s.Material = img
//	w, h := s.Material.Size()
//	s.SetWH(math.Round(float64(w)), math.Round(float64(h)))
//	s.SetScale(1, 1)
//	s.Move = func(a *Sprite) {}
//	// s.SetXY(0, 0)
//}
//
////销毁
//func (s *Sprite) Destroy() bool {
//	s.Base.Destroy = true
//	return true
//}
//
//func (s *Sprite) AttrDel() bool {
//	return s.Base.Destroy
//}

// func (s *Sprite) Move(x, y float64) {

// }

//图片转换
func Byte2Image(b []byte) (retImg *ebiten.Image) {
	img, _, err := image.Decode(bytes.NewReader(b))
	if err != nil {
		log.Fatal("Byte2Image:", err)
	}
	retImg = ebiten.NewImageFromImage(img)
	return
}

//同步碰撞
//func (s *Sprite) UpdateResolv() {
//	//if s.Collision != nil {
//	//	s.Collision.SetXY(int32(s.X), int32(s.Y))
//	//	s.Collision.W = int32(s.W)
//	//	s.Collision.H = int32(s.H)
//	//}
//}

//func (s *Sprite) AddComponent(component SpriteComponent) {
//	s.Component = append(s.Component, component)
//	component.OnLoad()
//	component.Start()
//}
//
//func (s *Sprite) GetComponent() (components []SpriteComponent) {
//	return s.Component
//}

//func (s *Sprite) GetVisible() (b bool) {
//	return s.Visible
//}
//func (s *Sprite) SetVisible(b bool) {
//	s.Visible = b
//}
//
//func (s *Sprite) SetSkewXY(skewX, skewY float64) {
//	s.SkewX = skewX
//	s.SkewY = skewY
//}

//func (s *Sprite) KeyY() (p int) {
//	if ebiten.IsKeyPressed(ebiten.KeyUp) {
//		if v := inpututil.KeyPressDuration(ebiten.KeyUp); 0 < v {
//			p = -1
//		}
//	}
//	if ebiten.IsKeyPressed(ebiten.KeyDown) {
//		if v := inpututil.KeyPressDuration(ebiten.KeyDown); 0 < v {
//			p = 1
//		}
//	}
//	return p
//}
//
//func (s *Sprite) KeyX() (p int) {
//	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
//		if v := inpututil.KeyPressDuration(ebiten.KeyLeft); 0 < v {
//			p = -1
//		}
//	}
//	if ebiten.IsKeyPressed(ebiten.KeyRight) {
//		if v := inpututil.KeyPressDuration(ebiten.KeyRight); 0 < v {
//			p = 1
//		}
//	}
//	return p
//}
//
//func (s *Sprite) KeySpace() bool {
//	if ebiten.IsKeyPressed(ebiten.KeySpace) {
//		if v := inpututil.KeyPressDuration(ebiten.KeySpace); 0 < v {
//			if v == 1 {
//				return true
//			}
//		}
//	}
//	return false
//}

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

//type SpriteGroup struct {
//	GroupName string
//	Sprite    []SpriteComponent //有顺序要求
//	Prefab    map[string]SpriteComponent
//	Ui        []UiType
//}

type UiType interface {
	OnLoad()
	Start()
	Update(screen *ebiten.Image) (err error)
	GetName() (name string)
}

type Core struct {
	Width, Height int
	Scale         float64
	title         string
}

func (c *Core) GetWH() (w, h int) {
	return c.Width, c.Height
}

//type Prefab struct {
//	Sprite
//}
