package render

import (
	"fmt"
	"github.com/n-parisi/settlers-of-text/pkg/board"
	"strings"
)

const numberOfEdges = 54

func PrintBoard(board board.Board, showCoords bool) {
	tiles, structures := board.Tiles, board.Structures

	edges := make([]string, numberOfEdges)
	resources, values := make([]string, len(tiles)), make([]string, len(tiles))

	//Create edges
	for i := 0; i < len(edges); i++ {
		if showCoords {
			edges[i] = fmt.Sprintf("%02d", i)
		} else {
			edges[i] = "()"
		}
	}

	//Fill in structures
	for _, s := range structures {
		edges[s.Edge] = getStructureDisplay(s)
	}

	//Create tiles
	for i, tile := range tiles {
		tileResource := getTileResourceDisplay(tile)

		var tileValue string
		if tile.IsRobbed {
			tileValue = "XX"
		} else {
			tileValue = fmt.Sprintf("%02d", tile.Value)
		}

		resources[i], values[i] = tileResource, tileValue
	}

	textBoard := genBoard(edges, resources, values)

	//TODO: Replace these values same way as edges, not with replaceAll
	for i := 0; i < len(tiles); i++ {
		tile := tiles[i]

		//tVal := fmt.Sprintf("%02d", tile.Value)
		tType := getTileResourceDisplay(tile)

		tileId := fmt.Sprintf("%02d", i)

		//textBoard = strings.ReplaceAll(textBoard, "V"+tileId, tVal)
		textBoard = strings.ReplaceAll(textBoard, "TYPE"+tileId, tType)
	}

	fmt.Println(textBoard)
}

func getStructureDisplay(s board.Structure) string {
	if s.Type == board.CITY {
		return fmt.Sprint("C", s.Owner)
	} else {
		return fmt.Sprint("S", s.Owner)
	}
}

func getTileResourceDisplay(t board.Tile) string {
	var displayStr string
	switch t.Resource {
	case board.BRICK:
		displayStr = "BBBBBB"
	case board.LUMBER:
		displayStr = "LLLLLL"
	case board.ORE:
		displayStr = "OOOOOO"
	case board.GRAIN:
		displayStr = "GGGGGG"
	case board.WOOL:
		displayStr = "WWWWWW"
	default:
		displayStr = "      "
	}

	return displayStr
}

func genBoard(e, r, v []string) string {
	return "                      " + e[0] + " ---- " + e[1] + "\n" +
		"                      /   " + v[0] + "   \\\n" +
		"                     /          \\\n" +
		"            " + e[2] + " ---- " + e[3] + "  TYPE00  " + e[4] + " ---- " + e[5] + "\n" +
		"            /   " + v[1] + "   \\  TYPE00  /   " + v[11] + "   \\\n" +
		"           /          \\ TYPE00 /          \\\n" +
		"  " + e[6] + " ---- " + e[7] + "  TYPE01  " + e[8] + " ---- " + e[9] + "  TYPE11  " + e[10] + " ---- " + e[11] + "\n" +
		"  /   " + v[2] + "   \\  TYPE01  /   " + v[12] + "   \\  TYPE11  /   " + v[10] + "   \\\n" +
		" /          \\ TYPE01 /          \\ TYPE11 /          \\\n" +
		"" + e[12] + "  TYPE02  " + e[13] + " ---- " + e[14] + "  TYPE12  " + e[15] + " ---- " + e[16] + "  TYPE10  " + e[17] + "\n" +
		" \\  TYPE02  /   " + v[13] + "   \\  TYPE12  /   " + v[17] + "   \\  TYPE10  /\n" +
		"  \\ TYPE02 /          \\ TYPE12 /          \\ TYPE10 /\n" +
		"  " + e[18] + " ---- " + e[19] + "  TYPE13  " + e[20] + " ---- " + e[21] + "  TYPE17  " + e[22] + " ---- " + e[23] + "\n" +
		"  /   " + v[3] + "   \\  TYPE13  /   " + v[18] + "   \\  TYPE17  /   " + v[9] + "   \\\n" +
		" /          \\ TYPE13 /          \\ TYPE17 /          \\\n" +
		"" + e[24] + "  TYPE03  " + e[25] + " ---- " + e[26] + "  TYPE18  " + e[27] + " ---- " + e[28] + "  TYPE09  " + e[29] + "\n" +
		" \\  TYPE03  /   " + v[14] + "   \\  TYPE18  /   " + v[16] + "   \\  TYPE09  /\n" +
		"  \\ TYPE03 /          \\ TYPE18 /          \\ TYPE09 /\n" +
		"  " + e[30] + " ---- " + e[31] + "  TYPE14  " + e[32] + " ---- " + e[33] + "  TYPE16  " + e[34] + " ---- " + e[35] + "\n" +
		"  /   " + v[4] + "   \\  TYPE14  /   " + v[15] + "   \\  TYPE16  /   " + v[8] + "   \\\n" +
		" /          \\ TYPE14 /          \\ TYPE16 /          \\\n" +
		"" + e[36] + "  TYPE04  " + e[37] + " ---- " + e[38] + "  TYPE15  " + e[39] + " ---- " + e[40] + "  TYPE08  " + e[41] + "\n" +
		" \\  TYPE04  /   " + v[5] + "   \\  TYPE15  /   " + v[7] + "   \\  TYPE08  /\n" +
		"  \\ TYPE04 /          \\ TYPE15 /          \\ TYPE08 /\n" +
		"  " + e[42] + " ---- " + e[43] + "  TYPE05  " + e[44] + " ---- " + e[45] + "  TYPE07  " + e[46] + " ---- " + e[47] + "\n" +
		"           \\  TYPE05  /   " + v[6] + "   \\  TYPE07  /\n" +
		"            \\ TYPE05 /          \\ TYPE07 /\n" +
		"            " + e[48] + " ---- " + e[49] + "  TYPE06  " + e[50] + " ---- " + e[51] + "\n" +
		"                     \\  TYPE06  /\n" +
		"                      \\ TYPE06 /\n" +
		"                      " + e[52] + " ---- " + e[53]

}
