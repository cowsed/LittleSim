package simgraphics

import (
	simworld "LittleSim/World"
	simdata "LittleSim/WorldData"
	"fmt"
)

//Figure out which part to use if a tile is at a border
type TileSurrounding [3][3]bool

func MirrorX(in TileSurrounding) TileSurrounding {
	out := in
	for y := 0; y < 3; y++ {
		out[y][0] = in[y][2]
		out[y][2] = in[y][0]
	}
	return out
}
func MirrorY(in TileSurrounding) TileSurrounding {
	out := in
	for x := 0; x < 3; x++ {

		out[0][x] = in[2][x]
		out[2][x] = in[0][x]
	}
	return out
}
func RotateCW(in TileSurrounding) TileSurrounding {
	out := in
	for x := 0; x < 3; x++ {
		//New Top Row from Old Left Row
		out[0][x] = in[2-x][0]
		//New Bottom Row from Old Right Row
		out[2][x] = in[2-x][2]
	}

	//Middle row
	//Left Middle from bottom middle
	out[1][0] = in[2][1]
	//Right Middle from top middle
	out[1][2] = in[0][1]
	//Center piece is same

	return out
}

var C = TileSurrounding{
	[3]bool{true, true, true},
	[3]bool{true, true, true},
	[3]bool{true, true, true},
}

//Interior corners
var TL_CORNER = TileSurrounding{
	[3]bool{false, true, true},
	[3]bool{true, true, true},
	[3]bool{true, true, true},
}
var BL_CORNER = MirrorY(TL_CORNER)
var TR_CORNER = MirrorX(TL_CORNER)
var BR_CORNER = MirrorY(TR_CORNER)

//Exterior corners
var TL = TileSurrounding{
	[3]bool{false, false, false},
	[3]bool{false, true, true},
	[3]bool{false, true, true},
}
var TL2 = TileSurrounding{
	[3]bool{false, false, true},
	[3]bool{false, true, true},
	[3]bool{false, true, true},
}

var TL3 = TileSurrounding{
	[3]bool{false, false, false},
	[3]bool{false, true, true},
	[3]bool{true, true, true},
}
var TL4 = TileSurrounding{
	[3]bool{false, false, true},
	[3]bool{false, true, true},
	[3]bool{true, true, true},
}
var TR = MirrorX(TL)

var TR2 = MirrorX(TL2)
var TR3 = MirrorX(TL3)
var TR4 = MirrorX(TL4)

var BL = MirrorY(TL)

var BL2 = MirrorY(TL2)

var BL3 = MirrorY(TL3)
var BL4 = MirrorY(TL4)

var BR = MirrorX(BL)
var BR2 = MirrorX(BL2)
var BR3 = MirrorX(BL3)
var BR4 = MirrorX(BL4)

var L = TileSurrounding{
	[3]bool{false, true, true},
	[3]bool{false, true, true},
	[3]bool{false, true, true},
}
var L2 = TileSurrounding{
	[3]bool{false, true, true},
	[3]bool{false, true, true},
	[3]bool{true, true, true},
}
var L3 = MirrorY(L2)

var R = MirrorX(L)
var R2 = MirrorX(L2)
var R3 = MirrorX(L3)

var B = RotateCW(R)
var B2 = RotateCW(R2)
var B3 = RotateCW(R3)

var T = MirrorY(B)
var T2 = MirrorY(B2)
var T3 = MirrorY(B3)

var T_END = TileSurrounding{
	[3]bool{false, false, false},
	[3]bool{false, true, false},
	[3]bool{false, true, false},
}
var T_END2 = TileSurrounding{
	[3]bool{false, false, false},
	[3]bool{false, true, false},
	[3]bool{false, true, true},
}
var T_END3 = TileSurrounding{
	[3]bool{false, false, false},
	[3]bool{false, true, false},
	[3]bool{true, true, false},
}
var T_END4 = TileSurrounding{
	[3]bool{false, false, false},
	[3]bool{false, true, false},
	[3]bool{true, true, true},
}

var R_END = RotateCW(T_END)
var R_END2 = RotateCW(T_END2)
var R_END3 = RotateCW(T_END3)
var R_END4 = RotateCW(T_END4)

var L_END = MirrorX(R_END)
var L_END2 = MirrorX(R_END2)
var L_END3 = MirrorX(R_END3)
var L_END4 = MirrorX(R_END4)

var B_END = MirrorY(T_END)
var B_END2 = MirrorY(T_END2)
var B_END3 = MirrorY(T_END3)
var B_END4 = MirrorY(T_END4)

var X = TileSurrounding{
	[3]bool{false, false, false},
	[3]bool{true, true, true},
	[3]bool{false, false, false},
}
var X2 = TileSurrounding{
	[3]bool{false, false, true},
	[3]bool{true, true, true},
	[3]bool{false, false, false},
}
var X3 = MirrorY(X2)
var X4 = MirrorX(X2)
var X5 = MirrorX(X3)

