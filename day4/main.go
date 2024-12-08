package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

const XMAS = "XMAS"
const XMASlen = len(XMAS)
const MAS = "MAS"
const MASlen = len(MAS)

func B2i(b bool) int {
	if b {
		return 1
	}
	return 0
}

func checkXMAS(wordSearch [][]string, x int, y int, maxW int, maxH int) int {
	count := 0
	// horizontal ➡️
	if y+XMASlen <= maxW {
		row := wordSearch[x]
		word := row[y : y+XMASlen]
		count += B2i(strings.Join(word, "") == XMAS)
	}
	//horizontal reverse ⬅️
	if y+1-XMASlen >= 0 {
		row := make([]string, maxW)
		copy(row, wordSearch[x])
		word := row[y+1-XMASlen : y+1]
		slices.Reverse(word)
		count += B2i(strings.Join(word, "") == XMAS)
	}
	//vertical ⬇️
	if x+XMASlen <= maxH {
		var word []string
		for _, row := range wordSearch[x : x+XMASlen] {
			word = append(word, row[y])
		}
		count += B2i(strings.Join(word, "") == XMAS)
	}
	//vertical reverse ⬆️
	if x+1-XMASlen >= 0 {
		var word []string
		for _, row := range wordSearch[x+1-XMASlen : x+1] {
			word = append(word, row[y])
		}
		slices.Reverse(word)
		count += B2i(strings.Join(word, "") == XMAS)
	}
	//diagonal ↘️
	if y+XMASlen <= maxW && x+XMASlen <= maxH {
		var word []string
		for i := 0; i < XMASlen; i++ {
			word = append(word, wordSearch[x+i][y+i])
		}
		count += B2i(strings.Join(word, "") == XMAS)
	}
	//diagonal ↗️
	if y+XMASlen <= maxW && x+1-XMASlen >= 0 {
		var word []string
		for i := 0; i < XMASlen; i++ {
			word = append(word, wordSearch[x-i][y+i])
		}
		count += B2i(strings.Join(word, "") == XMAS)
	}
	//diagonal ↙️
	if y+1-XMASlen >= 0 && x+XMASlen <= maxH {
		var word []string
		for i := 0; i < XMASlen; i++ {
			word = append(word, wordSearch[x+i][y-i])
		}
		count += B2i(strings.Join(word, "") == XMAS)
	}
	//diagonal ↖️
	if y+1-XMASlen >= 0 && x+1-XMASlen >= 0 {
		var word []string
		for i := 0; i < XMASlen; i++ {
			word = append(word, wordSearch[x-i][y-i])
		}
		count += B2i(strings.Join(word, "") == XMAS)
	}

	return count
}

func checkX_MAS(wordSearch [][]string, x int, y int, maxW int, maxH int) int {
	count := 0
	if y > 0 && x > 0 && y+1 < maxW && x+1 < maxH {
		diagonalA, diagonalB, AReversed, BReversed := make([]string, MASlen), make([]string, MASlen), make([]string, MASlen), make([]string, MASlen)
		for i := 0; i < MASlen; i++ {
			diagonalA[i] = wordSearch[x+i-1][y+i-1]
			diagonalB[i] = wordSearch[x-i+1][y+i-1]
		}

		copy(AReversed, diagonalA)
		copy(BReversed, diagonalB)
		slices.Reverse(AReversed)
		slices.Reverse(BReversed)

		masA := strings.Join(diagonalA, "") == MAS || strings.Join(AReversed, "") == MAS
		masB := strings.Join(diagonalB, "") == MAS || strings.Join(BReversed, "") == MAS
		count += B2i(masA && masB)
	}
	return count
}

func main() {
	input, _ := os.ReadFile("day4/input.txt")
	var rows = strings.Split(string(input), "\n")

	wordSearch := make([][]string, len(rows))
	xIndexes, aIndexes := make([][2]int, 0), make([][2]int, 0)
	width, height := len(rows), 0

	for i, row := range rows {
		items := strings.Split(row, "")
		wordSearch[i] = items
		// search all X positions
		height = len(items)
		for j, char := range items {
			if char == "X" {
				xIndexes = append(xIndexes, [2]int{i, j})
			}
			if char == "A" {
				aIndexes = append(aIndexes, [2]int{i, j})
			}
		}
	}

	countXMAS, countX_MAS := 0, 0
	for _, position := range xIndexes {
		countXMAS += checkXMAS(wordSearch, position[0], position[1], width, height)
	}
	for _, position := range aIndexes {
		countX_MAS += checkX_MAS(wordSearch, position[0], position[1], width, height)
	}

	fmt.Println("Part 1:", countXMAS)
	fmt.Println("Part 2:", countX_MAS)

}
