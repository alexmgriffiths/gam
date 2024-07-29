package entities

import "github.com/hajimehoshi/ebiten/v2"

type BoundingBox struct {
	X      int
	Y      int
	Width  int
	Height int
}

type Renderable interface {
	Render(screen *ebiten.Image)
	GetY() int
	SetPosition(x, y int)
	GetBoundingBox() BoundingBox
}
