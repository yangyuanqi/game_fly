package plays

import (
	"bytes"
	"game_fly/core"
	"game_fly/core/prefab"
	"game_fly/core/timer"
	"game_fly/core/ui"
	"game_fly/plays/assets/images"
	"github.com/hajimehoshi/ebiten"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"math"
	"time"
)

type Role struct {
	core.Sprite
	RootNode       *Game
	Left           int
	Right          int
	Img            *ebiten.Image
	Fraction       int //分数
	StartAnimation bool
}

var RoleImg *ebiten.Image

func NewRole() (role *Role) {
	r := &Role{}
	r.Sprite.Create("role")
	r.SetScale(0.3, 0.3)
	return r
}

func (r *Role) OnLoad() {
	//角色贴图
	img, _, err := image.Decode(bytes.NewReader(images.My_1))
	if err != nil {
		log.Fatal(6, err)
	}
	RoleImg, _ = ebiten.NewImageFromImage(img, ebiten.FilterDefault)
	r.Img = RoleImg

	w, h := RoleImg.Size()
	r.SetWH(math.Round(float64(w)*r.ScaleW), math.Round(float64(h)*r.ScaleH))

	game := core.GetScene("game").(*Game)
	r.RootNode = game
	r.SetXY(math.Ceil(float64((game.W-r.W)/2)), float64(game.H))

	//自动子弹
	timer.SetTicker(time.Millisecond*200, func() {
		bullet := NewBullet(BulletImg)
		bullet.X = bullet.X + 10
		bullet.Move = func(bullet2 *Bullet) {
			bullet2.Y -= 4
		}
		prefab.AddPrefab(bullet, "roleBullet")

		bullet = NewBullet(BulletImg)
		bullet.X = bullet.X - 10
		bullet.Move = func(bullet2 *Bullet) {
			bullet2.Y -= 4
		}
		prefab.AddPrefab(bullet, "roleBullet")
	}, 0)

	r.AddComponent(NewInput())
}

func (r *Role) Start() {

}

func (r *Role) Update(screen *ebiten.Image) (err error) {
	ui.ReplaceUi("ui", "number1", ui.NewNumber(r.Fraction, 200, 10, 0.5))
	r.Drow(screen)
	return nil
}

func (r *Role) Drow(screen *ebiten.Image) {
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Reset()
	opts.GeoM.Scale(r.ScaleW, r.ScaleH)

	//飞机初始动画
	if r.StartAnimation == false {
		r.Move(0, -2)
		if r.Y < r.RootNode.H-r.H {
			r.StartAnimation = true
		}
	}

	//灰机倾斜处理
	if r.Right > 0 || r.Left > 0 {
		if r.Right <= 20 {
			opts.GeoM.Skew(0, float64(r.Right)/100)
		} else if r.Right > 20 {
			opts.GeoM.Skew(0, 0.2)
		}
		if r.Left <= 20 {
			opts.GeoM.Skew(0, -float64(r.Left)/100)
		} else if r.Left > 20 {
			opts.GeoM.Skew(0, -0.2)
		}
	}

	opts.GeoM.Translate(r.X, r.Y)
	screen.DrawImage(r.Img, opts)
}
