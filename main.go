package main

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"

	simgraphics "LittleSim/Graphics"
	simworld "LittleSim/World"

	"github.com/inkyblackness/imgui-go/v4"
)

var gameWorld *simworld.World
var worldDrawer *simgraphics.RenderSystem

//rs       *RenderSystem

func main() {
	fmt.Println("Beginning World Gen")
	gameWorld = simworld.GenWorld(2)
	fmt.Println("Finished World Gen")

	ebiten.SetWindowSize(800, 600)
	gg := &G{}
	fmt.Println("Beginning Render System")
	worldDrawer = simgraphics.NewRenderSystem(gameWorld)
	fmt.Println("Finished Render System")

	ebiten.RunGame(gg)
}

type G struct {
	clearColor [3]float32
}

func (g *G) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{uint8(g.clearColor[0] * 255), uint8(g.clearColor[1] * 255), uint8(g.clearColor[2] * 255), 255})

	worldDrawer.Draw(screen)
	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %.2f", ebiten.CurrentTPS()))

}

func (g *G) Update() error {
	worldDrawer.Update()

	//inpututil.IsMouseButtonJustPressed
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		if !imgui.CurrentIO().WantCaptureMouse() { //Dont get the mouse if imgui wants it
			x, y := ebiten.CursorPosition()
			worldDrawer.HandleClickAt(x, y)
		}
	}

	return nil
}

func (g *G) Layout(outsideWidth, outsideHeight int) (int, int) {
	worldDrawer.Mgr.SetDisplaySize(float32(800), float32(600))
	return 800, 600
}