var X6 = TileSurrounding{
	[3]bool{true, false, true},
	[3]bool{true, true, true},
	[3]bool{false, false, false},
}
var X7 = MirrorY(X6)
var X8 = TileSurrounding{
	[3]bool{true, false, false},
	[3]bool{true, true, true},
	[3]bool{true, false, false},
}
var X9 = MirrorX(X8)
var X10 = TileSurrounding{
	[3]bool{true, false, true},
	[3]bool{true, true, true},
	[3]bool{false, false, true},
}
var X11 = MirrorX(X10)
var X12 = MirrorY(X10)
var X13 = MirrorY(X11)
var X14 = TileSurrounding{
	[3]bool{true, false, true},
	[3]bool{true, true, true},
	[3]bool{true, false, true},
}
var Y = RotateCW(X)
var Y2 = RotateCW(X2)
var Y3 = RotateCW(X3)
var Y4 = MirrorY(Y2)
var Y5 = MirrorY(Y3)
var Y6 = RotateCW(X6)
var Y7 = RotateCW(X7)
var Y8 = RotateCW(X8)
var Y9 = MirrorY(Y8)
var Y10 = RotateCW(X10)
var Y11 = RotateCW(X11)
var Y12 = MirrorX(Y10)
var Y13 = MirrorX(Y11)
var Y14 = RotateCW(X2)

var BR_TURN = TileSurrounding{
	[3]bool{false, false, false},
	[3]bool{false, true, true},
	[3]bool{false, true, false},
}
var BL_TURN = RotateCW(BR_TURN)
var TL_TURN = RotateCW(BL_TURN)
var TR_TURN = RotateCW(TL_TURN)

var JUNC_4WAY = TileSurrounding{
	[3]bool{false, true, false},
	[3]bool{true, true, true},
	[3]bool{false, true, false},
}

//Left top right
var JUNC_ltr = TileSurrounding{
	[3]bool{false, true, false},
	[3]bool{true, true, true},
	[3]bool{false, false, false},
}

var JUNC_trb = RotateCW(JUNC_ltr)
var JUNC_rbl = RotateCW(JUNC_trb)
var JUNC_blu = RotateCW(JUNC_rbl)

var AutotileGuide = map[TileSurrounding]string{
	C: "c",

	X:   "x",
	X2:  "x",
	X3:  "x",
	X4:  "x",
	X5:  "x",
	X6:  "x",
	X7:  "x",
	X8:  "x",
	X9:  "x",
	X10: "x",
	X11: "x",
	X12: "x",
	X13: "x",
	X14: "x",

	Y:   "y",
	Y2:  "y",
	Y3:  "y",
	Y4:  "y",
	Y5:  "y",
	Y6:  "y",
	Y7:  "y",
	Y8:  "y",
	Y9:  "y",
	Y10: "y",
	Y11: "y",
	Y12: "y",
	Y13: "y",
	Y14: "y",

	JUNC_4WAY: "4way",

	BL_TURN: "turn_bl",
	BR_TURN: "turn_br",
	TL_TURN: "turn_tl",
	TR_TURN: "turn_tr",

	T_END:  "t_end",
	T_END2: "t_end",
	T_END3: "t_end",
	T_END4: "t_end",

	B_END:  "b_end",
	B_END2: "b_end",
	B_END3: "b_end",
	B_END4: "b_end",

	R_END:  "r_end",
	R_END2: "r_end",
	R_END3: "r_end",
	R_END4: "r_end",

	L_END:  "l_end",
	L_END2: "l_end",
	L_END3: "l_end",
	L_END4: "l_end",

	TL_CORNER: "corner_tl",
	BL_CORNER: "corner_bl",
	TR_CORNER: "corner_tr",
	BR_CORNER: "corner_br",

	JUNC_ltr: "ltr",
	JUNC_trb: "trb",
	JUNC_rbl: "rbl",
	JUNC_blu: "blu",

	TL:  "tl",
	TL2: "tl",
	TL3: "tl",
	TL4: "tl",

	TR:  "tr",
	TR2: "tr",
	TR3: "tr",
	TR4: "tr",

	BL:  "bl",
	BL2: "bl",
	BL3: "bl",
	BL4: "bl",

	BR:  "br",
	BR2: "br",
	BR3: "br",
	BR4: "br",

	L:  "cl",
	L2: "cl",
	L3: "cl",

	R:  "cr",
	R2: "cr",
	R3: "cr",

	T:  "t",
	T2: "t",
	T3: "t",

	B:  "b",
	B2: "b",
	B3: "b",
}

func Autotile(cx, cy int, ID simdata.ChunkCoord, world *simworld.World, debug bool) simdata.TilemapKey {
	tracker := [3][3]bool{
		{false, false, false},
		{false, false, false},
		{false, false, false},
	}

	if debug {
		fmt.Println("Autotile", cx, cy)
	}
	me := GetTileSafe(cx, cy, ID, world)
	for y := -1; y <= 1; y++ {
		for x := -1; x <= 1; x++ {

			if me == GetTileSafe(cx+x, cy+y, ID, world) {
				tracker[y+1][x+1] = true
			}

		}
	}
	t := AutotileGuide[tracker]
	if debug {
		fmt.Println(tracker)
	}
	if t == "c" || t == "" {
		if debug {
			fmt.Println(simdata.ReverseTileNames[me])
		}
		return me
	}

	if debug {
		fmt.Println(simdata.ReverseTileNames[me] + "_" + t)
	}
	if tileNum, found := simdata.TileNames[simdata.ReverseTileNames[me]+"_"+t]; found {
		return tileNum
	}
	return me

}
