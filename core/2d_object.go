package core

import (
	"github.com/SolarLune/resolv/resolv"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
	"math"
)

type SpriteComponent interface {
	OnLoad()
	Start()
	Update(screen *ebiten.Image) (err error)
	//GetResolv() (collision interface{})
	//UpdateResolv()
	GetId() (id string)
	//GetName() (name string)
	//SetScreen(image *ebiten.Image)
	//GetComponent() (components []SpriteComponent)
	//GetVisible() (b bool)
	//SetVisible(b bool)
	//GetDestroy() (b bool)
	//SetDestroy(b bool)
	GetPosition() (x, y, w, h float64)
	GetChild() (child []SpriteComponent)
	GetPrefab() (prefab []SpriteComponent)
	//DelPrefab(id string)
	UpdateWrap(screen *ebiten.Image)
	GetAction() (LR, UD int)
	Move(x, y int32)
}

type Object2d struct {
	Id        string //唯一标识
	Name      string //变量名称
	ScaleW    float64
	ScaleH    float64
	SkewX     float64
	SkewY     float64
	Timing    int
	Destroy   bool              //销毁
	Material  *ebiten.Image     //材质
	Component []SpriteComponent //组件
	Prefab    []SpriteComponent //预制体
	Child     []SpriteComponent
}

//获取移动方向
func (o *Object2d) Input() {

}

func (o *Object2d) UpdateTiming() {
	o.Timing++
	if o.Timing > 1000 {
		o.Timing = 0
	}
}
func (o *Object2d) OnLoad() {

}
func (o *Object2d) Start() {

}
func (o *Object2d) Update(screen *ebiten.Image) (err error) {
	return
}

func (s *Object2d) GetPrefab() (prefab []SpriteComponent) {
	return s.Prefab
}
func (s *Object2d) GetChild() (child []SpriteComponent) {
	return s.Child
}

type Rectangle struct {
	Object2d
	resolv.Rectangle

	//Visible   bool //隐藏
	//Collision *resolv.Rectangle
}

func (o *Rectangle) SetScaleWH(scaleW, scaleH float64) {
	o.ScaleW, o.ScaleH = scaleW, scaleH
	o.W, o.H = int32(float64(o.W)*scaleW), int32(float64(o.H)*scaleH)
}

func (o *Rectangle) SetSkewXY(skewX, skewY float64) {
	o.SkewX, o.SkewY = skewX, skewY
}

func (o *Rectangle) GetPosition() (x, y, w, h float64) {
	return float64(o.X), float64(o.Y), float64(o.W), float64(o.H)
}

type Circle struct {
	resolv.Circle
	Object2d
}

type Line struct {
	resolv.Line
	Object2d
}

type RectangleSprite struct {
	Rectangle
	LR         int //左右
	UD         int //上下
	OpenCamera bool
}

func (o *RectangleSprite) Update(screen *ebiten.Image) (err error) {
	//left
	if inpututil.IsKeyJustReleased(ebiten.KeyLeft) {
		o.LR = 0
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyLeft) {
		o.LR = -1
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		if v := inpututil.KeyPressDuration(ebiten.KeyLeft); 0 < v {
			if v == 1 || v%1 == 0 {

			}
		}
	}

	//right
	if inpututil.IsKeyJustReleased(ebiten.KeyRight) {
		o.LR = 0
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyRight) {
		o.LR = 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		if v := inpututil.KeyPressDuration(ebiten.KeyRight); 0 < v {
			if v == 1 || v%1 == 0 {
			}
		}
	}

	//up
	if inpututil.IsKeyJustReleased(ebiten.KeyUp) {
		o.UD = 0
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyUp) {
		o.UD = -1
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		if v := inpututil.KeyPressDuration(ebiten.KeyUp); 0 < v {
			if v == 1 || v%1 == 0 {

			}
		}
	}
	//down
	if inpututil.IsKeyJustReleased(ebiten.KeyDown) {
		o.UD = 0
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyDown) {
		o.UD = 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		if v := inpututil.KeyPressDuration(ebiten.KeyDown); 0 < v {
			if v == 1 || v%1 == 0 {

			}
		}
	}

	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Skew(o.SkewX, o.SkewY)
	opts.GeoM.Scale(o.ScaleW, o.ScaleH)
	opts.GeoM.Translate(float64(o.X), float64(o.Y))
	screen.DrawImage(o.Material, opts)
	return
}

func (o *RectangleSprite) UpdateWrap(screen *ebiten.Image) {
	o.Update(screen)
}

func (o *RectangleSprite) OnLoad() {

}

func (r *RectangleSprite) Start() {
}

func (s *RectangleSprite) GetResolv() (collision interface{}) {
	//collision = s.Collision
	return
}

func (s *RectangleSprite) SetMaterial(img *ebiten.Image) {
	s.Material = img
	w, h := s.Material.Size()
	s.W = int32(math.Round(float64(w)))
	s.H = int32(math.Round(float64(h)))

	//fmt.Println(s.W, s.H)
	//s.SetWH(math.Round(float64(w)), math.Round(float64(h)))
	//s.SetScale(1, 1)
}
func (s *RectangleSprite) GetMaterial() (img *ebiten.Image) {
	return s.Material
}

func (s *RectangleSprite) GetName() (name string) {
	return s.Name
}

func (s *RectangleSprite) GetId() (id string) {
	return s.Id
}

func (s *RectangleSprite) UpdateResolv() {
	//if s.Collision != nil {
	//	s.Collision.SetXY(int32(s.X), int32(s.Y))
	//	s.Collision.W = int32(s.W)
	//	s.Collision.H = int32(s.H)
	//}
}

func (s *RectangleSprite) DelPrefab(id string) {
	for k, v := range s.Prefab {
		if v.GetId() == id {
			s.Prefab = append(s.Prefab[0:k], s.Prefab[k+1:]...)
		}
	}
}

func (s *RectangleSprite) GetAction() (LR, UD int) {
	return s.LR, s.UD
}
