package game

import (
	"fmt"
	"image/color"
	"log"
	"math"
	"sort"

	"github.com/alexmgriffiths/gam/entities"
	"github.com/alexmgriffiths/gam/entities/objects"
	"github.com/alexmgriffiths/gam/entities/player"
	"github.com/alexmgriffiths/gam/entities/tiles"
	"github.com/alexmgriffiths/gam/levels"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	screenWidth  = 240
	screenHeight = 180
	tileSize     = 16
)

var (
	bgImage        *ebiten.Image
	fgImage        *ebiten.Image
	maskedFgImage  = ebiten.NewImage(240, 180)
	spotLightImage *ebiten.Image
)

type Game struct {
	layers  [][]int
	objects [][]int

	gameTiles   [][]tiles.Tile
	gameObjects [][]objects.Object

	lightMask *ebiten.Image

	tilemap        *tiles.Tilemap
	renderables    []entities.Renderable
	objectEntities []objects.Object
	player         *player.Player
	keys           []ebiten.Key
	camera         *Camera
}

func (g *Game) Update() error {
	g.keys = inpututil.AppendPressedKeys(g.keys[:0])
	g.player.Tick(g.keys, g.renderables)
	g.camera.Tick(g.player)
	for _, x := range g.objectEntities {
		x.Tick()
	}
	return nil
}

func (g *Game) LoadLevel() {
	// Initialize the gameTiles and gameObjects with the same size as layers and objects
	g.gameTiles = make([][]tiles.Tile, len(g.layers))
	g.gameObjects = make([][]objects.Object, len(g.objects))
	for y := 0; y < len(g.layers); y++ {
		g.gameTiles[y] = make([]tiles.Tile, len(g.layers[0]))
		for x := 0; x < len(g.layers[0]); x++ {
			tileType := g.layers[y][x]
			tile := tiles.NewTile(g.tilemap, tileType, x*tileSize, y*tileSize)
			g.gameTiles[y][x] = tile
		}
	}

	for y := 0; y < len(g.objects); y++ {
		g.gameObjects[y] = make([]objects.Object, len(g.objects[0]))
		for x := 0; x < len(g.objects[0]); x++ {
			objectType := g.objects[y][x]
			object := objects.NewObject(g.tilemap, objectType, x*tileSize, y*tileSize)
			g.gameObjects[y][x] = object
		}
	}
}

func (g *Game) Draw(screen *ebiten.Image) {

	cameraBuffer := 2 // Number of extra tiles to render around the screen edges

	startX := (g.camera.X / tileSize) - cameraBuffer
	startY := (g.camera.Y / tileSize) - cameraBuffer

	endX := (g.camera.X+screenWidth)/tileSize + cameraBuffer
	endY := (g.camera.Y+screenHeight)/tileSize + cameraBuffer

	// Ensure start indices do not go below 0
	if startX < 0 {
		startX = 0
	}
	if startY < 0 {
		startY = 0
	}

	// Ensure end indices do not exceed the map dimensions
	if endX >= len(g.layers[0]) {
		endX = len(g.layers[0])
	}
	if endY >= len(g.layers) {
		endY = len(g.layers)
	}

	for y := startY; y < endY; y++ {
		for x := startX; x < endX; x++ {
			t := g.layers[y][x]

			screenX := (x * tileSize) - g.camera.X
			screenY := (y * tileSize) - g.camera.Y

			tileAt := tiles.NewTile(g.tilemap, t, screenX, screenY)
			if tileAt != nil {
				tileAt.Render(screen)
			}
		}
	}

	var renderables []entities.Renderable

	for y := startY; y < endY; y++ {
		for x := startX; x < endX; x++ {
			t := g.gameObjects[y][x]
			if t != nil {
				t.SetPosition((x*tileSize)-g.camera.X, (y*tileSize)-g.camera.Y)
				renderables = append(renderables, t)
				g.objectEntities = append(g.objectEntities, t)
			}
		}
	}
	renderables = append(renderables, g.player)
	// Sort renderables by Y position, taking into account the mid-point logic
	sort.SliceStable(renderables, func(i, j int) bool {
		iY := renderables[i].GetY()
		jY := renderables[j].GetY()

		// Sort so that entities with lower base Y are rendered first
		return iY < jY
	})
	g.renderables = renderables
	// Render sorted objects and player
	for _, r := range renderables {
		r.Render(screen)
	}

	// Debug
	// for y := startY; y < endY; y++ {
	// 	for x := startX; x < endX; x++ {
	// 		screenX := (x * tileSize) - g.camera.X
	// 		screenY := (y * tileSize) - g.camera.Y

	// 		vector.StrokeRect(screen, float32(screenX), float32(screenY), 16, 16, 1, color.Black, true)
	// 	}
	// }

	// Display debug information
	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f", ebiten.ActualTPS()))
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%d,%d", g.player.X, g.player.Y), 0, 20)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func createLightMask(radius int) *ebiten.Image {
	diameter := radius * 2
	mask := ebiten.NewImage(diameter, diameter)
	mask.Fill(color.RGBA{100, 100, 100, 100}) // Start with a black (fully transparent) background

	for y := 0; y < diameter; y++ {
		for x := 0; x < diameter; x++ {
			dx := float64(x - radius)
			dy := float64(y - radius)
			distance := math.Sqrt(dx*dx + dy*dy)
			if distance < float64(radius) {
				alpha := uint8(255 * (1 - distance/float64(radius)))
				mask.Set(x, y, color.RGBA{255, 255, 255, alpha})
			}
		}
	}

	return mask
}

func Start() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Gam")

	homeLevel := levels.HomeLevel()

	// Clamp the camera to the edges of the game world
	maxCameraX := len(homeLevel.Heightmap[0])*tileSize - screenWidth
	maxCameraY := len(homeLevel.Heightmap)*tileSize - screenHeight

	game := &Game{
		camera:  NewCamera(maxCameraX, maxCameraY),
		player:  player.NewPlayer(tiles.NewTilemap("assets/Tilemap/tilemap2_packed.png"), 0, 0),
		tilemap: tiles.NewTilemap("assets/Tilemap/tilemap_packed.png"),
		layers:  homeLevel.Heightmap,
		objects: homeLevel.Objectmap,
	}

	lightMask := createLightMask(100)
	game.lightMask = lightMask
	game.LoadLevel()

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
