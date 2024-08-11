package game

import (
	_ "embed"
	"fmt"
	"log"
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

//go:embed lighting.kage
var lighting_shader []byte

const (
	screenWidth  = 240
	screenHeight = 180
	tileSize     = 16
)

type Game struct {
	layers  [][]int
	objects [][]int

	time     float64
	lights   []entities.LightEmitter
	shader   *ebiten.Shader
	vertices [4]ebiten.Vertex

	gameTiles   [][]tiles.Tile
	gameObjects [][]objects.Object

	tilemap        *tiles.Tilemap
	normalmap      *tiles.Tilemap
	renderables    []entities.Renderable
	objectEntities []objects.Object
	player         *player.Player
	keys           []ebiten.Key
	camera         *Camera

	renderNormals bool
}

func (g *Game) Update() error {
	g.time = 5.01
	g.keys = inpututil.AppendPressedKeys(g.keys[:0])
	g.player.Tick(g.keys, g.renderables)
	g.camera.Tick(g.player)
	for _, x := range g.objectEntities {
		x.Tick()
	}

	if inpututil.IsKeyJustReleased(ebiten.KeyArrowRight) {
		g.renderNormals = true
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
			tile := tiles.NewTile(g.tilemap, g.normalmap, tileType, x*tileSize, y*tileSize)
			g.gameTiles[y][x] = tile
		}
	}

	for y := 0; y < len(g.objects); y++ {
		g.gameObjects[y] = make([]objects.Object, len(g.objects[0]))
		for x := 0; x < len(g.objects[0]); x++ {
			objectType := g.objects[y][x]
			object := objects.NewObject(g.tilemap, g.normalmap, objectType, x*tileSize, y*tileSize)
			g.gameObjects[y][x] = object
		}
	}
}

func (g *Game) PreRender(startX, startY, endX, endY int) (*ebiten.Image, *ebiten.Image) {
	buffer := ebiten.NewImage(screenWidth, screenHeight)
	normalBuffer := ebiten.NewImage(screenWidth, screenHeight)

	// Render tiles
	for y := startY; y < endY; y++ {
		for x := startX; x < endX; x++ {
			t := g.layers[y][x]

			screenX := (x * tileSize) - g.camera.X
			screenY := (y * tileSize) - g.camera.Y

			tileAt := tiles.NewTile(g.tilemap, g.normalmap, t, screenX, screenY)
			if tileAt != nil {
				tileAt.Render(buffer, normalBuffer)
			}
		}
	}

	// Prepare entities
	var renderables []entities.Renderable
	var lights []entities.LightEmitter

	for y := startY; y < endY; y++ {
		for x := startX; x < endX; x++ {
			t := g.gameObjects[y][x]
			if t != nil {
				t.SetPosition((x*tileSize)-g.camera.X, ((y*tileSize)-g.camera.Y)-16)
				renderables = append(renderables, t)
				g.objectEntities = append(g.objectEntities, t)
				isLight, lightParameters := t.GetLightParameters()
				if isLight {
					lights = append(lights, lightParameters)
				}
			}
		}
	}

	// Render player
	renderables = append(renderables, g.player)

	// Render entities
	// Sort renderables by Y position, taking into account the mid-point logic
	sort.SliceStable(renderables, func(i, j int) bool {
		iY := renderables[i].GetY()
		jY := renderables[j].GetY()

		if renderables[i] == g.player {
			return iY-g.player.CameraY < jY
		} else if renderables[j] == g.player {
			return jY-g.player.CameraY < iY
		}
		// Sort so that entities with lower base Y are rendered first
		return iY < jY
	})

	g.renderables = renderables
	g.lights = lights
	// Render sorted objects and player
	for _, r := range renderables {
		r.Render(buffer, normalBuffer)
	}
	return buffer, normalBuffer
}

func (g *Game) RenderCamera() (int, int, int, int) {
	cameraBuffer := 8 // Number of extra tiles to render around the screen edges

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
	return startX, startY, endX, endY
}

