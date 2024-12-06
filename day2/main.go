package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func checkFirstlevels(num int, prev int) bool {
	differ := math.Abs(float64(num - prev))
	return num != prev && (differ > 0 && differ <= 3)
}

func checkLevels(num int, prev int, increasing bool) bool {
	if num == prev {
		return false
	}
	differ := math.Abs(float64(num - prev))
	if increasing {
		return (num > prev && differ > 0 && differ <= 3)
	} else {
		return (num < prev && differ > 0 && differ <= 3)
	}
}

func checkSafe(levels []string) bool {
	var (
		isSafe     bool = true
		prev       int
		increasing bool = true
	)

	for i, l := range levels {
		num, _ := strconv.Atoi(l)
		switch {
		case i == 1:
			{
				isSafe = checkFirstlevels(num, prev)
				increasing = num > prev
			}
		case i > 1:
			isSafe = checkLevels(num, prev, increasing)
		}
		if !isSafe {
			break
		}
		prev = num
	}
	return isSafe
}

func main() {
	input, _ := os.ReadFile("day2/input.txt")
	var reports = strings.Split(string(input), "\n")
	var countPart1 int
	var countPart2 int

	for _, v := range reports {
		levels := strings.Split(v, " ")
		if checkSafe(levels) {
			countPart1++
		} else {
			var reportProblemDampener [][]string
			for i := 0; i < len(levels); i++ {
				subArray := append([]string{}, levels[:i]...)
				subArray = append(subArray, levels[i+1:]...)
				reportProblemDampener = append(reportProblemDampener, subArray)
			}
			for _, v := range reportProblemDampener {
				if checkSafe(v) {
					countPart2++
					break
				}
			}
		}
	}

	fmt.Println("Part 1: ", countPart1)
	fmt.Println("Part 2: ", +countPart1+countPart2)

}
