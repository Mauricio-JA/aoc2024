package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	input, _ := os.ReadFile("day1/input.txt")

	var pairs = strings.Split(string(input), "\n")

	leftList, rigthList := []int{}, []int{}

	for i := 0; i < len(pairs); i++ {
		var pair = strings.Split(pairs[i], "   ")
		first, _ := strconv.Atoi(pair[0])
		second, _ := strconv.Atoi(pair[1])

		leftList = append(leftList, first)
		rigthList = append(rigthList, second)
	}
	sort.Ints(leftList)
	sort.Ints(rigthList)

	var sumDistance int64 = 0
	for i := 0; i < len(leftList); i++ {
		distance := rigthList[i] - leftList[i]
		sumDistance += int64(math.Abs(float64(distance)))
	}

	fmt.Println("Part 1:", sumDistance)

	var similarityScore = 0
	for i := 0; i < len(leftList); i++ {
		num := leftList[i]
		count := 0
		for j := 0; j < len(rigthList); j++ {
			if num == rigthList[j] {
				count++
			}
		}
		similarityScore += num * count
	}

	fmt.Println("Part 2:", similarityScore)

}
