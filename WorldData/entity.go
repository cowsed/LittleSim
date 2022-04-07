package simdata

import (
	"github.com/google/uuid"
)

type Entity interface {
	Position() Location
	UUID() uuid.UUID
	AnimationFrames() []TilemapKey
	String() string
	SubEntities() []Entity
}

//Repeats frame[i] ns[i] times
//Example frame 1 * 3, frame 2 * 2
//frame1,frame1,frame1,frame2,frame2
//MakeFrameList([]TilemapKey{DEFUALT_TILE, TileNames["dirt_path"]}, []int{3, 2}): [{0 1329} {0 1329} {0 1329} {0 568} {0 568}]
func MakeFrameList(frames []TilemapKey, ns []int) []TilemapKey {
	//Find the total number of frames needed
	t := 0
	for _, v := range ns {
		t += v
	}
	l := make([]TilemapKey, t)
	frameIndex := 0     //Which frame we're getting from
	sameFramecount := 0 //count of the same frame. if n is 5, will count up to 5

	for i := range l {
		l[i] = frames[frameIndex]
		sameFramecount++
		if sameFramecount == ns[frameIndex] {
			//Go to next frame
			sameFramecount = 0
			frameIndex++
		}
	}

	return l
}
