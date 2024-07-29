package utils

import "github.com/alexmgriffiths/gam/entities"

func CheckCollision(box1, box2 entities.BoundingBox) bool {
	return box1.X < box2.X+box2.Width &&
		box1.X+box1.Width > box2.X &&
		box1.Y < box2.Y+box2.Height &&
		box1.Y+box1.Height > box2.Y
}
