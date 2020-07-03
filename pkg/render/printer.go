package render

import (
	"fmt"
	"github.com/n-parisi/settlers-of-text/pkg/components"
	"strings"
)

func PrintBoard(board components.Game, showCoords bool) {
	tiles := board.Tiles
	structures := board.Structures

	//Create edges
	edges := make([]string, 54)
	for i := 0; i < 54; i++ {
		if showCoords {
			edges[i] = fmt.Sprintf("%02d", i)
		} else {
			edges[i] = "()"
		}
	}

	//Fill in structures
	for _, s := range structures {
		edges[s.Edge] = getDisplay(s)
	}

	textBoard := genBoard(edges)

	//TODO: Replace these values same way as edges, not with replaceAll
	for i := 0; i < len(tiles); i++ {
		tile := tiles[i]

		tVal := fmt.Sprintf("%02d", tile.Value)
		tType := displayTileType(tile)

		tileId := fmt.Sprintf("%02d", i)

		textBoard = strings.ReplaceAll(textBoard, "V"+tileId, tVal)
		textBoard = strings.ReplaceAll(textBoard, "TYPE"+tileId, tType)
	}

	fmt.Println(textBoard)
}

func getDisplay(s components.Structure) string {
	if s.Type == components.CITY {
		return fmt.Sprint("C", s.Owner)
	} else {
		return fmt.Sprint("S", s.Owner)
	}
}

func displayTileType(t components.Tile) string {
	displayStr := ""
	switch typ := t.Type; typ {
	case components.BRICK:
		displayStr = "BBBBBB"
	case components.LUMBER:
		displayStr = "LLLLLL"
	case components.ORE:
		displayStr = "OOOOOO"
	case components.GRAIN:
		displayStr = "GGGGGG"
	case components.WOOL:
		displayStr = "WWWWWW"
	default:
		displayStr = "      "
	}

	return displayStr
}

func genBoard(e []string) string {
	return "                      " + e[0] + " ---- " + e[1] + "\n" +
		"                      /   V00   \\\n" +
		"                     /          \\\n" +
		"            " + e[2] + " ---- " + e[3] + "  TYPE00  " + e[4] + " ---- " + e[5] + "\n" +
		"            /   V01   \\  TYPE00  /   V11   \\\n" +
		"           /          \\ TYPE00 /          \\\n" +
		"  " + e[6] + " ---- " + e[7] + "  TYPE01  " + e[8] + " ---- " + e[9] + " TYPE11  " + e[10] + " ---- " + e[11] + "\n" +
		"  /   V02   \\  TYPE01  /   V12   \\  TYPE11  /   V10   \\\n" +
		" /          \\ TYPE01 /          \\ TYPE11 /          \\\n" +
		"" + e[12] + "  TYPE02  " + e[13] + " ---- " + e[14] + "  TYPE12  " + e[15] + " ---- " + e[16] + "  TYPE10  " + e[17] + "\n" +
		" \\  TYPE02  /   V13   \\  TYPE12  /   V17   \\  TYPE10  /\n" +
		"  \\ TYPE02 /          \\ TYPE12 /          \\ TYPE10 /\n" +
		"  " + e[18] + " ---- " + e[19] + "  TYPE13  " + e[20] + " ---- " + e[21] + "  TYPE17  " + e[22] + " ---- " + e[23] + "\n" +
		"  /   V03   \\  TYPE13  /   V18   \\  TYPE17  /   V09   \\\n" +
		" /          \\ TYPE13 /          \\ TYPE17 /          \\\n" +
		"" + e[24] + "  TYPE03  " + e[25] + " ---- " + e[26] + "  TYPE18  " + e[27] + " ---- " + e[28] + "  TYPE09  " + e[29] + "\n" +
		" \\  TYPE03  /   V14   \\  TYPE18  /   V16   \\  TYPE09  /\n" +
		"  \\ TYPE03 /          \\ TYPE18 /          \\ TYPE09 /\n" +
		"  " + e[30] + " ---- " + e[31] + "  TYPE14  " + e[32] + " ---- " + e[33] + "  TYPE16  " + e[34] + " ---- " + e[35] + "\n" +
		"  /   V04   \\  TYPE14  /   V15   \\  TYPE16  /   V08   \\\n" +
		" /          \\ TYPE14 /          \\ TYPE16 /          \\\n" +
		"" + e[36] + "  TYPE04  " + e[37] + " ---- " + e[38] + "  TYPE15  " + e[39] + " ---- " + e[40] + "  TYPE08  " + e[41] + "\n" +
		" \\  TYPE04  /   V05   \\  TYPE15  /   V07   \\  TYPE08  /\n" +
		"  \\ TYPE04 /          \\ TYPE15 /          \\ TYPE08 /\n" +
		"  " + e[42] + " ---- " + e[43] + "  TYPE05  " + e[44] + " ---- " + e[45] + "  TYPE07  " + e[46] + " ---- " + e[47] + "\n" +
		"           \\  TYPE05  /   V06   \\  TYPE07  /\n" +
		"            \\ TYPE05 /          \\ TYPE07 /\n" +
		"            " + e[48] + " ---- " + e[49] + "  TYPE06  " + e[50] + " ---- " + e[51] + "\n" +
		"                     \\  TYPE06  /\n" +
		"                      \\ TYPE06 /\n" +
		"                      " + e[52] + " ---- " + e[53]

}
