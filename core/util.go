package core

import (
	"bytes"
	"fmt"
	"game_fly/core/data"
	"github.com/hajimehoshi/ebiten"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"strconv"
)

func GetXYOfDegree(x, y float64, degree float64, speed float64) (mx, my float64) {
	if (degree < 90 && degree > 0) || (degree > 270 && degree < 360) {
		if degree < 90 {
			x += speed
		} else {
			x -= speed
		}
		y -= 1
		return x, y
	} else {
		if degree < 180 {
			x += speed
		} else {
			x -= speed
		}
		y += 1
		return x, y
	}
	return
}

func ScaliEq(a1, a2 float64) (s float64) {
	if a1 > a2 {
		s, _ = strconv.ParseFloat(fmt.Sprintf("%.4f", a2/a1), 64)
	} else {
		s, _ = strconv.ParseFloat(fmt.Sprintf("%.4f", a1/a2), 64)
	}
	return
}

func Byte2Image(b []byte) (retImg *ebiten.Image) {
	img, _, err := image.Decode(bytes.NewReader(b))
	if err != nil {
		log.Fatal("Byte2Image:", err)
	}
	retImg, _ = ebiten.NewImageFromImage(img, ebiten.FilterDefault)
	return
}

func LoadSprite(img []byte) (sprite *data.Sprite) {
	s := &data.Sprite{}
	s.SetMaterial(Byte2Image(img))
	return s
}

func Going(role, scenes SpriteComponent, mfunc string) {
	rx, _, rw, _ := role.GetPosition()
	rLR, _ := role.GetAction()

	//sx, _, _, _ := scenes.GetPosition()

	//switch mfunc {
	//case "move":
	if rLR == 1 {
		if rx <= 320-CameraGlobal.LW-rw {
			role.Move(2, 0)
		} else {
			//if sx >= 0 {
				scenes.Move(-2, 0)
			//}
		}
	}

	if rLR == -1 {
		if rx >= CameraGlobal.LW {
			role.Move(-2, 0)
		} else {
			//if sx <= 0 {
				scenes.Move(2, 0)
			//}
		}
	}
	//}

	//if rUD == 1 {
	//	r.Move(0, 2)
	//}
	//if rUD == -1 {
	//	r.Move(0, -2)
	//}
}