func (g *Game) Draw(screen *ebiten.Image) {

	// Setup shaders
	// Dest (screen)
	bounds := screen.Bounds()
	g.vertices[0].DstX = float32(bounds.Min.X) // top-left
	g.vertices[0].DstY = float32(bounds.Min.Y) // top-left
	g.vertices[1].DstX = float32(bounds.Max.X) // top-right
	g.vertices[1].DstY = float32(bounds.Min.Y) // top-right
	g.vertices[2].DstX = float32(bounds.Min.X) // bottom-left
	g.vertices[2].DstY = float32(bounds.Max.Y) // bottom-left
	g.vertices[3].DstX = float32(bounds.Max.X) // bottom-right
	g.vertices[3].DstY = float32(bounds.Max.Y) // bottom-right

	// Source (rendered screen)
	srcBounds := screen.Bounds()
	g.vertices[0].SrcX = float32(srcBounds.Min.X) // top-left
	g.vertices[0].SrcY = float32(srcBounds.Min.Y) // top-left
	g.vertices[1].SrcX = float32(srcBounds.Max.X) // top-right
	g.vertices[1].SrcY = float32(srcBounds.Min.Y) // top-right
	g.vertices[2].SrcX = float32(srcBounds.Min.X) // bottom-left
	g.vertices[2].SrcY = float32(srcBounds.Max.Y) // bottom-left
	g.vertices[3].SrcX = float32(srcBounds.Max.X) // bottom-right
	g.vertices[3].SrcY = float32(srcBounds.Max.Y) // bottom-right

	startX, startY, endX, endY := g.RenderCamera()
	buffer, normalBuffer := g.PreRender(startX, startY, endX, endY)

	op := &ebiten.DrawTrianglesShaderOptions{}
	op.Images[0] = buffer
	op.Images[1] = normalBuffer
	op.Uniforms = make(map[string]interface{})

	// Pack light data into uniform arrays
	var lightPositions = make([]float32, 600) // Max lights of 100

	for i, light := range g.lights {
		base := i * 6
		lightPositions[base] = float32(light.X)
		lightPositions[base+1] = float32(light.Y)
		lightPositions[base+2] = light.Intensity
		lightPositions[base+3] = light.Color[0]
		lightPositions[base+4] = light.Color[1]
		lightPositions[base+5] = light.Color[2]
	}

	cx, cy := ebiten.CursorPosition()
	op.Uniforms["Cursor"] = []float32{float32(cx), float32(cy)}
	op.Uniforms["Time"] = float32(g.time)
	op.Uniforms["LightCount"] = len(lightPositions)
	op.Uniforms["LightPositions"] = lightPositions

	indices := []uint16{0, 1, 2, 2, 1, 3} // map vertices to triangles
	screen.DrawTrianglesShader(g.vertices[:], indices, g.shader, op)
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

func Start() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Gam")

	homeLevel := levels.HomeLevel()

	// Clamp the camera to the edges of the game world
	maxCameraX := len(homeLevel.Heightmap[0])*tileSize - screenWidth
	maxCameraY := len(homeLevel.Heightmap)*tileSize - screenHeight

	// Load shader
	shader, shaderErr := ebiten.NewShader(lighting_shader)
	if shaderErr != nil {
		log.Fatalf("Failed to load shader %s", shaderErr.Error())
	}

	game := &Game{
		camera:    NewCamera(maxCameraX, maxCameraY),
		player:    player.NewPlayer(tiles.NewTilemap("assets/Tilemap/tilemap2_packed.png"), tiles.NewTilemap("assets/Tilemap/tilemap2_packed_n.png"), 0, 0),
		tilemap:   tiles.NewTilemap("assets/Tilemap/tilemap_packed.png"),
		normalmap: tiles.NewTilemap("assets/Tilemap/tilemap_packed_n.png"),
		layers:    homeLevel.Heightmap,
		objects:   homeLevel.Objectmap,
	}

	game.lights = []entities.LightEmitter{}

	game.time = 0
	game.LoadLevel()
	game.shader = shader
	game.renderNormals = false
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
