package simworld

import (
	communications "LittleSim/Communications"
	simdata "LittleSim/WorldData"
	"math"
	"math/rand"
)

type World struct {
	WorldMap map[simdata.ChunkCoord]*simdata.Chunk
}

func (w *World) ChunkExists(ID simdata.ChunkCoord) bool {
	_, ok := w.WorldMap[ID]
	return ok
}
func (w *World) GetChunk(ID simdata.ChunkCoord) *simdata.Chunk {
	return w.WorldMap[ID]
}
func (w *World) SetTile(x, y int, ID simdata.ChunkCoord, TileID int) {
	if !w.ChunkExists(ID) {
		return
	}
	w.WorldMap[ID].SetTile(x, y, simdata.EnvBlock{Material: TileID})
	communications.InvalidateChunk(x, y, ID)
}

func GenWorld(n int) *World {
	w := World{
		WorldMap: map[simdata.ChunkCoord]*simdata.Chunk{},
	}
	for y := 0; y < n; y++ {
		for x := 0; x < n; x++ {
			w.WorldMap[simdata.ChunkCoord{x, y}] = GenChunk(x*16, y*16)
		}
	}

	return &w
}

func Dist(x1, y1, x2, y2 int) int {
	return int(math.Sqrt(math.Pow(float64(x2-x1), 2) + math.Pow(float64(y2-y1), 2)))
}

func GenChunk(xStart, yStart int) *simdata.Chunk {
	c := simdata.Chunk{
		ChunkData: [16][16]simdata.EnvBlock{},
	}

	//Make Base Layer
	avail := []int{simdata.TileNames["green_ground"]}
	for y := 0; y < 16; y++ {
		for x := 0; x < 16; x++ {
			//Seed = xStart+yStart
			c.ChunkData[y][x].Material = avail[rand.Intn(1)]
			//Place circular pond in middle of tile
			if Dist(8, 8, x, y) < 8 {
				c.ChunkData[y][x].Material = simdata.TileNames["dirt_path"]
			}
			if Dist(8, 8, x, y) < 4 {
				c.ChunkData[y][x].Material = simdata.TileNames["green_ground"]
			}

		}
	}

	return &c
}
