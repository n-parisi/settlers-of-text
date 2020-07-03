package components

import (
	"math/rand"
	"time"
)

type Game struct {
	Tiles      []Tile
	Structures []Structure
}

//Add player struct

type Tile struct {
	Id     int
	Type   int
	Value  int
	Robber bool
}

type Structure struct {
	Type  int
	Edge  int
	Owner int
}

//Surely this is stupid and we can find enum-like pattern / separate files for less confusion
const (
	BRICK  int = 0
	LUMBER int = 1
	ORE    int = 2
	GRAIN  int = 3
	WOOL   int = 4
	DESERT int = 6

	SETTLEMENT int = 7
	CITY       int = 8
)

func (g *Game) AddStructure(sType int, sEdge int, sOwner int) {
	st := Structure{sType, sEdge, sOwner}
	g.Structures = append(g.Structures, st)
}

func NewGame() *Game {
	//Create Tiles
	tiles := make([]Tile, 0)

	//Randomize order of types
	randomTypes := make([]int, len(TileTypes))
	rand.Seed(time.Now().UTC().UnixNano())
	perm := rand.Perm(len(TileTypes))
	for i, v := range perm {
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

	game := Game{tiles, make([]Structure, 0)}
	return &game
}

var TileValues = []int{5, 2, 6, 3, 8, 10, 9, 12, 11, 4, 8, 10, 9, 4, 5, 6, 3, 11}
var TileTypes = []int{BRICK, BRICK, BRICK, LUMBER, LUMBER, LUMBER, LUMBER, ORE, ORE, ORE,
	GRAIN, GRAIN, GRAIN, GRAIN, WOOL, WOOL, WOOL, WOOL, DESERT}
