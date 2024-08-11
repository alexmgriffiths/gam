package entities

import "github.com/hajimehoshi/ebiten/v2"

type BoundingBox struct {
	X      int
	Y      int
	Width  int
	Height int
}

type Renderable interface {
	Render(screen *ebiten.Image, normalBuffer *ebiten.Image)
	GetY() int
	GetLightParameters() (bool, LightEmitter)
	SetPosition(x, y int)
	GetBoundingBox() BoundingBox
}

type LightEmitter struct {
	X, Y      int
	Intensity float32
	Color     [3]float32
}
