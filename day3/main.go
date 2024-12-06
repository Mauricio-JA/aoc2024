package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func startWithMul(text []string, i int) bool {
	return strings.Join(text[i:i+4], "") == "mul("
}

func searchDo(text []string, i int) bool {
	return strings.Join(text[i:i+4], "") == "do()"
}

func searchDont(text []string, i int) bool {
	return strings.Join(text[i:i+7], "") == "don't()"
}

func checkDigit(char string) bool {
	if _, err := strconv.Atoi(char); err == nil {
		return true
	} else {
		return false
	}
}

func checkFirstDigit(firstDigit []string, char string) bool {
	return firstDigit == nil && checkDigit(char)
}

func checkComma(comma bool, char string) bool {
	return !comma && char == ","
}

func checkCloseBracket(comma bool, char string) bool {
	return comma && char == ")"
}

func findMultiplication(text []string, i int, posibleIndexes []int) int {
	var (
		firstDigit  []string
		secondDigit []string
		comma       bool
	)
	for _, char := range text[posibleIndexes[i]+4 : posibleIndexes[i]+12] {
		if checkFirstDigit(firstDigit, char) {
			firstDigit = append(firstDigit, char)
			continue
		}
		if !comma && checkDigit(char) {
			firstDigit = append(firstDigit, char)
			continue
		}
		if checkComma(comma, char) {
			comma = true
			continue
		}
		if comma && checkDigit(char) {
			secondDigit = append(secondDigit, char)
			continue
		}
		if checkCloseBracket(comma, char) {
			firsNumber, _ := strconv.Atoi(strings.Join(firstDigit, ""))
			secondNumber, _ := strconv.Atoi(strings.Join(secondDigit, ""))
			return firsNumber * secondNumber
		}
		return 0
	}
	return 0
}

func checkEnabledInstructions(indexToCheck int, instructions map[int]bool, keys []int) bool {
	lastInstruction := true
	for _, v := range keys {
		if indexToCheck > v {
			lastInstruction = instructions[v]
		} else {
			break
		}
	}
	return lastInstruction
}

func main() {
	input, _ := os.ReadFile("day3/input.txt")
	text := strings.Split(string(input), "")
	// fmt.Println(text)
	var posibleIndexes []int
	var instructions = make(map[int]bool)
	for i := range text[:len(text)-8] {
		if startWithMul(text, i) {
			posibleIndexes = append(posibleIndexes, i)
		}
		if searchDo(text, i) {
			instructions[i] = true
		}
		if searchDont(text, i) {
			instructions[i] = false
		}
	}

	keys := make([]int, len(instructions))
	j := 0
	for k := range instructions {
		keys[j] = k
		j++
	}
	sort.Ints(keys)

	sumP1 := 0
	sumP2 := 0

	for i := 0; i < len(posibleIndexes); i++ {
		multiplication := findMultiplication(text, i, posibleIndexes)
		sumP1 += multiplication
		if checkEnabledInstructions(posibleIndexes[i], instructions, keys) {
			sumP2 += multiplication
		}
	}
	fmt.Println("Part  1", sumP1)
	fmt.Println("Part  2", sumP2)

}
