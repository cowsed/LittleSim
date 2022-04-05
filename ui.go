package main

import (
	"fmt"

	"github.com/inkyblackness/imgui-go/v4"
)

var currentDrawTileID int32 = 0

func TileViewWindow(tileID int, size int) {
	ix, iy := CalcTileMapPosition(int(tileID))
	x, y := float32(ix)/float32(TilesImageWidth), float32(iy)/float32(TilesImageHeight)
	w, h := float32(tileSize)/float32(TilesImageWidth), float32(tileSize)/float32(TilesImageHeight)

	imgui.ImageV(10, imgui.Vec2{X: float32(size), Y: float32(size)}, imgui.Vec2{X: x, Y: y}, imgui.Vec2{X: x + w, Y: y + h}, imgui.Vec4{X: 1, Y: 1, Z: 1, W: 1}, imgui.Vec4{X: 1, Y: 1, Z: 1, W: 1})

}

func DrawTileExplorer() {
	//imgui compatible viewer for the tiles
	imgui.SetNextItemOpen(true, imgui.ConditionAppearing)
	if imgui.CollapsingHeader("Tile Explorer") {

		if SelectedX == -1 {
			imgui.Text("No Tile Selected")
		} else {
			imgui.Text(fmt.Sprintf("Chunk: %dx%d -  Tile %dx%d Selected", SelectedChunkX, SelectedChunkY, SelectedX, SelectedY))
			imgui.Text(fmt.Sprint(ReverseTileNames[int(currentTileID)]))
		}
		imgui.InputIntV("Tile Index", &currentTileID, 1, 1, 0)
		TileViewWindow(int(currentTileID), 128)

		if imgui.BeginTableV("IDs", 2, imgui.TableFlagsBorders, imgui.ContentRegionAvail(), 200) {
			for _, v := range JustNames {
				clicked := false

				imgui.TableNextColumn()

				clicked = imgui.Selectable(v)

				imgui.TableNextColumn()

				clicked = imgui.Selectable(fmt.Sprint(TileNames[v])) || clicked
				imgui.TableNextRow()
				if clicked {
					currentTileID = int32(TileNames[v])
				}

			}
			imgui.EndTable()
		}
	}
}

func DrawTilePallette() {
	//imgui compatible viewer for the tiles
	imgui.SetNextItemOpen(true, imgui.ConditionAppearing)
	if imgui.CollapsingHeader("Tile Explorer") {

		imgui.InputIntV("Tile Index", &currentDrawTileID, 1, 1, 0)
		TileViewWindow(int(currentDrawTileID), 128)

		if imgui.BeginTableV("IDs", 2, imgui.TableFlagsBorders, imgui.ContentRegionAvail(), 200) {
			for _, v := range WorldlyTiles {
				clicked := false
				imgui.TableNextColumn()

				clicked = imgui.Selectable(v)

				imgui.TableNextColumn()

				clicked = imgui.Selectable(fmt.Sprint(TileNames[v])) || clicked
				imgui.TableNextRow()
				if clicked {
					fmt.Println("selected")
					currentDrawTileID = int32(MapToDefualt[v])
				}
			}
			imgui.EndTable()
		}
	}
}
