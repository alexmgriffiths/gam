package lighting

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type Lightsource struct {
	X, Y   int
	Radius int
	Image  *ebiten.Image
}

func NewLightsource(x, y, radius int) *Lightsource {
	lightSource := createLightSource(radius)
	if radius == 0 {
		return nil
	}

	return &Lightsource{
		X:      x,
		Y:      y,
		Radius: radius,
		Image:  lightSource,
	}
}

func (l *Lightsource) CarveOnto(darkness *ebiten.Image) {
	spotlightOps := &ebiten.DrawImageOptions{}

	tempLightSize := l.Image.Bounds().Size()
	tempLightWidth := tempLightSize.X
	tempLightHeight := tempLightSize.Y

	drawX := float64(l.X) - float64(tempLightWidth)/2.5
	drawY := float64(l.Y) - float64(tempLightHeight)/2.5

	spotlightOps.GeoM.Translate(float64(drawX), float64(drawY))
	//spotlightOps.Blend = ebiten.BlendSourceAtop for colors but gross
	spotlightOps.ColorScale.ScaleWithColor(color.RGBA{100, 50, 0, 255})
	spotlightOps.Blend = ebiten.BlendDestinationOut

	darkness.DrawImage(l.Image, spotlightOps)
}

func createLightSource(radius int) *ebiten.Image {
	diameter := radius * 2
	lightSource := ebiten.NewImage(diameter, diameter)
	lightSource.Fill(color.Transparent) // Start with fully transparent

	center := diameter / 2
	for y := 0; y < diameter; y++ {
		for x := 0; x < diameter; x++ {
			dx := float64(x - center)
			dy := float64(y - center)
			distance := math.Sqrt(dx*dx + dy*dy)

			if distance < float64(radius) {
				// Calculate the alpha: closer to the center = more opaque
				alpha := uint8(255 * (1 - distance/float64(radius)))
				lightSource.Set(x, y, color.RGBA{255, 255, 255, alpha})
			}
		}
	}

	return lightSource
}
