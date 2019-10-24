package plays

//
//type Bullet struct {
//	data.Sprite
//	roleNode *Role
//	gameNode Game
//	object   data.SpriteComponent
//}
//
//const (
//	BulletType = iota
//)
//
//func NewBullet(img *ebiten.Image, object data.SpriteComponent) (bullet *Bullet) {
//	b := &Bullet{}
//	b.Sprite.Create("bullet")
//	//获取节点
//	b.object = object
//	b.gameNode = *core.GetScene("game").(*Game)
//	//贴图
//	b.Material = img
//	//初始化属性
//	b.SetScale(0.2, 0.2)
//
//	//添加碰撞
//	b.Collision = resolv.NewRectangle(b.GetPositionInt32())
//	return b
//}

//func (b *Bullet) OnLoad() {
//	if b.GetFSM(BulletType) == 1 {
//		//b.roleNode = sprite.GetSprite("game", "role").(*Role)
//		//b.SetXY(b.roleNode.X+int32(math.Ceil(float64(b.roleNode.W-b.W)/2)), b.roleNode.Y)
//	}
//
//	if b.GetFSM(BulletType) == 2 {
//		x, y, w, h := b.object.GetPosition()
//		b.SetXY(int32(int32(x)+(int32(w)-b.W)/2), int32(y+h))
//	}
//}
//
//func (b *Bullet) Start() {
//
//}
//
//func (b *Bullet) Update(screen *ebiten.Image) (err error) {
//	//到达边界删除b
//	//if float64(b.Y) > float64(b.gameNode.H) || b.Y < 0 {
//	//	prefab.DelPrefab(b.Id)
//	//}
//	//
//	//if b.GetFSM(BulletType) == 1 {
//	//	b.Move(0, -5)
//	//}
//	//if b.GetFSM(BulletType) == 2 {
//	//	b.Move(0, 3)
//	//}
//
//	//roleBullet := prefab.GetPrefabAll("roleBullet")
//	//enemy := prefab.GetPrefabAll("enemy")
//	//for _, v := range roleBullet {
//	//	for _, v2 := range enemy {
//	//		colliding := v.(*Bullet).IsColliding(v2.(*Enemy))
//	//		if colliding {
//	//			fmt.Println("碰撞1")
//	//			v2.(*Enemy).SetFSM(SouShang, v2.(*Enemy).GetFSM(SouShang)+1)
//	//			//v.SetDestroy(true)
//	//			prefab.DelPrefab(v.GetId())
//	//		}
//	//	}
//	//}
//
//	if ebiten.IsDrawingSkipped() {
//		return nil
//	}
//
//	b.Sprite.Update(screen)
//	return
//}
