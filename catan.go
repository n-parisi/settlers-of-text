package main

import (
	"nparisi.com/user/catan/game"
)

func main() {
	gameInstance := game.NewGame()
	game.PrintBoard(*gameInstance)
}
