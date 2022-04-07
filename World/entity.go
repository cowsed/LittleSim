package simworld

import (
	simdata "LittleSim/WorldData"
	"fmt"

	"github.com/google/uuid"
)

/*
type Entity interface {
	Position() Location
	UUID() uuid.UUID
	AnimationFrames() []TilemapKey
	String() string
	SubEntities() []Entity
}
*/

type Flame struct {
	ID           uuid.UUID
	Pos          simdata.Location
	Heat         float64
	Light        float64
	Extinguished bool
	Frames       []simdata.Sprite
}

func (f *Flame) String() string {
	return fmt.Sprintf("Flame - Heat: %.2f - Light: %.2f", f.Heat, f.Light)
}
func (f *Flame) UUID() uuid.UUID {
	return f.ID
}
func (f *Flame) AnimationFrames() []simdata.Sprite {
	return f.Frames
}
func (f *Flame) SubEntities() []simdata.Entity {
	return nil
}
func (f *Flame) Position() simdata.Location {
	return f.Pos
}
