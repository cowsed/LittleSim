package simgraphics

import (
	simdata "LittleSim/WorldData"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

var startDragX int
var startDragY int

var initialCamX float32
var initialCamY float32

var RMBHeld bool = false

func (rs *RenderSystem) HandleMouse() {
	x, y := ebiten.CursorPosition()

	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		rs.HandleLeftDragAt(x, y)
	}

	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonRight) {
		startDragX = x
		startDragY = y

		initialCamX = rs.camera.focusPointX
		initialCamY = rs.camera.focusPointY

		RMBHeld = true
	}
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight) {
		RMBHeld = true
		dx := x - startDragX
		dy := y - startDragY
		//fmt.Println(dx, dy)
		rs.camera.focusPointX = initialCamX - float32(dx)/rs.camera.zoom //float32(dx) + rs.camera.focusPointX
		rs.camera.focusPointY = initialCamY - float32(dy)/rs.camera.zoom //float32(dy) + rs.camera.focusPointX

	}
	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonRight) {
		RMBHeld = false
	}
}

func (rs *RenderSystem) HandleLeftDragAt(mouseX, mouseY int) {
	//SelectedX = (mouseX / tileSize)
	//SelectedY = (mouseY / tileSize)
	//SelectedChunkX = SelectedX / 16
	//SelectedChunkY = SelectedY / 16
	//
	//SelectedX = SelectedX % 16
	//SelectedY = SelectedY % 16
	//cc := simdata.ChunkCoord{SelectedChunkX, SelectedChunkY}

	worldLoc := rs.camera.ScreenPosToWorldPos(float64(mouseX), float64(mouseY))
	SelectedX = int(worldLoc.X)
	SelectedY = int(worldLoc.Y)
	SelectedChunkX = worldLoc.CC[0]
	SelectedChunkY = worldLoc.CC[1]

	if CurrentLeftClickMode == 0 {
		if !rs.myWorld.ChunkExists(worldLoc.CC) {
			currentTilemapID = simdata.DEFUALT_TILE
			return
		}
		currentTilemapID = rs.chunkRenderers[worldLoc.CC].myTiledChunk[SelectedY][SelectedX]
		Autotile(SelectedX, SelectedY, worldLoc.CC, rs.myWorld, true)

	} else if CurrentLeftClickMode == 1 {
		if !(SelectedX < 0 || SelectedY < 0) {
			rs.myWorld.SetTile(SelectedX, SelectedY, worldLoc.CC, currentDrawTileID)
		}
	}
}
