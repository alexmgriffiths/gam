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

func (c *Campfire) GetLightParameters() (bool, entities.LightEmitter) {
	posX, posY := c.GetPosition()
	return true, entities.LightEmitter{
		X:         (posX + 16),
		Y:         (posY + 16),
		Intensity: 15,
		Color:     [3]float32{1.0, 0.65, 0.3},
	}
}

func (c *Campfire) GetBoundingBox() entities.BoundingBox {
	return entities.BoundingBox{}
}

func (c *Campfire) GetY() int {
	return c.Y
}

func (c *Campfire) GetPosition() (int, int) {
	return c.X, c.Y
}

func (c *Campfire) SetPosition(x, y int) {
	c.X = x
	c.Y = y
}

func (c *Campfire) Render(screen *ebiten.Image, normal *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(c.X), float64(c.Y))
	screen.DrawImage(c.getImage(), op)
	normal.DrawImage(c.getImageNormal(), op)
}

func (c *Campfire) Tick() {

}

func (c *Campfire) getImageNormal() *ebiten.Image {
	tileset := &tiles.Tilemap{}
	width := 32
	height := 32
	frameDuration := 10
	maxFrames := 6

	if c.Type == 0 {
		tileset = tiles.NewTilemap("assets/Tilemap/campfire_lit_n.png")
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

func (c *Campfire) getImage() *ebiten.Image {
	tileset := &tiles.Tilemap{}
	width := 32
	height := 32
	frameDuration := 10
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
