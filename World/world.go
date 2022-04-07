package simworld

import (
	simdata "LittleSim/WorldData"

	"github.com/google/uuid"
)

type ChunkData struct {
	Chunk         *simdata.Chunk
	LocalEntities map[uuid.UUID]simdata.Entity
}

type World struct {
	WorldMap map[simdata.ChunkCoord]*ChunkData
}

func (w *World) TickSimulation() {
}
