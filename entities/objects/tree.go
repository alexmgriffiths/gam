package objects

import (
	"github.com/alexmgriffiths/gam/entities"
	"github.com/hajimehoshi/ebiten/v2"
)

type Tree struct {
	Image  *ebiten.Image
	Normal *ebiten.Image
	X      int
	Y      int
	Width  int
	Height int
}

func (t *Tree) GetLightParameters() (bool, entities.LightEmitter) {
	return false, entities.LightEmitter{}
}

func (t *Tree) GetY() int {
	return t.Y + 20
}

func (t *Tree) SetPosition(x, y int) {
	t.X = x
	t.Y = y
}

func NewTree(image *ebiten.Image, normal *ebiten.Image, x, y, width, height int) *Tree {
	return &Tree{
		Image:  image,
		Normal: normal,
		X:      x,
		Y:      y,
		Width:  width,
		Height: height,
	}
}

func (t *Tree) Render(screen *ebiten.Image, normalBuffer *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(t.X), float64(t.Y))
	screen.DrawImage(t.Image, op)
	normalBuffer.DrawImage(t.Normal, op)
}

func (t *Tree) GetBoundingBox() entities.BoundingBox {
	return entities.BoundingBox{
		X:      t.X,
		Y:      t.Y + 24, // Only care about the bottom quater of tree
		Width:  t.Width * 16,
		Height: (t.Height * 16) / 4,
	}
}

func (t *Tree) Tick() {

}
