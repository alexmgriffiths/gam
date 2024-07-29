package objects

import (
	"github.com/alexmgriffiths/gam/entities"
	"github.com/alexmgriffiths/gam/entities/tiles"
	"github.com/hajimehoshi/ebiten/v2"
)

type Object interface {
	Tick()
	Render(screen *ebiten.Image)
	GetY() int
	SetPosition(x, y int)
	GetBoundingBox() entities.BoundingBox
}

const (
	// 0 is nothing
	TREE_1       = 1
	TREE_2       = 2
	HOUSE_0      = 3
	WELL         = 4
	BUSH_1       = 5
	CAMPFIRE_LIT = 6
	CAMPFIRE_OUT = 7
)

func NewObject(tilemap *tiles.Tilemap, objectType, x, y int) Object {
	switch objectType {
	case TREE_1:
		return &Tree{
			Image:  getTileImage(tilemap, 3, 0, 1, 2),
			X:      x,
			Y:      y - 16,
			Width:  1,
			Height: 2,
		}
	case TREE_2:
		return &Tree{
			Image:  getTileImage(tilemap, 4, 0, 1, 2),
			X:      x,
			Y:      y - 16,
			Width:  1,
			Height: 2,
		}
	case BUSH_1:
		return &Bush{
			Image:  getTileImage(tilemap, 5, 0, 1, 1),
			X:      x,
			Y:      y,
			Width:  1,
			Height: 1,
		}
	case HOUSE_0:
		return &Tree{
			Image:  getTileImage(tilemap, 0, 1, 7, 4),
			X:      x,
			Y:      y - 16,
			Width:  7,
			Height: 10,
		}
	case WELL:
		return &Well{
			Image:  getTileImageCustom(tilemap, 128, 120, 17, 25),
			X:      x,
			Y:      y - 16,
			Width:  1,
			Height: 2,
		}
	case CAMPFIRE_LIT:
		return &Campfire{
			Type: 0,
			X:    x,
			Y:    y,
		}
	case CAMPFIRE_OUT:
		return &Campfire{
			Type: 1,
			X:    x,
			Y:    y,
		}
	default:
		return nil
	}
}

func getTileImage(tilemapImage *tiles.Tilemap, x, y, width, height int) *ebiten.Image {
	return tilemapImage.GetTileImage(x, y, width*16, height*16)
}

func getTileImageCustom(tilemapImage *tiles.Tilemap, x, y, width, height int) *ebiten.Image {
	return tilemapImage.GetTileImageCustom(x, y, width, height)
}
