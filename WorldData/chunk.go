package simdata

const ChunkSize = 16

type ChunkCoord [2]int

type EnvBlock struct {
	Material TilemapKey
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
