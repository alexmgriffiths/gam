package objects

import (
	"github.com/alexmgriffiths/gam/entities"
	"github.com/hajimehoshi/ebiten/v2"
)

type Well struct {
	Image, Normal *ebiten.Image
	X, Y          int
	Width, Height int
}

func NewWell(image *ebiten.Image, normal *ebiten.Image, x, y, width, height int) *Well {
	return &Well{
		Image:  image,
		Normal: normal,
		X:      x,
		Y:      y,
		Width:  width,
		Height: height,
	}
}

func (t *Well) GetLightParameters() (bool, entities.LightEmitter) {
	return false, entities.LightEmitter{}
}

func (w *Well) Render(screen *ebiten.Image, normalBuffer *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(w.X), float64(w.Y))
	screen.DrawImage(w.Image, op)
	normalBuffer.DrawImage(w.Normal, op)
}

func (w *Well) Tick() {

}

func (w *Well) GetY() int {
	return w.Y + 10
}

func (w *Well) SetPosition(x, y int) {
	w.X = x
	w.Y = y
}

func (w *Well) GetBoundingBox() entities.BoundingBox {
	return entities.BoundingBox{
		X:      w.X - 2,
		Y:      w.Y + 16,
		Width:  (w.Width * 16) + 7,
		Height: (w.Height * 16) / 4,
	}
}
