package component

type Script interface {
	OnLoad()
	Start()
	//Update(screen *ebiten.Image) (err error)
}
