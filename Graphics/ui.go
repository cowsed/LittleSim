package simgraphics

import (
	simdata "LittleSim/WorldData"
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/inkyblackness/imgui-go/v4"
)

func DrawDebugUI(rs *RenderSystem) {
	imgui.ColorEdit3("clear color", &rs.clearColor) // Edit 3 floats representing a color
	imgui.Text(fmt.Sprintf("FPS : %.2f\nTPS: %.2f", ebiten.CurrentFPS(), ebiten.CurrentFPS()))

	if imgui.Checkbox("Do Autotile", &rs.doAutotile) {
		rs.InvalidateAllChunks()
	}
	imgui.DragFloatV("Zoom", &rs.camera.zoom, 0.01, 0, 10, "%f", 0)
	imgui.DragFloat("X", &rs.camera.focusPointX)
	imgui.DragFloat("Y", &rs.camera.focusPointY)

	imgui.RadioButtonInt("Draw info", &CurrentLeftClickMode, 0)
	imgui.RadioButtonInt("Draw Tiles", &CurrentLeftClickMode, 1)
	imgui.RadioButtonInt("Tile Info", &CurrentLeftClickMode, 2)

	if CurrentLeftClickMode == 0 {
		DrawTileExplorer()
	} else if CurrentLeftClickMode == 1 {
		DrawWorldPalette()
	}
}

//Tile preview makes an imgui image that displays the requested tile of TileID at a specified size
//TODO make it honor the sheet arguement
func TilePreview(tileID simdata.TilemapKey, size int) {
	ix, iy := CalcTileMapPosition(int(tileID.Tile))
	x, y := float32(ix)/float32(TilesImageWidth), float32(iy)/float32(TilesImageHeight)
	w, h := float32(tileSize)/float32(TilesImageWidth), float32(tileSize)/float32(TilesImageHeight)

	imgui.ImageV(10, imgui.Vec2{X: float32(size), Y: float32(size)}, imgui.Vec2{X: x, Y: y}, imgui.Vec2{X: x + w, Y: y + h}, imgui.Vec4{X: 1, Y: 1, Z: 1, W: 1}, imgui.Vec4{X: 1, Y: 1, Z: 1, W: 1})

}

func InputInt(label string, val *int) {
	v := int32(*val)
	imgui.InputInt(label, &v)
	*val = int(v)

}

//Shows all available tiles from the tilemap
func DrawTileExplorer() {
	imgui.Begin("Tile Pallette")

	if SelectedX == -1 {
		imgui.Text("No Tile Selected")
	} else {
		imgui.Text(fmt.Sprintf("Chunk: %dx%d -  Tile %dx%d Selected", SelectedChunkX, SelectedChunkY, SelectedX, SelectedY))
		imgui.Text(fmt.Sprint(simdata.ReverseTileNames[currentTilemapID]))
	}
	//imgui.InputIntV("Tile Index", (*int32)(&currentTilemapID.Tile), 1, 1, 0)
	InputInt("Tile Index", &currentTilemapID.Tile)
	TilePreview(currentTilemapID, 128)

	if imgui.BeginTableV("IDs", 2, imgui.TableFlagsBorders, imgui.ContentRegionAvail(), 200) {
		for _, v := range simdata.JustNames {
			clicked := false

			imgui.TableNextColumn()

			clicked = imgui.Selectable(v)

			imgui.TableNextColumn()

			clicked = imgui.Selectable(fmt.Sprint(simdata.TileNames[v])) || clicked
			imgui.TableNextRow()
			if clicked {
				currentTilemapID = simdata.TileNames[v]
			}

		}
		imgui.EndTable()
	}
	imgui.End()
}

//Shows all the types of materials that can be drawn into the world
func DrawWorldPalette() {
	//imgui compatible viewer for the tiles
	imgui.Begin("World Blocks")

	InputInt("Tile Index", &currentDrawTileID.Tile)
	TilePreview(currentDrawTileID, 128)

	if imgui.BeginTableV("IDs", 2, imgui.TableFlagsBorders, imgui.ContentRegionAvail(), 200) {
		for _, v := range simdata.WorldlyTiles {
			clicked := false
			imgui.TableNextColumn()

			clicked = imgui.Selectable(v)

			imgui.TableNextColumn()

			clicked = imgui.Selectable(fmt.Sprint(simdata.TileNames[v])) || clicked
			imgui.TableNextRow()
			if clicked {
				fmt.Println("selected")
				currentDrawTileID = (simdata.MapToDefualt[v])
			}
		}
		imgui.EndTable()
	}
	imgui.End()
}
