package main

import (
	"fmt"
	"github.com/n-parisi/settlers-of-text/pkg/board"
	"github.com/n-parisi/settlers-of-text/pkg/render"
)

func main() {
	game := board.NewBoard()
	fmt.Println("Create a new game board!")
	render.PrintBoard(*game, false)
	fmt.Println("Show board with coordinates!")
	render.PrintBoard(*game, true)
	fmt.Println("Place a settlement on Edge 21 for Player 3")
	game.AddStructure(board.SETTLEMENT, 21, 3)
	render.PrintBoard(*game, false)
	fmt.Println("Place a city on Edge 12 for Player 1")
	game.AddStructure(board.CITY, 12, 1)
	render.PrintBoard(*game, false)
}
