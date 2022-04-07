package main

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"

	simgraphics "LittleSim/Graphics"
	simworld "LittleSim/World"
	simdata "LittleSim/WorldData"

	"github.com/inkyblackness/imgui-go/v4"
)

func main() {

	//Window setup
	ebiten.SetWindowSize(800, 600)
	ebiten.SetWindowResizable(true)
	gg := &G{}

	gg.gameWorld = simworld.GenWorld(2)

	u := uuid.New()
	u2 := uuid.New()

	fire := &simworld.Flame{
		ID:           u,
		Pos:          simdata.Location{X: 0, Y: 0, CC: simdata.ChunkCoord{0, 0}},
		Heat:         1,
		Light:        1,
		Extinguished: false,
		Frames:       simdata.MakeFrameList([]simdata.TilemapKey{simdata.TileNames["campfire_1"], simdata.TileNames["campfire_2"]}, []int{16, 16}),
	}
	fire2 := &simworld.Flame{
		ID:           u2,
		Pos:          simdata.Location{X: 2, Y: 0, CC: simdata.ChunkCoord{0, 0}},
		Heat:         1,
		Light:        1,
		Extinguished: false,
		Frames:       simdata.MakeFrameList([]simdata.TilemapKey{simdata.TileNames["small_fireplace_1"], simdata.TileNames["small_fireplace_2"]}, []int{16, 16}),
	}

	gg.gameWorld.WorldMap[fire.Pos.CC].LocalEntities[u] = fire
	gg.gameWorld.WorldMap[fire2.Pos.CC].LocalEntities[u2] = fire2

	gg.worldDrawer = simgraphics.NewRenderSystem(gg.gameWorld)

	ebiten.RunGame(gg)
}

type G struct {
	gameWorld   *simworld.World
	worldDrawer *simgraphics.RenderSystem
}

//Draws the world and ui to the screen
func (g *G) Draw(screen *ebiten.Image) {

	g.worldDrawer.DrawWorld(screen)
	g.worldDrawer.DrawUI(screen)

	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %.2f", ebiten.CurrentTPS()))

}

//Update runs the game update function.
//Steps the simulation, Updates the ui
//Handles the input
func (g *G) Update() error {
	g.gameWorld.TickSimulation()
	g.worldDrawer.Update()

	//inpututil.IsMouseButtonJustPressed
	if !imgui.CurrentIO().WantCaptureMouse() { //Dont get the mouse if imgui wants it
		g.worldDrawer.HandleMouse()

	}

	return nil
}

//Handles the sizing of the window
func (g *G) Layout(outsideWidth, outsideHeight int) (int, int) {
	g.worldDrawer.SetSize(outsideWidth, outsideHeight)
	return outsideWidth, outsideHeight
}
