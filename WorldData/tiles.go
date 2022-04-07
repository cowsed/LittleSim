package simdata

import "sort"

type Location struct {
	X, Y float64
	CC   ChunkCoord
}

type TilemapKey struct {
	Sheet int
	Tile  int
}

type Sprite struct {
	//In tiles, the dimensions of the sprite
	Width, Height int
	Sprites       [][]TilemapKey
}

var DEFUALT_TILE TilemapKey = TilemapKey{0, 1329}

var TileNames = map[string]TilemapKey{
	"water":    {0, 0},
	"water2":   {0, 1},
	"water_tl": {0, 2},
	"water_t":  {0, 3},
	"water_tr": {0, 4},
	"water_bl": {0, 114},
	"water_b":  {0, 115},
	"water_br": {0, 116},
	"water_cl": {0, 58},
	"water_c":  {0, 59},
	"water_cr": {0, 60},

	"green_ground":  {0, 5},
	"green_ground2": {0, 61},

	"brown_ground":  {0, 6},
	"brown_ground2": {0, 62},

	"gray_ground":  {0, 7},
	"gray_ground2": {0, 63},

	"tan_ground":  {0, 8},
	"tan_ground2": {0, 64},
	"pebbles":     {0, 9},

	"pebbled_grass": {0, 65},

	"water_corner_br": {0, 56},
	"water_corner_bl": {0, 57},
	"water_corner_tr": {0, 112},
	"water_corner_tl": {0, 113},

	"held_water": {0, 227},

	"held_water_bl":        {0, 282},
	"held_water_b":         {0, 283},
	"held_water_br":        {0, 284},
	"held_water_tl":        {0, 170},
	"held_water_t":         {0, 171},
	"held_water_tr":        {0, 172},
	"held_water_cl":        {0, 226},
	"held_water_cr":        {0, 228},
	"held_water_corner_tr": {0, 224},
	"held_water_corner_tl": {0, 225},
	"held_water_corner_br": {0, 168},
	"held_water_corner_bl": {0, 169},

	"red_flower_tl": {0, 338},
	"red_flower_t":  {0, 339},
	"red_flower_tr": {0, 340},
	"red_flower_cl": {0, 394},
	"red_flower":    {0, 395},
	"red_flower_cr": {0, 396},
	"red_flower_bl": {0, 450},
	"red_flower_b":  {0, 451},
	"red_flower_br": {0, 452},

	"red_flower_corner_tl": {0, 393},
	"red_flower_corner_tr": {0, 392},
	"red_flower_corner_bl": {0, 337},
	"red_flower_corner_br": {0, 336},

	"white_flower_tl": {0, 338 + 56*3},
	"white_flower_t":  {0, 339 + 56*3},
	"white_flower_tr": {0, 340 + 56*3},
	"white_flower_cl": {0, 394 + 56*3},
	"white_flower":    {0, 395 + 56*3},
	"white_flower_cr": {0, 396 + 56*3},
	"white_flower_bl": {0, 450 + 56*3},
	"white_flower_b":  {0, 451 + 56*3},
	"white_flower_br": {0, 452 + 56*3},

	"white_flower_corner_tl": {0, 393 + 56*3},
	"white_flower_corner_tr": {0, 392 + 56*3},
	"white_flower_corner_bl": {0, 337 + 56*3},
	"white_flower_corner_br": {0, 336 + 56*3},

	"blue_flower_tl":        {0, 338 + 56*6},
	"blue_flower_t":         {0, 339 + 56*6},
	"blue_flower_tr":        {0, 340 + 56*6},
	"blue_flower_cl":        {0, 394 + 56*6},
	"blue_flower":           {0, 395 + 56*6},
	"blue_flower_cr":        {0, 396 + 56*6},
	"blue_flower_bl":        {0, 450 + 56*6},
	"blue_flower_b":         {0, 451 + 56*6},
	"blue_flower_br":        {0, 452 + 56*6},
	"blue_flower_corner_tl": {0, 393 + 56*6},
	"blue_flower_corner_tr": {0, 392 + 56*6},
	"blue_flower_corner_bl": {0, 337 + 56*6},
	"blue_flower_corner_br": {0, 336 + 56*6},

	"green_bush_tl":        {0, 338 + 56*9},
	"green_bush_t":         {0, 339 + 56*9},
	"green_bush_tr":        {0, 340 + 56*9},
	"green_bush_cl":        {0, 394 + 56*9},
	"green_bush":           {0, 395 + 56*9},
	"green_bush_cr":        {0, 396 + 56*9},
	"green_bush_bl":        {0, 450 + 56*9},
	"green_bush_b":         {0, 451 + 56*9},
	"green_bush_br":        {0, 452 + 56*9},
	"green_bush_corner_tl": {0, 393 + 56*9},
	"green_bush_corner_tr": {0, 392 + 56*9},
	"green_bush_corner_bl": {0, 337 + 56*9},
	"green_bush_corner_br": {0, 336 + 56*9},

	"orange_bush_tl":        {0, 338 + 56*12},
	"orange_bush_t":         {0, 339 + 56*12},
	"orange_bush_tr":        {0, 340 + 56*12},
	"orange_bush_cl":        {0, 394 + 56*12},
	"orange_bush":           {0, 395 + 56*12},
	"orange_bush_cr":        {0, 396 + 56*12},
	"orange_bush_bl":        {0, 450 + 56*12},
	"orange_bush_b":         {0, 451 + 56*12},
	"orange_bush_br":        {0, 452 + 56*12},
	"orange_bush_corner_tl": {0, 393 + 56*12},
	"orange_bush_corner_tr": {0, 392 + 56*12},
	"orange_bush_corner_bl": {0, 337 + 56*12},
	"orange_bush_corner_br": {0, 336 + 56*12},

	"purple_bush_tl":        {0, 338 + 56*15},
	"purple_bush_t":         {0, 339 + 56*15},
	"purple_bush_tr":        {0, 340 + 56*15},
	"purple_bush_cl":        {0, 394 + 56*15},
	"purple_bush":           {0, 395 + 56*15},
	"purple_bush_cr":        {0, 396 + 56*15},
	"purple_bush_bl":        {0, 450 + 56*15},
	"purple_bush_b":         {0, 451 + 56*15},
	"purple_bush_br":        {0, 452 + 56*15},
	"purple_bush_corner_tl": {0, 393 + 56*15},
	"purple_bush_corner_tr": {0, 392 + 56*15},
	"purple_bush_corner_bl": {0, 337 + 56*15},
	"purple_bush_corner_br": {0, 336 + 56*15},

	"dirt_path":         {0, 568},
	"dirt_path_cl":      {0, 567},
	"dirt_path_cr":      {0, 569},
	"dirt_path_t":       {0, 512},
	"dirt_path_tl":      {0, 511},
	"dirt_path_tr":      {0, 513},
	"dirt_path_b":       {0, 624},
	"dirt_path_bl":      {0, 623},
	"dirt_path_br":      {0, 625},
	"dirt_path_b_end":   {0, 621},
	"dirt_path_t_end":   {0, 622},
	"dirt_path_r_end":   {0, 677},
	"dirt_path_l_end":   {0, 678},
	"dirt_path_x":       {0, 457},
	"dirt_path_y":       {0, 401},
	"dirt_path_turn_tr": {0, 455},
	"dirt_path_turn_tl": {0, 456},
	"dirt_path_turn_br": {0, 399},
	"dirt_path_turn_bl": {0, 400},
	"dirt_path_4way":    {0, 679},
	"dirt_path_ltr":     {0, 454},
	"dirt_path_trb":     {0, 453},
	"dirt_path_rbl":     {0, 397},
	"dirt_path_blu":     {0, 398},

	"gray_path":         {0, 568 + 56*6},
	"gray_path_cl":      {0, 567 + 56*6},
	"gray_path_cr":      {0, 569 + 56*6},
	"gray_path_t":       {0, 512 + 56*6},
	"gray_path_tl":      {0, 511 + 56*6},
	"gray_path_tr":      {0, 513 + 56*6},
	"gray_path_b":       {0, 624 + 56*6},
	"gray_path_bl":      {0, 623 + 56*6},
	"gray_path_br":      {0, 625 + 56*6},
	"gray_path_b_end":   {0, 621 + 56*6},
	"gray_path_t_end":   {0, 622 + 56*6},
	"gray_path_r_end":   {0, 677 + 56*6},
	"gray_path_l_end":   {0, 678 + 56*6},
	"gray_path_x":       {0, 457 + 56*6},
	"gray_path_y":       {0, 401 + 56*6},
	"gray_path_turn_tr": {0, 455 + 56*6},
	"gray_path_turn_tl": {0, 456 + 56*6},
	"gray_path_turn_br": {0, 399 + 56*6},
	"gray_path_turn_bl": {0, 400 + 56*6},
	"gray_path_4way":    {0, 679 + 56*6},
	"gray_path_ltr":     {0, 454 + 56*6},
	"gray_path_trb":     {0, 453 + 56*6},
	"gray_path_rbl":     {0, 397 + 56*6},
	"gray_path_blu":     {0, 398 + 56*6},

	"white_path":         {0, 568 + 56*12},
	"white_path_cl":      {0, 567 + 56*12},
	"white_path_cr":      {0, 569 + 56*12},
	"white_path_t":       {0, 512 + 56*12},
	"white_path_tl":      {0, 511 + 56*12},
	"white_path_tr":      {0, 513 + 56*12},
	"white_path_b":       {0, 624 + 56*12},
	"white_path_bl":      {0, 623 + 56*12},
	"white_path_br":      {0, 625 + 56*12},
	"white_path_b_end":   {0, 621 + 56*12},
	"white_path_t_end":   {0, 622 + 56*12},
	"white_path_r_end":   {0, 677 + 56*12},
	"white_path_l_end":   {0, 678 + 56*12},
	"white_path_x":       {0, 457 + 56*12},
	"white_path_y":       {0, 401 + 56*12},
	"white_path_turn_tr": {0, 455 + 56*12},
	"white_path_turn_tl": {0, 456 + 56*12},
	"white_path_turn_br": {0, 399 + 56*12},
	"white_path_turn_bl": {0, 400 + 56*12},
	"white_path_4way":    {0, 679 + 56*12},
	"white_path_ltr":     {0, 454 + 56*12},
	"white_path_trb":     {0, 453 + 56*12},
	"white_path_rbl":     {0, 397 + 56*12},
	"white_path_blu":     {0, 398 + 56*12},

	"campfire_0": {0, 461},
	"campfire_1": {0, 462},
	"campfire_2": {0, 463},

	"small_fireplace_1": {0, 390},
	"small_fireplace_2": {0, 391},

	"medium_fireplace_1": {0, 446},
	"medium_fireplace_2": {0, 447},

	"medium_wood_fireplace_1": {0, 502},
	"medium_wood_fireplace_2": {0, 503},

	"large_fireplace_top_1": {0, 558},
	"large_fireplace_top_2": {0, 559},

	"large_fireplace_bot_1": {0, 558 + 56},
	"large_fireplace_bot_2": {0, 559 + 56},

	"green_tent_bl": {0, 662},
	"green_tent_br": {0, 663},
	"green_tent_tl": {0, 606},
	"green_tent_tr": {0, 607},
}

