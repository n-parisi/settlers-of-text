package board

import (
	"math/rand"
	"time"
)

//CONSTANTS
const (
	BRICK  = iota
	LUMBER = iota
	ORE    = iota
	GRAIN  = iota
	WOOL   = iota
	DESERT = iota
)

const (
	SETTLEMENT = iota
	CITY       = iota
)

//BOARD
type Board struct {
	Tiles      []Tile
	Structures []Structure
}

//Standard board
var tileValues = []int{5, 2, 6, 3, 8, 10, 9, 12, 11, 4, 8, 10, 9, 4, 5, 6, 3, 11}
var tileTypes = []int{BRICK, BRICK, BRICK, LUMBER, LUMBER, LUMBER, LUMBER, ORE, ORE, ORE,
	GRAIN, GRAIN, GRAIN, GRAIN, WOOL, WOOL, WOOL, WOOL, DESERT}

func (b *Board) AddStructure(sType int, sEdge int, sOwner int) {
	st := Structure{sType, sEdge, sOwner}
	b.Structures = append(b.Structures, st)
}

func NewBoard() *Board {
	//Create Tiles
	tiles := make([]Tile, 0)

	//Randomize order of types
	randomTypes := make([]int, len(tileTypes))
	rand.Seed(time.Now().UTC().UnixNano())
	perm := rand.Perm(len(tileTypes))
	for i, v := range perm {
		randomTypes[v] = tileTypes[i]
	}

	//Keep track of remaining tile values
	remainingValues := tileValues[:]

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

	game := Board{tiles, make([]Structure, 0)}
	return &game
}

//OTHER STRUCTS
//Add player struct

//Should some of these values be private with getters?
type Tile struct {
	Id       int
	Resource int
	Value    int
	IsRobbed bool
}

type Structure struct {
	Type  int
	Edge  int
	Owner int
}
