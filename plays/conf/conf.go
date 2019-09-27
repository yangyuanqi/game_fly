package conf

var CC Conf

type Conf struct {
	ScenesWidth, ScenesHeight int
	Scale                     float64
	Title                     string
}

func (c *Conf) SetScenes(w, h int, s float64, title string) {
	c.ScenesWidth, c.ScenesHeight, c.Scale, c.Title = w, h, s, title
}

func (c *Conf) GetScenesWH() (w, h float64) {
	return float64(c.ScenesWidth), float64(c.ScenesHeight)
}
