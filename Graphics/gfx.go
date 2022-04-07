package simgraphics

import (
	communications "LittleSim/Communications"
	simworld "LittleSim/World"
	simdata "LittleSim/WorldData"
	"fmt"

	"image/color"
	_ "image/png"

	"github.com/gabstv/ebiten-imgui/renderer"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

//Tile Map parameters
var TileMap *ebiten.Image

var NumTilesX int
var NumTilesY int
var TilesImageWidth int
var TilesImageHeight int

//Tile Drawing parameters
const ChunkSize = simdata.ChunkSize
const tileSize = 16  //px
const tileBorder = 1 //px

//Coordinates and ID of selected tike
var SelectedX = -1
var SelectedY = -1
var SelectedChunkX = -1
var SelectedChunkY = -1

var currentTilemapID simdata.TilemapKey
var currentDrawTileID simdata.TilemapKey

//0 View Drawing information - gfx specific
//1 Placing tiles
//2 View tile info
//3 View Entity info
const (
	TileGraphicsInfo int = iota
	TileDrawing
	TileInfo
	EntityInfo
)

//Controls what left clicking will do
var CurrentLeftClickMode int = TileDrawing

type Camera struct {
	focusPointX    float32
	focusPointY    float32
	zoom           float32
	viewportWidth  float32
	viewportHeight float32
}

type RenderSystem struct {
	myWorld        *simworld.World
	camera         *Camera
	animationTick  int
	doAutotile     bool
	chunkRenderers map[[2]int]*ChunkRenderer
	GUImgr         *renderer.Manager
	clearColor     [3]float32
}

func NewRenderSystem(world *simworld.World) *RenderSystem {
	//Load tiles from file
	LoadTiles()

	r := RenderSystem{
		myWorld: world,
		camera: &Camera{
			focusPointX: 0,
			focusPointY: 0,
			zoom:        1,
		},
		doAutotile: true,
		GUImgr:     renderer.New(nil),
	}
	r.GUImgr.Cache.SetTexture(10, TileMap)
	r.chunkRenderers = r.MakeChunkRenderers()

	return &r
}

//SetSize updates all render systems that need to know the size of the window
func (rs *RenderSystem) SetSize(w, h int) {
	rs.GUImgr.SetDisplaySize(float32(w), float32(h))
	rs.camera.viewportHeight = float32(h)
	rs.camera.viewportWidth = float32(w)
}
func (rs *RenderSystem) DrawWorld(screen *ebiten.Image) {
	rs.animationTick++

	//Clear Screen
	screen.Fill(color.RGBA{uint8(rs.clearColor[0] * 255), uint8(rs.clearColor[1] * 255), uint8(rs.clearColor[2] * 255), 255})

	//Draw game world
	for i := range rs.chunkRenderers {
		rs.chunkRenderers[i].Draw(screen, rs.myWorld, rs.animationTick, rs.doAutotile, rs.camera)
	}
}

//Actually draw the UI to the screen
func (rs *RenderSystem) DrawUI(screen *ebiten.Image) {
	rs.GUImgr.Draw(screen)
}
func (rs *RenderSystem) Update() {
	rs.UpdateTileMaps()

	//Begin UI
	rs.GUImgr.Update(1.0 / 60.0)
	rs.GUImgr.BeginFrame()
	//Create  UI
	DrawDebugUI(rs)

	//End UI
	rs.GUImgr.EndFrame()
}
func (rs *RenderSystem) UpdateTileMaps() {
	//Check for updates to the tilemaps
	more_updates := true
	for more_updates {
		select {
		case x, ok := <-communications.InvalidateChunkChannel:
			if ok {
				rs.InvalidateChunkAt(x.X, x.Y, x.ID)
			} else {
				fmt.Println("Channel Closed")
			}
		default:
			more_updates = false
		}
	}
}

func (rs *RenderSystem) MakeChunkRenderers() map[[2]int]*ChunkRenderer {
	crs := map[[2]int]*ChunkRenderer{}
	for yChunk := 0; yChunk < 2; yChunk++ {
		for xChunk := 0; xChunk < 2; xChunk++ {
			myCoords := [2]int{xChunk, yChunk}
			crs[myCoords] = &ChunkRenderer{
				myChunk:       rs.myWorld.WorldMap[myCoords],
				myChunkCoords: myCoords,
				drawnChunk:    ebiten.NewImage(ChunkSize*tileSize, ChunkSize*tileSize),
				entityImage:   ebiten.NewImage(ChunkSize*tileSize, ChunkSize*tileSize),

				upToDate: false,
			}
		}
	}
	return crs
}

func (rs *RenderSystem) InvalidateAllChunks() {
	for i := range rs.chunkRenderers {
		rs.chunkRenderers[i].upToDate = false
	}
}
func (rs *RenderSystem) InvalidateChunkAt(x, y int, ID [2]int) {
	if !rs.myWorld.ChunkExists(ID) {
		return
	}
	rs.chunkRenderers[ID].upToDate = false
	//If on the border, invalidate neighbouring chunks. Autotiling across chunk borders
	if x == 0 {
		rs.InvalidateChunk([2]int{ID[0] - 1, ID[1]})
	} else if x == 15 {
		rs.InvalidateChunk([2]int{ID[0] + 1, ID[1]})
	}
	if y == 0 {
		rs.InvalidateChunk([2]int{ID[0], ID[1] - 1})
	} else if y == 15 {
		rs.InvalidateChunk([2]int{ID[0], ID[1] + 1})
	}

}
func (rs *RenderSystem) InvalidateChunk(ID [2]int) {
	if !rs.myWorld.ChunkExists(ID) {
		return
	}
	rs.chunkRenderers[ID].upToDate = false
}

func CalcTileMapPosition(tileID int) (int, int) {
	sx := (tileID % NumTilesX) * (tileSize + tileBorder)
	sy := (tileID / NumTilesX) * (tileSize + tileBorder)
	return sx, sy
}

func LoadTiles() {
	eimg, _, err := ebitenutil.NewImageFromFile("Assets/tiles.png")
	check(err)
	w, h := eimg.Bounds().Dx(), eimg.Bounds().Dy()
	TilesImageWidth = w
	TilesImageHeight = h

	NumTilesX = w / (tileSize + tileBorder)
	NumTilesY = h / (tileSize + tileBorder)

	TileMap = eimg
}

func GetTileSafe(x, y int, ID simdata.ChunkCoord, world *simworld.World) simdata.TilemapKey {
	if (x > 16 || x < (-16)) || (y > 16 || y < (-16)) {
		panic("Out of bounds really hard")
	}
	if x < 0 {
		x = 16 - x
		ID[0] -= 1
	}
	if x > 15 {
		x = x - 16
		ID[0] += 1
	}
	if y < 0 {
		y = 16 - y
		ID[1] -= 1
	}
	if y > 15 {
		y = y - 16
		ID[1] += 1
	}

	//Check if chunk exists
	if !world.ChunkExists(ID) {
		return simdata.DEFUALT_TILE
	}

	return world.GetChunk(ID).GetTile(x, y).Material
}
