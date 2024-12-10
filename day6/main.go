package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
	"time"
)

func rotateClockwise(direction string) string {
	switch direction {
	case "^":
		return ">"
	case ">":
		return "v"
	case "v":
		return "<"
	case "<":
		return "^"
	default:
		return direction
	}
}

func checkLimit(x int, y int, w int, h int) bool {
	return (x >= 0 && y >= 0 && x < w && y < h)
}

type Guard struct {
	X         int
	Y         int
	direction string
}

func followPath(guard Guard, mappedArea [][]string, maxW int, maxH int, obstacle [2]int) (visitedPositions [][2]int, isLoop bool) {
	visitedPositions = append(visitedPositions, [2]int{guard.X, guard.Y})
	visitedPaths := make([][4]int, 0)

	pathId := [4]int{guard.X, guard.Y, 0, 0}

	for checkLimit(guard.X, guard.Y, maxW, maxH) {
		nextPosition := [2]int{0, 0}
		switch guard.direction {
		case "^":
			nextPosition = [2]int{guard.X - 1, guard.Y}
		case ">":
			nextPosition = [2]int{guard.X, guard.Y + 1}
		case "v":
			nextPosition = [2]int{guard.X + 1, guard.Y}
		case "<":
			nextPosition = [2]int{guard.X, guard.Y - 1}
		default:
			guard.X = -1
			guard.Y = -1
		}
		if checkLimit(nextPosition[0], nextPosition[1], maxW, maxH) {
			if mappedArea[nextPosition[0]][nextPosition[1]] == "#" || nextPosition == obstacle {
				pathId[2] = guard.X
				pathId[3] = guard.Y
				if slices.Contains(visitedPaths, pathId) {
					isLoop = true
					guard.X = -1
					guard.Y = -1
				} else {
					visitedPaths = append(visitedPaths, pathId)
					pathId = [4]int{guard.X, guard.Y, 0, 0}
					guard.direction = rotateClockwise(guard.direction)
				}
			} else {
				guard.X = nextPosition[0]
				guard.Y = nextPosition[1]
				if !slices.Contains(visitedPositions, nextPosition) {
					visitedPositions = append(visitedPositions, nextPosition)
				}
			}
		} else {
			guard.X = -1
			guard.Y = -1
		}
	}

	return visitedPositions, isLoop
}

func main() {
	input, _ := os.ReadFile("day6/input.txt")
	var rows = strings.Split(string(input), "\n")

	//Separate rules and updates
	guardPosition := Guard{0, 0, ""}
	maxW, maxH := len(rows), 0
	mappedArea := make([][]string, maxW)
	for i, row := range rows {
		line := strings.Split(row, "")
		maxH = len(line)
		for j, char := range line {
			if char != "." && char != "#" {
				guardPosition.X = i
				guardPosition.Y = j
				guardPosition.direction = char
			}
			mappedArea[i] = append(mappedArea[i], char)
		}
	}
	visitedPositions, _ := followPath(guardPosition, mappedArea, maxW, maxH, [2]int{-1, -1})
	fmt.Println("Part 1", len(visitedPositions))
	
	countPosibleLoops := 0
	startTime := time.Now()
	// Se podrá optimizar esto?? Respuesta: SI
	// Tiempo anterior: 2m6.37s
	// Tiempo actual: 20.42s
	for _, position := range visitedPositions{
			_, isLoop := followPath(guardPosition, mappedArea, maxW, maxH, [2]int{position[0], position[1]})
			if isLoop {
				countPosibleLoops++
			}
	}

	fmt.Println("Part 2", countPosibleLoops," tooks ", time.Since(startTime))

}