package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func combineOperations(numbers []int, acc int) (results []int) {
	if len(numbers) == 1 {
		results = append(results, numbers[0]+acc)
		results = append(results, numbers[0]*acc)

		strArr := []string{strconv.Itoa(acc), strconv.Itoa(numbers[0])}
		concatnumber, _ := strconv.Atoi(strings.Join(strArr, ""))
		results = append(results, concatnumber)
		return results
	}
	if len(numbers) >= 2 {
		sum := combineOperations(numbers[1:], acc+numbers[0])
		results = append(results, sum...)

		if acc != 0 {
			product := combineOperations(numbers[1:], acc*numbers[0])
			results = append(results, product...)

			//Part 2
			strArr := []string{strconv.Itoa(acc), strconv.Itoa(numbers[0])}
			concatnumber, _ := strconv.Atoi(strings.Join(strArr, ""))
			concat := combineOperations(numbers[1:], concatnumber)
			results = append(results, concat...)
		}

	}

	return results

}

func main() {
	input, _ := os.ReadFile("day7/input.txt")
	rows := strings.Split(string(input), "\n")

	sum := 0
	for _, row := range rows {
		equation := strings.Split(row, ":")
		testValue, _ := strconv.Atoi(equation[0])
		list := strings.Split(strings.Trim(equation[1], " "), " ")

		numbers := make([]int, len(list))
		for i, value := range list {
			num, _ := strconv.Atoi(value)
			numbers[i] = num
		}

		results := combineOperations(numbers, 0)

		if slices.Contains(results, testValue) {
			fmt.Println(testValue)
			sum += testValue
		}

	}

	fmt.Println(sum)

}
