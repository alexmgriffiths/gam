package tiles

import "github.com/hajimehoshi/ebiten/v2"

// Leaving this as grass and not generic for sound later (and maybe anim)
type GrassTile struct {
	Image *ebiten.Image
	X     int
	Y     int
}

func NewGrassTile(image *ebiten.Image, x, y int) *GrassTile {
	return &GrassTile{
		Image: image,
		X:     x,
		Y:     y,
	}
}

func (t *GrassTile) Render(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(t.X), float64(t.Y))
	screen.DrawImage(t.Image, op)
}

func (t *GrassTile) Tick() {}
