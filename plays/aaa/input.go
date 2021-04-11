package plays

//
//type Input struct {
//	data.Sprite
//	role *Role
//	game *Game
//}
//
//func NewInput() (input *Input) {
//	return &Input{}
//}
//
//func (i *Input) OnLoad() {
//	//i.role = sprite.GetSprite("game", "role").(*Role)
//	i.game = core.GetScene("game").(*Game)
//}
//
//func (i *Input) Start() {
//
//}
//
//func (i *Input) Update(screen *ebiten.Image) (err error) {
//	//if ebiten.IsKeyPressed(ebiten.KeyUp) {
//	//	if v := inpututil.KeyPressDuration(ebiten.KeyUp); 0 < v {
//	//		if v == 1 || v%1 == 0 {
//	//			i.role.Move(0, -5)
//	//		}
//	//	}
//	//	if i.role.Y <= 0 {
//	//		i.role.Y = 0
//	//	}
//	//}
//	//if ebiten.IsKeyPressed(ebiten.KeyDown) {
//	//	if v := inpututil.KeyPressDuration(ebiten.KeyDown); 0 < v {
//	//		if v == 1 || v%1 == 0 {
//	//			i.role.Move(0, 5)
//	//		}
//	//	}
//	//	if float64(i.role.Y) >= float64(i.game.H)-float64(i.role.H) {
//	//		i.role.Y = i.game.H - i.role.H
//	//	}
//	//}
//	//
//	////left
//	//if inpututil.IsKeyJustReleased(ebiten.KeyLeft) {
//	//	i.role.QingXieLeft = false
//	//}
//	//if inpututil.IsKeyJustPressed(ebiten.KeyLeft) {
//	//	i.role.QingXieLeft = true
//	//}
//	//if ebiten.IsKeyPressed(ebiten.KeyLeft) {
//	//	if v := inpututil.KeyPressDuration(ebiten.KeyLeft); 0 < v {
//	//		if v == 1 || v%1 == 0 {
//	//			i.role.Move(-5, 0)
//	//		}
//	//		//i.role.SetFSM(Left, v)
//	//	}
//	//	if i.role.X <= 0 {
//	//		i.role.X = 0
//	//	}
//	//}
//	//
//	////right
//	//if inpututil.IsKeyJustReleased(ebiten.KeyRight) {
//	//	i.role.QingXieRight = false
//	//}
//	//if inpututil.IsKeyJustPressed(ebiten.KeyRight) {
//	//	i.role.QingXieRight = true
//	//}
//	//if ebiten.IsKeyPressed(ebiten.KeyRight) {
//	//	if v := inpututil.KeyPressDuration(ebiten.KeyRight); 0 < v {
//	//		if v == 1 || v%1 == 0 {
//	//			i.role.Move(5, 0)
//	//		}
//	//		//i.role.SetFSM(Right, v)
//	//	}
//	//	if float64(i.role.X) >= float64(i.game.W)-float64(i.role.W) {
//	//		i.role.X = i.game.W - i.role.W
//	//	}
//	//}
//	//
//	//if ebiten.IsKeyPressed(ebiten.KeySpace) {
//	//	if v := inpututil.KeyPressDuration(ebiten.KeySpace); 0 < v {
//	//		if v == 1 {
//	//			bullet := NewBullet(BulletImg, i.role)
//	//			prefab.AddPrefab(bullet, "roleBullet")
//	//		}
//	//	}
//	//}
//
//	return nil
//}
