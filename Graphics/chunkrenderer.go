package simgraphics

import (
	simworld "LittleSim/World"
	simdata "LittleSim/WorldData"
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

type ChunkRenderer struct {
	myChunk       *simworld.ChunkData
	myTiledChunk  [16][16]simdata.TilemapKey
	myChunkCoords simdata.ChunkCoord
	drawnChunk    *ebiten.Image
	entityImage   *ebiten.Image
	upToDate      bool
}

//TODO have multiple tilemaps indexed by the sheet variable
func BlitTile(dest, tilemap *ebiten.Image, x, y int, tileID simdata.TilemapKey) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(x*tileSize), float64(y*tileSize))
	sx, sy := CalcTileMapPosition(tileID.Tile)

	dest.DrawImage(tilemap.SubImage(image.Rect(sx, sy, sx+tileSize, sy+tileSize)).(*ebiten.Image), op)
}

func (c *ChunkRenderer) ChunkToImage(chunkImage *ebiten.Image, w *simworld.World, auto bool) {
	BaseMaterial := simdata.TileNames["green_ground"]
	for y := range c.myChunk.Chunk.ChunkData {
		for x := range c.myChunk.Chunk.ChunkData[y] {
			BlitTile(chunkImage, TileMap, x, y, BaseMaterial)

			m := c.myChunk.Chunk.ChunkData[y][x].Material
			if auto {
				m = Autotile(x, y, c.myChunkCoords, w, false)
			}
			c.myTiledChunk[y][x] = m
			BlitTile(chunkImage, TileMap, x, y, m)
		}
	}
}

func (c *ChunkRenderer) Draw(screen *ebiten.Image, w *simworld.World, animationTick int, auto bool, cam *Camera) {
	//Update Chunks
	if !c.upToDate {
		c.drawnChunk.Clear()
		c.ChunkToImage(c.drawnChunk, w, auto)
		c.upToDate = true
	}
	//Update entities
	c.DrawEntities(animationTick, cam)

	op := &ebiten.DrawImageOptions{}
	posX, posY, scale := cam.WorldPosToScreenPos(simdata.Location{CC: c.myChunkCoords})

	op.GeoM.Translate(float64(posX), float64(posY))
	op.GeoM.Scale(scale, scale)

	screen.DrawImage(c.drawnChunk, op)
	screen.DrawImage(c.entityImage, op)

}
func (c *ChunkRenderer) DrawEntities(animationTick int, cam *Camera) {
	c.entityImage.Clear()
	for _, v := range c.myChunk.LocalEntities {
		//x, y, z := cam.WorldPosToScreenPos(v.Position())
		//fmt.Println("Drawing", k, "at", x, y, z)
		animFram := animationTick % len(v.AnimationFrames())
		BlitTile(c.entityImage, TileMap, int(v.Position().X), int(v.Position().Y), v.AnimationFrames()[animFram])

	}
}

func (c *Camera) WorldPosToScreenPos(loc simdata.Location) (x, y, scale float64) {
	var posX float64 = float64(loc.CC[0] * tileSize * ChunkSize)
	posX += loc.X
	var posY float64 = float64(loc.CC[1] * tileSize * ChunkSize)
	posY += loc.Y

	posX += float64(c.viewportWidth) / 2
	posY += float64(c.viewportHeight) / 2

	posX -= float64(c.focusPointX)
	posY -= float64(c.focusPointY)

	return posX, posY, float64(c.zoom)

}
func (c *Camera) ScreenPosToWorldPos(x, y float64) simdata.Location {
	var loc simdata.Location
	x /= float64(c.zoom)
	y /= float64(c.zoom)

	x += float64(c.focusPointX)
	y += float64(c.focusPointY)

	x -= float64(c.viewportWidth) / 2
	y -= float64(c.viewportHeight) / 2

	loc.CC[0] = int(x) / (tileSize * ChunkSize)
	loc.CC[1] = int(y) / (tileSize * ChunkSize)

	x -= float64(loc.CC[0] * tileSize * ChunkSize)
	y -= float64(loc.CC[1] * tileSize * ChunkSize)
	loc.X = x / tileSize
	loc.Y = y / tileSize

	return loc
}
func check(err error) {
	if err != nil {
		panic(err)
	}
}
