package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Refactor this horrible mess
func main() {
	ruleMap, pages := getInput()

	correctPageLists := getCorrectPageLists(ruleMap, pages, false)
	pageListSumWithoutCorrection := getSumOfMiddlePages(correctPageLists)

	fmt.Println("Sum Pt.1:", pageListSumWithoutCorrection)

	fixedPageLists := getCorrectPageLists(ruleMap, pages, true)
	pageListSumWithCorrection := getSumOfMiddlePages(fixedPageLists)

	fmt.Println("Sum Pt.2:", pageListSumWithCorrection-pageListSumWithoutCorrection)
}

func getSumOfMiddlePages(pages [][]int) int {
	totalSum := 0

	for _, pageList := range pages {
		if len(pageList) == 0 {
			continue
		}

		midIndex := len(pageList) / 2
		totalSum += pageList[midIndex]
	}

	return totalSum
}

func getCorrectPageLists(rules map[int][]int, pages [][]int, correctionEnabled bool) [][]int {
	var correctPageLists [][]int

	for i := range pages {
		page := checkPage(rules, pages[i], correctionEnabled)
		if page != nil {
			correctPageLists = append(correctPageLists, page)
		}
	}
	return correctPageLists
}

func checkPage(rules map[int][]int, page []int, correctionEnabled bool) []int {
	for i := range page {
		for j := i + 1; j < len(page); j++ {
			rule := rules[page[j]]
			if contains(rule, page[i]) {
				if correctionEnabled {
					newList := applyRule(i, j, page)
					return checkPage(rules, newList, correctionEnabled)
				}
				return nil
			}
		}
	}

	return page
}

// Put second index value before first one and return new corrected list
func applyRule(firstIndex, secondIndex int, list []int) []int {
	value := list[secondIndex]
	list = append(list[:secondIndex], list[secondIndex+1:]...)
	list = append(list[:firstIndex], append([]int{value}, list[firstIndex:]...)...)
	return list
}

func getInput() (map[int][]int, [][]int) {
	isReadingRules := true

	rulesMap := make(map[int][]int)
	var pages [][]int

	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if isReadingRules {
			if line == "" {
				isReadingRules = false
				continue
			}

			chars := strings.Split(line, "|")
			rulesMap[stringToInt(chars[0])] = append(rulesMap[stringToInt(chars[0])], stringToInt(chars[1]))
			continue
		}

		chars := strings.Split(line, ",")
		var pageList []int
		for _, c := range chars {
			pageList = append(pageList, stringToInt(c))
		}
		pages = append(pages, pageList)
	}

	return rulesMap, pages
}

func contains(arr []int, target int) bool {
	for _, num := range arr {
		if num == target {
			return true
		}
	}
	return false
}

func stringToInt(s string) int {
	result, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return result
}
