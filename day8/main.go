package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

type node struct {
	X, Y      int
	Frecuency string
}

func checkMapRange(x int, y int, w int, h int) bool {
	return x >= 0 && x < w && y >= 0 && y < h
}

func getAntinodes(antena1 node, antena2 node, maxW int, maxH int) (antinodePair [2]node, antinodeLine []node) {
	distance := [2]int{0, 0}
	distance[0] = antena2.X - antena1.X
	distance[1] = antena2.Y - antena1.Y
	pos1, pos2 := node{antena1.X, antena1.Y, "#"}, node{antena2.X, antena2.Y, "#"}

	for checkMapRange(pos1.X, pos1.Y, maxW, maxH) {

		antinodeLine = append(antinodeLine, pos1)
		pos1.X = pos1.X - distance[0]
		pos1.Y = pos1.Y - distance[1]
		if antinodePair[0].Frecuency == "" && checkMapRange(pos1.X, pos1.Y, maxW, maxH) {
			antinodePair[0] = pos1
		}
	}

	for checkMapRange(pos2.X, pos2.Y, maxW, maxH) {
		antinodeLine = append(antinodeLine, pos2)
		pos2.X = pos2.X + distance[0]
		pos2.Y = pos2.Y + distance[1]
		if antinodePair[1].Frecuency == "" && checkMapRange(pos2.X, pos2.Y, maxW, maxH) {
			antinodePair[1] = pos2
		}
	}

	return antinodePair, antinodeLine
}

func mergeNodeLists(mainList []node, listToMerge []node) []node {
	for _, an := range listToMerge {
		if !slices.Contains(mainList, an) {
			mainList = append(mainList, an)
		}
	}
	return mainList
}

func searchAntinodes(antena node, cityMap [][]string, maxW int, maxH int) (antinodePairs []node, antinodeLines []node) {
	for i, row := range cityMap {
		if i >= antena.X {
			for j, char := range row {
				//Si esta por delante de la antena actual
				if (i == antena.X && j > antena.Y) || i > antena.X {
					if char == antena.Frecuency {
						pair, line := getAntinodes(antena, node{i, j, char}, maxW, maxH)
						//Pushear los antinodos de la parte 1
						if pair[0].Frecuency == "#" && !slices.Contains(antinodePairs, pair[0]) {
							antinodePairs = append(antinodePairs, pair[0])
						}
						if pair[1].Frecuency == "#" && !slices.Contains(antinodePairs, pair[1]) {
							antinodePairs = append(antinodePairs, pair[1])
						}
						//Pushear los antinodos de la parte 2
						antinodeLines = mergeNodeLists(antinodeLines, line)
					}
				}
			}
		}
	}
	return antinodePairs, antinodeLines
}

func main() {
	input, _ := os.ReadFile("day8/input.txt")
	var rows = strings.Split(string(input), "\n")

	maxW, maxH := len(rows), len(rows[0])

	cityMap := make([][]string, 0)

	for _, row := range rows {
		line := strings.Split(row, "")
		cityMap = append(cityMap, line)
	}

	antinodesPart1, antinodesPart2 := []node{}, []node{}
	// countAntinodes := 0
	for i, row := range cityMap {
		for j, char := range row {
			if char != "." {
				pairs, lines := searchAntinodes(node{i, j, char}, cityMap, maxW, maxH)
				antinodesPart1 = mergeNodeLists(antinodesPart1, pairs)
				antinodesPart2 = mergeNodeLists(antinodesPart2, lines)
			}
		}
	}
	fmt.Println("part1", len(antinodesPart1))
	fmt.Println("part2", len(antinodesPart2))

	//Show map de antinodes
	// for _, an := range antinodesPart2 {
	// 	cityMap[an.X][an.Y] = an.Frecuency
	// }
	// for _, an := range cityMap {
	// 	fmt.Println(an)
	// }

}
