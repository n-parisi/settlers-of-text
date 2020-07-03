package main

import (
	"fmt"
	"nparisi.com/user/catan/game"
)

func main() {
	gameInstance := game.NewGame()
	fmt.Println("Create a new game board!")
	game.PrintBoard(*gameInstance, false)
	fmt.Println("Show board with coordinates!")
	game.PrintBoard(*gameInstance, true)
	fmt.Println("Place a settlement on Edge 21 for Player 3")
	gameInstance.AddStructure(game.SETTLEMENT, 21, 3)
	game.PrintBoard(*gameInstance, false)
	fmt.Println("Place a city on Edge 12 for Player 1")
	gameInstance.AddStructure(game.CITY, 12, 1)
	game.PrintBoard(*gameInstance, false)
}
