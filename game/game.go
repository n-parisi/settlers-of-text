package game

import (
	"math/rand"
	"time"
)

type Game struct {
	Tiles []Tile
}

type Tile struct {
	Id int
	Type int
	Value int
	Robber bool
}

const (
	BRICK int = 0
	LUMBER int = 1
	ORE int = 2
	GRAIN int = 3
	WOOL int = 4
	DESERT int = 6
)
var TileValues = []int{5,2,6,3,8,10,9,12,11,4,8,10,9,4,5,6,3,11}
var TileTypes = []int{BRICK, BRICK, BRICK, LUMBER, LUMBER, LUMBER, LUMBER, ORE, ORE, ORE,
	GRAIN, GRAIN, GRAIN, GRAIN, WOOL, WOOL, WOOL, WOOL, DESERT}

func (t Tile) displayType() string {
	displayStr := ""
	switch typ := t.Type; typ {
	case BRICK:
		displayStr = "BBBBBB"
	case LUMBER:
		displayStr = "LLLLLL"
	case ORE:
		displayStr = "OOOOOO"
	case GRAIN:
		displayStr = "GGGGGG"
	case WOOL:
		displayStr = "WWWWWW"
	default:
		displayStr = "      "
	}

	return displayStr
}

func NewGame() *Game {
	//Create Tiles
	tiles := make([]Tile, 0)

	//Randomize order of types
	randomTypes := make([]int, len(TileTypes))
	rand.Seed(time.Now().UTC().UnixNano())
	perm := rand.Perm(len(TileTypes))
	for i,v := range perm {
		randomTypes[v] = TileTypes[i]
	}

	//Keep track of remaining tile values
	remainingValues := TileValues[:]

	for i := 0; i < len(randomTypes); i++ {
		tileType := randomTypes[i]
		tileValue := 0
		robber := false
		if tileType != DESERT {
			tileValue = remainingValues[0]
			remainingValues = remainingValues[1:]
		} else {
			robber = true
		}

		tile := Tile{i, tileType, tileValue, robber}
		tiles = append(tiles, tile)
	}

	game := Game{tiles}
	return &game
}