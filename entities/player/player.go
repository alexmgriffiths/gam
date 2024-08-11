package player

import (
	"github.com/alexmgriffiths/gam/entities"
	"github.com/alexmgriffiths/gam/entities/tiles"
	"github.com/alexmgriffiths/gam/utils"
	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct {
	Tilemap          *tiles.Tilemap
	Normalmap        *tiles.Tilemap
	X                int
	Y                int
	VelocityX        int
	VelocityY        int
	CameraX          int
	CameraY          int
	IsMoving         bool
	Rotation         int
	AnimationFrame   int
	AnimationCounter int
}

func (p *Player) GetY() int {
	return p.Y
}

func (p *Player) SetPosition(x, y int) {
	p.X = x
	p.Y = y
}

func (p *Player) GetHeight() int {
	return 16
}

func (p *Player) GetBoundingBox() entities.BoundingBox {
	return entities.BoundingBox{
		X:      p.X + 4,
		Y:      p.Y + 8,
		Width:  8,
		Height: 8,
	}
}

func (p *Player) getPlayerImage() (*ebiten.Image, *ebiten.Image) {
	frameDuration := 10
	if p.IsMoving {
		p.AnimationCounter++
		if p.AnimationCounter >= frameDuration {
			p.AnimationCounter = 0
			p.AnimationFrame++
			if p.AnimationFrame >= 3 {
				p.AnimationFrame = 0
			}
		}
	} else {
		p.AnimationFrame = 0
	}

	if p.Rotation == 1 {
		return p.Tilemap.GetTileImage(25, p.AnimationFrame, 16, 16), p.Normalmap.GetTileImage(25, p.AnimationFrame, 16, 16)
	} else if p.Rotation == 2 {
		return p.Tilemap.GetTileImage(26, p.AnimationFrame, 16, 16), p.Normalmap.GetTileImage(26, p.AnimationFrame, 16, 16)
	} else if p.Rotation == 3 {
		return p.Tilemap.GetTileImage(23, p.AnimationFrame, 16, 16), p.Normalmap.GetTileImage(23, p.AnimationFrame, 16, 16)
	}
	return p.Tilemap.GetTileImage(24, p.AnimationFrame, 16, 16), p.Normalmap.GetTileImage(24, p.AnimationFrame, 16, 16)
}

func NewPlayer(tilemap *tiles.Tilemap, normalmap *tiles.Tilemap, x, y int) *Player {
	return &Player{
		Tilemap:          tilemap,
		Normalmap:        normalmap,
		X:                x,
		Y:                y,
		CameraX:          0,
		CameraY:          0,
		IsMoving:         false,
		Rotation:         0,
		AnimationFrame:   0,
		AnimationCounter: 0,
	}
}

func (p *Player) Tick(keysPressed []ebiten.Key, objects []entities.Renderable) {
	p.IsMoving = false
	if len(keysPressed) > 0 {
		currentKeyPress := keysPressed[len(keysPressed)-1].String()
		if currentKeyPress == "S" {
			p.VelocityY = 1
			p.Rotation = 0
			p.IsMoving = true
		} else if currentKeyPress == "W" {
			p.VelocityY = -1
			p.Rotation = 1
			p.IsMoving = true
		} else if currentKeyPress == "D" {
			p.VelocityX = 1
			p.Rotation = 2
			p.IsMoving = true
		} else if currentKeyPress == "A" {
			p.VelocityX = -1
			p.Rotation = 3
			p.IsMoving = true
		} else {
			p.IsMoving = false
		}
	} else {
		p.IsMoving = false
	}

	// Update the player's position with velocity
	nextX := p.X + p.VelocityX
	nextY := p.Y + p.VelocityY

	// Check for collisions
	if !p.CheckCollision(nextX, nextY, objects) {
		p.X = nextX
		p.Y = nextY
	} else {
		p.VelocityX = 0
		p.VelocityY = 0
	}

	// Reset velocities
	p.VelocityX = 0
	p.VelocityY = 0
}

func (p *Player) CheckCollision(nextX, nextY int, objects []entities.Renderable) bool {
	playerBox := entities.BoundingBox{
		X:      (nextX + 6) - p.CameraX,
		Y:      (nextY + 12) - p.CameraY,
		Width:  5,
		Height: 4,
	}
	for _, obj := range objects {
		if obj == nil || obj == p {
			continue
		}
		if utils.CheckCollision(playerBox, obj.GetBoundingBox()) {
			return true
		}
	}
	return false
}

func (p *Player) Render(screen *ebiten.Image, normalBuffer *ebiten.Image) {
	options := &ebiten.DrawImageOptions{}
	options.GeoM.Translate(float64(p.X-p.CameraX), float64(p.Y-p.CameraY))
	playerImage, playerNormal := p.getPlayerImage()
	screen.DrawImage(playerImage, options)
	normalBuffer.DrawImage(playerNormal, options)
}

func (p *Player) GetLightParameters() (bool, entities.LightEmitter) {
	return false, entities.LightEmitter{}
}
