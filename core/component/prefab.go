package component

import (
	"github.com/hajimehoshi/ebiten"
)

//只存储坐标和个别对象改变的信息，不存储多个预制体，节约资源

type Prefab struct {
	GameObject
}

func (p *Prefab) Draw(scenesGameObjects []Render, screen *ebiten.Image, prefabName string) {
	for _, v := range scenesGameObjects {
		if v.GetGameObject().IsPrefab && v.GetGameObject().Name == prefabName {
			v.GetCollider().Move(v.GetGameObject().Direction.SetMag(v.GetGameObject().Speed))
			opts := &ebiten.DrawImageOptions{}
			opts.GeoM.Scale(v.GetGameObject().Transform.ScaleX, v.GetGameObject().Transform.ScaleY)
			opts.GeoM.Translate(v.GetCollider().Transform.Position.X, v.GetCollider().Transform.Position.Y)
			p.Transform.Position = p.Collider.Transform.Position
			screen.DrawImage(v.GetGameObject().Sprite.Img, opts)
		}
	}
}
