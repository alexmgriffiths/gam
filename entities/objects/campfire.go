package objects

import (
	"github.com/alexmgriffiths/gam/entities"
	"github.com/alexmgriffiths/gam/entities/tiles"
	"github.com/hajimehoshi/ebiten/v2"
)

type Campfire struct {
	X, Y             int
	Type             int
	AnimationFrame   int
	AnimationCounter int
}

func NewCampfire(campfireType, x, y int) *Campfire {
	return &Campfire{
		X:                x,
		Y:                y,
		Type:             campfireType,
		AnimationFrame:   0,
		AnimationCounter: 0,
	}
}

func (c *Campfire) GetBoundingBox() entities.BoundingBox {
	return entities.BoundingBox{}
}

func (c *Campfire) GetY() int {
	return c.Y
}

func (c *Campfire) SetPosition(x, y int) {
	c.X = x
	c.Y = y
}

func (c *Campfire) Render(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(c.X), float64(c.Y))
	screen.DrawImage(c.getImage(), op)
}

func (c *Campfire) Tick() {

}

func (c *Campfire) getImage() *ebiten.Image {
	tileset := &tiles.Tilemap{}
	width := 32
	height := 32
	frameDuration := 5
	maxFrames := 6

	if c.Type == 0 {
		tileset = tiles.NewTilemap("assets/Tilemap/campfire_lit.png")
	} else {
		tileset = tiles.NewTilemap("assets/Tilemap/campfire_out.png")
		height = 64
		frameDuration = 5
	}
	c.AnimationCounter++
	if c.AnimationCounter >= frameDuration {
		c.AnimationCounter = 0
		c.AnimationFrame++
		if c.AnimationFrame >= maxFrames {
			c.AnimationFrame = 0
		}
	}
	return tileset.GetTileImage(c.AnimationFrame, 0, width, height)
}
