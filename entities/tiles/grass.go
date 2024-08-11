package tiles

import "github.com/hajimehoshi/ebiten/v2"

// Leaving this as grass and not generic for sound later (and maybe anim)
type GrassTile struct {
	Image      *ebiten.Image
	Normal     *ebiten.Image
	X          int
	Y          int
	Brightness float64
}

func NewGrassTile(image *ebiten.Image, normal *ebiten.Image, x, y int) *GrassTile {
	return &GrassTile{
		Image:      image,
		Normal:     normal,
		X:          x,
		Y:          y,
		Brightness: 0,
	}
}
func (t *GrassTile) Render(screen *ebiten.Image, normalBuffer *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(t.X), float64(t.Y))
	screen.DrawImage(t.Image, op)
	normalBuffer.DrawImage(t.Normal, op)
}

func (t *GrassTile) Tick() {}
