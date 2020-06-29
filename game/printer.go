package game

import (
	"fmt"
	"strings"
)

func PrintBoard(board Game) {
	tiles := board.Tiles

	for i := 0 ; i < len(tiles); i++ {
		tile := tiles[i]

		tVal := fmt.Sprintf("%02d", tile.Value)
		tType := tile.displayType()

		tileId := fmt.Sprintf("%02d", i)

		textBoard = strings.ReplaceAll(textBoard, "V"+tileId, tVal)
		textBoard = strings.ReplaceAll(textBoard, "TYPE"+tileId, tType)
	}

	fmt.Println(textBoard)
}

var textBoard string =
	"                      00 ---- 01\n" +
	"                      /   V00   \\\n" +
	"                     /          \\\n" +
	"            02 ---- 03  TYPE00  04 ---- 05\n" +
	"            /   V01   \\  TYPE00  /   V11   \\\n" +
	"           /          \\ TYPE00 /          \\\n" +
	"  06 ---- 07  TYPE01   08 ---- 09 TYPE11  10 ---- 11\n" +
	"  /   V02   \\  TYPE01  /   V12   \\  TYPE11  /   V10   \\\n" +
	" /          \\ TYPE01 /          \\ TYPE11 /          \\\n" +
	"12  TYPE02  13 ---- 14  TYPE12  15 ---- 16  TYPE10  17\n" +
	" \\  TYPE02  /   V13   \\  TYPE12  /   V17   \\  TYPE10  /\n" +
	"  \\ TYPE02 /          \\ TYPE12 /          \\ TYPE10 /\n" +
	"  18 ---- 19  TYPE13  20 ---- 21  TYPE17  22 ---- 23\n" +
	"  /   V03   \\  TYPE13  /   V18   \\  TYPE17  /   V09   \\\n" +
	" /          \\ TYPE13 /          \\ TYPE17 /          \\\n" +
	"24  TYPE03  25 ---- 26  TYPE18  27 ---- 28  TYPE09  29\n" +
	" \\  TYPE03  /   V14   \\  TYPE18  /   V16   \\  TYPE09  /\n" +
	"  \\ TYPE03 /   RRRR   \\ TYPE18 /          \\ TYPE09 /\n" +
	"  30 ---- 31  TYPE14  32 ---- 33  TYPE16  34 ---- 35\n" +
	"  /   V04   \\  TYPE14  /   V15   \\  TYPE16  /   V08   \\\n" +
	" /          \\ TYPE14 /          \\ TYPE16 /          \\\n" +
	"36  TYPE04  37 ---- 38  TYPE15  39 ---- 40  TYPE08  41\n" +
	" \\  TYPE04  /   V05   \\  TYPE15  /   V07   \\  TYPE08  /\n" +
	"  \\ TYPE04 /          \\ TYPE15 /          \\ TYPE08 /\n" +
	"  42 ---- 43  TYPE05  44 ---- 45  TYPE07  46 ---- 47\n" +
	"           \\  TYPE05  /   V06   \\  TYPE07  /\n" +
	"            \\ TYPE05 /          \\ TYPE07 /\n" +
	"            48 ---- 49  TYPE06  50 ---- 51\n" +
	"                     \\  TYPE06  /\n" +
	"                      \\ TYPE06 /\n" +
	"                      52 ---- 53"
