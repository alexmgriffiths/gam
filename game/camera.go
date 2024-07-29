package game

import "github.com/alexmgriffiths/gam/entities/player"

type Camera struct {
	X         int
	Y         int
	MaxX      int
	MaxY      int
	smoothing float64
}

func NewCamera(maxX, maxY int) *Camera {
	return &Camera{
		X:         0,
		Y:         0,
		MaxX:      maxX,
		MaxY:      maxY,
		smoothing: 0.1,
	}
}

func (c *Camera) Tick(player *player.Player) {
	// Camera logic: Center the camera on the player
	targetCameraX := player.X - screenWidth/2
	targetCameraY := player.Y - screenHeight/2

	if targetCameraX < 0 {
		targetCameraX = 0
	} else if targetCameraX > c.MaxX {
		targetCameraX = c.MaxX
	}

	if targetCameraY < 0 {
		targetCameraY = 0
	} else if targetCameraY > c.MaxY {
		targetCameraY = c.MaxY
	}

	// Apply the clamped target position to the camera
	c.X = targetCameraX
	c.Y = targetCameraY
	player.CameraX = targetCameraX
	player.CameraY = targetCameraY
}
