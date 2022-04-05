package main

import (
	"math"
	"math/rand"
)

type ChunkCoord [2]int

type World struct {
	WorldMap map[ChunkCoord]*Chunk
	rs       *RenderSystem
}

func (w *World) ChunkExists(ID ChunkCoord) bool {
	_, ok := w.WorldMap[ID]
	return ok
}
func (w *World) GetChunk(ID ChunkCoord) *Chunk {
	return gameWorld.WorldMap[ID]
}
func (w *World) SetTile(x, y int, ID ChunkCoord, TileID int) {
	if !w.ChunkExists(ID) {
		return
	}
	gameWorld.WorldMap[ID].SetTile(x, y, EnvBlock{TileID})
	w.rs.InvalidateChunk(ID)
}

func GenWorld(n int) *World {
	w := World{
		WorldMap: map[ChunkCoord]*Chunk{},
	}
	for y := 0; y < n; y++ {
		for x := 0; x < n; x++ {
			w.WorldMap[ChunkCoord{x, y}] = GenChunk(x*16, y*16)
		}
	}

	return &w
}

func Dist(x1, y1, x2, y2 int) int {
	return int(math.Sqrt(math.Pow(float64(x2-x1), 2) + math.Pow(float64(y2-y1), 2)))
}

func GenChunk(xStart, yStart int) *Chunk {
	c := Chunk{
		ChunkData: [16][16]EnvBlock{},
	}

	//Make Base Layer
	avail := []int{TileNames["green_ground"]}
	for y := 0; y < 16; y++ {
		for x := 0; x < 16; x++ {
			//Seed = xStart+yStart
			c.ChunkData[y][x].Material = avail[rand.Intn(1)]
			//Place circular pond in middle of tile
			if Dist(8, 8, x, y) < 8 {
				c.ChunkData[y][x].Material = TileNames["dirt_path"]
			}
			if Dist(8, 8, x, y) < 6 {
				c.ChunkData[y][x].Material = TileNames["green_ground"]
			}

		}
	}

	return &c
}

type EnvBlock struct {
	Material int
}
type Chunk struct {
	ChunkData [16][16]EnvBlock
}

func (c *Chunk) GetTile(x, y int) EnvBlock {
	return c.ChunkData[y][x]
}

func (c *Chunk) SetTile(x, y int, Tile EnvBlock) {
	c.ChunkData[y][x] = Tile
}
