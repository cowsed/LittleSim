package communications

import (
	simdata "LittleSim/WorldData"
)

type InvalidateChunkData struct {
	X, Y int
	ID   simdata.ChunkCoord
}

//Data is x,y,chunk x, chunk y
var InvalidateChunkChannel chan InvalidateChunkData = make(chan InvalidateChunkData, 100)

func InvalidateChunk(x, y int, ID simdata.ChunkCoord) {
	InvalidateChunkChannel <- InvalidateChunkData{x, y, ID}
}
