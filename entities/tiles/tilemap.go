package tiles

import (
	"fmt"
	"image"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Tilemap struct {
	Image *ebiten.Image
}

var img *ebiten.Image
var err error

func NewTilemap(path string) *Tilemap {
	currentDir, _ := os.Getwd()
	img, _, err = ebitenutil.NewImageFromFile(fmt.Sprintf("%s/%s", currentDir, path))
	if err != nil {
		panic("Failed to load tilemap")
	}
	return &Tilemap{
		Image: img,
	}
}

func (tm *Tilemap) GetTileImage(spriteX, spriteY, tileWidth, tileHeight int) *ebiten.Image {
	sx := spriteX * tileWidth
	sy := spriteY * tileHeight
	rect := image.Rect(sx, sy, sx+tileWidth, sy+tileHeight)
	tileImage := tm.Image.SubImage(rect).(*ebiten.Image)
	return tileImage
}

func (tm *Tilemap) GetTileImageCustom(spriteX, spriteY, tileWidth, tileHeight int) *ebiten.Image {
	sx := spriteX
	sy := spriteY
	rect := image.Rect(sx, sy, sx+tileWidth, sy+tileHeight)
	tileImage := tm.Image.SubImage(rect).(*ebiten.Image)
	return tileImage
}
