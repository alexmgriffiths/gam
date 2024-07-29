package objects

import (
	"github.com/alexmgriffiths/gam/entities"
	"github.com/hajimehoshi/ebiten/v2"
)

type Bush struct {
	Image         *ebiten.Image
	X             int
	Y             int
	Width, Height int
}

func NewBush(image *ebiten.Image, x, y, width, height int) *Bush {
	return &Bush{
		Image:  image,
		X:      x,
		Y:      y - 32,
		Width:  width,
		Height: height,
	}
}

func (b *Bush) Tick() {
	b.X++
}

func (b *Bush) Render(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(b.X), float64(b.Y))
	screen.DrawImage(b.Image, op)
}

func (b *Bush) GetY() int {
	return b.Y
}

func (b *Bush) SetPosition(x, y int) {
	b.X = x
	b.Y = y
}

func (b *Bush) GetBoundingBox() entities.BoundingBox {
	return entities.BoundingBox{
		X:      b.X,
		Y:      b.Y + 12,
		Width:  b.Width * 16,
		Height: b.Height,
	}
}