//Maps the aesthetic tiles back to their worldly counterparts - border to center
var MapToDefualt = map[string]TilemapKey{
	"green_ground":  TileNames["green_ground"],
	"green_ground2": TileNames["green_ground"],

	"dirt_path":    TileNames["dirt_path"],
	"dirt_path_cl": TileNames["dirt_path"],
	"dirt_path_cr": TileNames["dirt_path"],
	"dirt_path_t":  TileNames["dirt_path"],
	"dirt_path_tl": TileNames["dirt_path"],
	"dirt_path_tr": TileNames["dirt_path"],
	"dirt_path_b":  TileNames["dirt_path"],
	"dirt_path_bl": TileNames["dirt_path"],
	"dirt_path_br": TileNames["dirt_path"],

	"gray_path":    TileNames["gray_path"],
	"gray_path_cl": TileNames["gray_path"],
	"gray_path_cr": TileNames["gray_path"],
	"gray_path_t":  TileNames["gray_path"],
	"gray_path_tl": TileNames["gray_path"],
	"gray_path_tr": TileNames["gray_path"],
	"gray_path_b":  TileNames["gray_path"],
	"gray_path_bl": TileNames["gray_path"],
	"gray_path_br": TileNames["gray_path"],

	"white_path":    TileNames["white_path"],
	"white_path_cl": TileNames["white_path"],
	"white_path_cr": TileNames["white_path"],
	"white_path_t":  TileNames["white_path"],
	"white_path_tl": TileNames["white_path"],
	"white_path_tr": TileNames["white_path"],
	"white_path_b":  TileNames["white_path"],
	"white_path_bl": TileNames["white_path"],
	"white_path_br": TileNames["white_path"],

	"white_flower":           TileNames["white_flower"],
	"white_flower_tl":        TileNames["white_flower"],
	"white_flower_t":         TileNames["white_flower"],
	"white_flower_tr":        TileNames["white_flower"],
	"white_flower_cl":        TileNames["white_flower"],
	"white_flower_cr":        TileNames["white_flower"],
	"white_flower_bl":        TileNames["white_flower"],
	"white_flower_b":         TileNames["white_flower"],
	"white_flower_br":        TileNames["white_flower"],
	"white_flower_corner_tl": TileNames["white_flower"],
	"white_flower_corner_tr": TileNames["white_flower"],
	"white_flower_corner_bl": TileNames["white_flower"],
	"white_flower_corner_br": TileNames["white_flower"],

	"blue_flower":           TileNames["blue_flower"],
	"blue_flower_tl":        TileNames["blue_flower"],
	"blue_flower_t":         TileNames["blue_flower"],
	"blue_flower_tr":        TileNames["blue_flower"],
	"blue_flower_cl":        TileNames["blue_flower"],
	"blue_flower_cr":        TileNames["blue_flower"],
	"blue_flower_bl":        TileNames["blue_flower"],
	"blue_flower_b":         TileNames["blue_flower"],
	"blue_flower_br":        TileNames["blue_flower"],
	"blue_flower_corner_tl": TileNames["blue_flower"],
	"blue_flower_corner_tr": TileNames["blue_flower"],
	"blue_flower_corner_bl": TileNames["blue_flower"],
	"blue_flower_corner_br": TileNames["blue_flower"],

	"green_bush":           TileNames["green_bush"],
	"green_bush_tl":        TileNames["green_bush"],
	"green_bush_t":         TileNames["green_bush"],
	"green_bush_tr":        TileNames["green_bush"],
	"green_bush_cl":        TileNames["green_bush"],
	"green_bush_cr":        TileNames["green_bush"],
	"green_bush_bl":        TileNames["green_bush"],
	"green_bush_b":         TileNames["green_bush"],
	"green_bush_br":        TileNames["green_bush"],
	"green_bush_corner_tl": TileNames["green_bush"],
	"green_bush_corner_tr": TileNames["green_bush"],
	"green_bush_corner_bl": TileNames["green_bush"],
	"green_bush_corner_br": TileNames["green_bush"],

	"orange_bush":           TileNames["orange_bush"],
	"orange_bush_tl":        TileNames["orange_bush"],
	"orange_bush_t":         TileNames["orange_bush"],
	"orange_bush_tr":        TileNames["orange_bush"],
	"orange_bush_cl":        TileNames["orange_bush"],
	"orange_bush_cr":        TileNames["orange_bush"],
	"orange_bush_bl":        TileNames["orange_bush"],
	"orange_bush_b":         TileNames["orange_bush"],
	"orange_bush_br":        TileNames["orange_bush"],
	"orange_bush_corner_tl": TileNames["orange_bush"],
	"orange_bush_corner_tr": TileNames["orange_bush"],
	"orange_bush_corner_bl": TileNames["orange_bush"],
	"orange_bush_corner_br": TileNames["orange_bush"],

	"purple_bush":           TileNames["purple_bush"],
	"purple_bush_tl":        TileNames["purple_bush"],
	"purple_bush_t":         TileNames["purple_bush"],
	"purple_bush_tr":        TileNames["purple_bush"],
	"purple_bush_cl":        TileNames["purple_bush"],
	"purple_bush_cr":        TileNames["purple_bush"],
	"purple_bush_bl":        TileNames["purple_bush"],
	"purple_bush_b":         TileNames["purple_bush"],
	"purple_bush_br":        TileNames["purple_bush"],
	"purple_bush_corner_tl": TileNames["purple_bush"],
	"purple_bush_corner_tr": TileNames["purple_bush"],
	"purple_bush_corner_bl": TileNames["purple_bush"],
	"purple_bush_corner_br": TileNames["purple_bush"],

	"red_flower":           TileNames["red_flower"],
	"red_flower_tl":        TileNames["red_flower"],
	"red_flower_t":         TileNames["red_flower"],
	"red_flower_tr":        TileNames["red_flower"],
	"red_flower_cl":        TileNames["red_flower"],
	"red_flower_cr":        TileNames["red_flower"],
	"red_flower_bl":        TileNames["red_flower"],
	"red_flower_b":         TileNames["red_flower"],
	"red_flower_br":        TileNames["red_flower"],
	"red_flower_corner_tl": TileNames["red_flower"],
	"red_flower_corner_tr": TileNames["red_flower"],
	"red_flower_corner_bl": TileNames["red_flower"],
	"red_flower_corner_br": TileNames["red_flower"],

	"water":    TileNames["water"],
	"water2":   TileNames["water"],
	"water_tl": TileNames["water"],
	"water_t":  TileNames["water"],
	"water_tr": TileNames["water"],
	"water_bl": TileNames["water"],
	"water_b":  TileNames["water"],
	"water_br": TileNames["water"],
	"water_cl": TileNames["water"],
	"water_c":  TileNames["water"],
	"water_cr": TileNames["water"],

	"held_water_bl":        TileNames["held_water"],
	"held_water_b":         TileNames["held_water"],
	"held_water_br":        TileNames["held_water"],
	"held_water_tl":        TileNames["held_water"],
	"held_water_t":         TileNames["held_water"],
	"held_water_tr":        TileNames["held_water"],
	"held_water_cl":        TileNames["held_water"],
	"held_water_cr":        TileNames["held_water"],
	"held_water_corner_tr": TileNames["held_water"],
	"held_water_corner_tl": TileNames["held_water"],
	"held_water_corner_br": TileNames["held_water"],
	"held_water_corner_bl": TileNames["held_water"],
}

var ReverseTileNames = map[TilemapKey]string{}
var JustNames = []string{}
var WorldlyTiles = []string{
	"green_ground", "water", "held_water", "green_ground", "dirt_path", "gray_path", "white_path", "red_flower", "white_flower", "blue_flower", "green_bush", "orange_bush", "purple_bush",
}

func init() {
	for k, v := range TileNames {
		ReverseTileNames[v] = k
	}
	for k := range TileNames {
		JustNames = append(JustNames, k)
	}
	sort.Strings(JustNames)

}
