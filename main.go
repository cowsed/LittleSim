package main

import (
	"fmt"
	"image/color"

	"github.com/gabstv/ebiten-imgui/renderer"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"

	"github.com/inkyblackness/imgui-go/v4"
)

var gameWorld *World
var worldDrawer *RenderSystem

func main() {
	fmt.Println("Beginning World Gen")
	gameWorld = GenWorld(2)
	fmt.Println("Finished World Gen")

	mgr := renderer.New(nil)

	ebiten.SetWindowSize(800, 600)
	gg := &G{
		mgr: mgr,
	}
	fmt.Println("Beginning Render System")
	worldDrawer = NewRenderSystem(gameWorld, gg)
	fmt.Println("Finished Render System")
	gameWorld.rs = worldDrawer

	//Autotile(4, 4, [2]int{0, 0}, gameWorld)

	ebiten.RunGame(gg)
}

type G struct {
	mgr *renderer.Manager

	clearColor [3]float32
}

func (g *G) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{uint8(g.clearColor[0] * 255), uint8(g.clearColor[1] * 255), uint8(g.clearColor[2] * 255), 255})

	worldDrawer.Draw(screen)
	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %.2f", ebiten.CurrentTPS()))
	g.mgr.Draw(screen)
}

func (g *G) Update() error {
	g.mgr.Update(1.0 / 60.0)
	g.mgr.BeginFrame()

	//inpututil.IsMouseButtonJustPressed
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		if !imgui.CurrentIO().WantCaptureMouse() { //Dont get the mouse if imgui wants it
			x, y := ebiten.CursorPosition()
			worldDrawer.HandleClickAt(x, y)
		}
	}

	//Windows
	{
		imgui.ColorEdit3("clear color", &g.clearColor) // Edit 3 floats representing a color

		if imgui.Checkbox("Autotile", &worldDrawer.doAutotile) {
			worldDrawer.RedrawAll()
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

	}
	g.mgr.EndFrame()
	return nil
}

func (g *G) Layout(outsideWidth, outsideHeight int) (int, int) {
	g.mgr.SetDisplaySize(float32(800), float32(600))
	return 800, 600
}
