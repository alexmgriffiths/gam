package tiles

import "github.com/hajimehoshi/ebiten/v2"

// Tile interface defines common behaviors for all tile types
type Tile interface {
	Tick()
	Render(screen *ebiten.Image, normalBuffer *ebiten.Image)
}

// Constants representing different types of grass tiles
const (
	GRASS_1 = 0
	GRASS_2 = 1
	GRASS_3 = 2

	// Grass path variations
	GP_TL = 3  // Grass path top left
	GP_TM = 4  // Grass path top center
	GP_TR = 5  // Grass path top right
	GP_ML = 6  // Grass path middle left
	GP_M  = 7  // Grass path middle
	GP_MR = 8  // Grass path middle right
	GP_BL = 9  // Grass path bottom left
	GP_BM = 10 // Grass path bottom center
	GP_BR = 11 // Grass path bottom right
)

// NewTile creates a new tile based on the tile type
func NewTile(tilemapImage *Tilemap, normalTilemap *Tilemap, tileType, x, y int) Tile {
	switch tileType {
	case GRASS_1:
		return &GrassTile{
			Image:  getTileImage(tilemapImage, 0, 0),
			Normal: getTileImage(normalTilemap, 0, 0),
			X:      x,
			Y:      y,
		}
	case GRASS_2:
		return &GrassTile{
			Image:  getTileImage(tilemapImage, 1, 0),
			Normal: getTileImage(normalTilemap, 1, 0),
			X:      x,
			Y:      y,
		}
	case GRASS_3:
		return &GrassTile{
			Image:  getTileImage(tilemapImage, 2, 0),
			Normal: getTileImage(normalTilemap, 2, 0),
			X:      x,
			Y:      y,
		}
	case GP_TL:
		return &GrassTile{
			Image:  getTileImage(tilemapImage, 0, 1),
			Normal: getTileImage(normalTilemap, 0, 1),
			X:      x,
			Y:      y,
		}
	case GP_TM:
		return &GrassTile{
			Image:  getTileImage(tilemapImage, 1, 1),
			Normal: getTileImage(normalTilemap, 1, 1),
			X:      x,
			Y:      y,
		}
	case GP_TR:
		return &GrassTile{
			Image:  getTileImage(tilemapImage, 2, 1),
			Normal: getTileImage(normalTilemap, 2, 1),
			X:      x,
			Y:      y,
		}
	case GP_ML:
		return &GrassTile{
			Image:  getTileImage(tilemapImage, 0, 2),
			Normal: getTileImage(normalTilemap, 0, 2),
			X:      x,
			Y:      y,
		}
	case GP_M:
		return &GrassTile{
			Image:  getTileImage(tilemapImage, 1, 2),
			Normal: getTileImage(normalTilemap, 1, 2),
			X:      x,
			Y:      y,
		}
	case GP_MR:
		return &GrassTile{
			Image:  getTileImage(tilemapImage, 2, 2),
			Normal: getTileImage(normalTilemap, 2, 2),
			X:      x,
			Y:      y,
		}
	case GP_BL:
		return &GrassTile{
			Image:  getTileImage(tilemapImage, 0, 3),
			Normal: getTileImage(normalTilemap, 0, 3),
			X:      x,
			Y:      y,
		}
	case GP_BM:
		return &GrassTile{
			Image:  getTileImage(tilemapImage, 1, 3),
			Normal: getTileImage(normalTilemap, 1, 3),
			X:      x,
			Y:      y,
		}
	case GP_BR:
		return &GrassTile{
			Image:  getTileImage(tilemapImage, 2, 3),
			Normal: getTileImage(normalTilemap, 2, 3),
			X:      x,
			Y:      y,
		}
	case 12:
		return &GrassTile{
			Image:  getTileImage(tilemapImage, 3, 0),
			Normal: getTileImage(normalTilemap, 3, 0),
			X:      x,
			Y:      y,
		}
	default:
		return nil // Return nil for unrecognized tile types
	}
}

// getTileImage retrieves a tile image from the tilemap based on x and y coordinates
func getTileImage(tilemapImage *Tilemap, x, y int) *ebiten.Image {
	return tilemapImage.GetTileImage(x, y, 16, 16)
}
