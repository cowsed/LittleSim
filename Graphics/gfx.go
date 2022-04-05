package simgraphics

import (
	communications "LittleSim/Communications"
	simworld "LittleSim/World"
	simdata "LittleSim/WorldData"
	"fmt"

	"image"
	_ "image/png"

	"github.com/gabstv/ebiten-imgui/renderer"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/inkyblackness/imgui-go/v4"
)

var NumTilesX int
var NumTilesY int
var TilesImageWidth int
var TilesImageHeight int

const ChunkSize = 16
const tileSize = 16
const tileBorder = 1

var TileMap *ebiten.Image

//Coordinates and ID of selected tike
var SelectedX = -1
var SelectedY = -1
var SelectedChunkX = -1
var SelectedChunkY = -1
var currentTileID int32 = 0

//0 View Drawing information
//1 Placing tiles
//2 View tile info
//3 View Entity info
var LeftClickMode int = 0

type Camera struct {
	focusPointX float32
	focusPointY float32
	zoom        float32
}

type RenderSystem struct {
	myWorld        *simworld.World
	camera         Camera
	doAutotile     bool
	chunkRenderers map[[2]int]*ChunkRenderer
	Mgr            *renderer.Manager
	clearColor     [3]float32
}

func NewRenderSystem(world *simworld.World) *RenderSystem {

	LoadTiles()

	r := RenderSystem{
		myWorld:    world,
		camera:     Camera{0, 0, 1},
		doAutotile: true,
		Mgr:        renderer.New(nil),
	}
	r.Mgr.Cache.SetTexture(10, TileMap)
	r.chunkRenderers = r.MakeChunkRenderers()

	return &r
}
func (rs *RenderSystem) Draw(screen *ebiten.Image) {
	//Check for updates
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

	//Draw game world
	for i := range rs.chunkRenderers {
		rs.chunkRenderers[i].Draw(screen, rs.myWorld, rs.doAutotile)
	}

	//Draw made UI
	rs.Mgr.Draw(screen)

}

func (rs *RenderSystem) Update() {
	rs.Mgr.Update(1.0 / 60.0)
	rs.Mgr.BeginFrame()
	//Draw UI
	imgui.ColorEdit3("clear color", &rs.clearColor) // Edit 3 floats representing a color

	if imgui.Checkbox("Do Autotile", &rs.doAutotile) {
		rs.RedrawAll()
	}

	imgui.RadioButtonInt("Draw info", &LeftClickMode, 0)
	imgui.RadioButtonInt("Draw Tiles", &LeftClickMode, 1)
	imgui.RadioButtonInt("Tile Info", &LeftClickMode, 2)

	if LeftClickMode == 0 {
		imgui.Begin("TileExplorer")
		DrawTileExplorer()
		imgui.End()
	} else if LeftClickMode == 1 {
		imgui.Begin("Tile Pallette")
		DrawTilePallette()
		imgui.End()
	}
	rs.Mgr.EndFrame()

}

func (rs *RenderSystem) HandleClickAt(mouseX, mouseY int) {
	SelectedX = (mouseX / tileSize)
	SelectedY = (mouseY / tileSize)
	SelectedChunkX = SelectedX / 16
	SelectedChunkY = SelectedY / 16

	SelectedX = SelectedX % 16
	SelectedY = SelectedY % 16
	cc := simdata.ChunkCoord{SelectedChunkX, SelectedChunkY}

	if LeftClickMode == 0 {
		if !rs.myWorld.ChunkExists(cc) {
			currentTileID = simdata.DEFUALT_TILE
			return
		}
		currentTileID = int32(rs.chunkRenderers[cc].myTiledChunk[SelectedY][SelectedX])
	} else if LeftClickMode == 1 {

		rs.myWorld.SetTile(SelectedX, SelectedY, cc, int(currentDrawTileID))
		Autotile(SelectedX, SelectedY, cc, rs.myWorld, true)
	}
}
func (rs *RenderSystem) MakeChunkRenderers() map[[2]int]*ChunkRenderer {
	crs := map[[2]int]*ChunkRenderer{}
	for yChunk := 0; yChunk < 2; yChunk++ {
		for xChunk := 0; xChunk < 2; xChunk++ {
			myCoords := [2]int{xChunk, yChunk}
			crs[myCoords] = &ChunkRenderer{
				myChunk:       (rs.myWorld.WorldMap[myCoords]),
				myChunkCoords: myCoords,
				drawnChunk:    ebiten.NewImage(ChunkSize*tileSize, ChunkSize*tileSize),
				upToDate:      false,
			}
		}
	}
	return crs
}

func (rs *RenderSystem) RedrawAll() {
	for i := range rs.chunkRenderers {
		rs.chunkRenderers[i].upToDate = false
	}
}
func (rs *RenderSystem) InvalidateChunkAt(x, y int, ID [2]int) {
	if !rs.myWorld.ChunkExists(ID) {
		return
	}
	rs.chunkRenderers[ID].upToDate = false
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

type ChunkRenderer struct {
	myChunk       *simdata.Chunk
	myTiledChunk  [16][16]int
	myChunkCoords simdata.ChunkCoord
	drawnChunk    *ebiten.Image
	upToDate      bool
}

func DrawTile(dest, tilemap *ebiten.Image, x, y int, tileID int) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(x*tileSize), float64(y*tileSize))
	sx, sy := CalcTileMapPosition(tileID)

	dest.DrawImage(tilemap.SubImage(image.Rect(sx, sy, sx+tileSize, sy+tileSize)).(*ebiten.Image), op)
}

func (c *ChunkRenderer) DrawChunkToImage(chunkImage *ebiten.Image, w *simworld.World, auto bool) {
	BaseMaterial := simdata.TileNames["green_ground"]
	for y := range c.myChunk.ChunkData {
		for x := range c.myChunk.ChunkData[y] {
			DrawTile(chunkImage, TileMap, x, y, BaseMaterial)

			m := c.myChunk.ChunkData[y][x].Material
			if auto {
				m = Autotile(x, y, c.myChunkCoords, w, false)
			}
			c.myTiledChunk[y][x] = m
			//m = DEFUALT_TILE
			DrawTile(chunkImage, TileMap, x, y, m)
		}
	}
}

func (c *ChunkRenderer) Draw(screen *ebiten.Image, w *simworld.World, auto bool) {
	if !c.upToDate {
		c.drawnChunk.Clear()
		c.DrawChunkToImage(c.drawnChunk, w, auto)
		c.upToDate = true
	}

	op := &ebiten.DrawImageOptions{}
	posX := c.myChunkCoords[0] * tileSize * ChunkSize
	posY := c.myChunkCoords[1] * tileSize * ChunkSize

	op.GeoM.Translate(float64(posX), float64(posY))
	screen.DrawImage(c.drawnChunk, op)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
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

func GetTileSafe(x, y int, ID simdata.ChunkCoord, world *simworld.World) int {
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
