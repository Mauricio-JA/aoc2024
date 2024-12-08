package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func filterCorrectUpdates(rules [][2]int, updates [][]int) (corrects [][]int, incorrects [][]int) {
	for _, pages := range updates {
		isCorrect := true
		for _, rule := range rules {
			firstPage := slices.Index(pages, rule[0])
			secondPage := slices.Index(pages, rule[1])
			if firstPage != -1 && secondPage != -1 {
				if firstPage >= secondPage {
					isCorrect = false
					incorrects = append(incorrects, pages)
					break
				}
			}
		}
		if isCorrect {
			corrects = append(corrects, pages)
		}
	}
	return corrects, incorrects
}

func rigthSortedPages(rules [][2]int, pages []int) (result []int) {
	result = pages
	for _, rule := range rules {
		firstPageIndex := slices.Index(pages, rule[0])
		secondPageIndex := slices.Index(pages, rule[1])
		if firstPageIndex != -1 && secondPageIndex != -1 {
			if firstPageIndex >= secondPageIndex {
				firstPageAux := pages[firstPageIndex]
				pages[firstPageIndex] = pages[secondPageIndex]
				pages[secondPageIndex] = firstPageAux
				result = rigthSortedPages(rules, pages)
			}
		}
	}

	return result
}

func main() {
	input, _ := os.ReadFile("day5/input.txt")
	var rows = strings.Split(string(input), "\n")

	//Separate rules and updates
	rules, updates := [][2]int{}, [][]int{}
	isRule := true
	for _, row := range rows {
		if row != "" {
			if isRule {
				rule := strings.Split(row, "|")
				num1, _ := strconv.Atoi(rule[0])
				num2, _ := strconv.Atoi(rule[1])
				rules = append(rules, [2]int{num1, num2})
			} else {
				pages := strings.Split(row, ",")
				pageNums := make([]int, len(pages))
				for i, page := range pages {
					pageNums[i], _ = strconv.Atoi(page)
				}
				updates = append(updates, pageNums)
			}
		} else {
			isRule = false
		}
	}

	sumPart1, sumpart2 := 0, 0
	corrects, incorrects := filterCorrectUpdates(rules, updates)
	for _, pages := range corrects {
		sumPart1 += pages[len(pages)/2]
	}
	for _, pages := range incorrects {
		correctUpdate := rigthSortedPages(rules, pages)
		sumpart2 += correctUpdate[len(correctUpdate)/2]
	}

	fmt.Println("Part 1", sumPart1)
	fmt.Println("part 2", sumpart2)
}
