package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Advent of Code 2024 Day 1 task -> Brute forced Array edition :(
// The task link: https://adventofcode.com/2024/day/1
func main() {
	leftList, rightList := loadData()

	sort.Ints(leftList)
	sort.Ints(rightList)

	distancePartOne := getDistancePartOne(&leftList, &rightList)
	distancePartTwo := getDistancePartTwo(&leftList, &rightList)

	fmt.Printf("Part One Result: %d\n", distancePartOne)
	fmt.Printf("Part Two Result: %d", distancePartTwo)
}

func loadData() ([]int, []int) {
	file, err := os.Open("input.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var leftList, rightList []int

	for scanner.Scan() {
		line := scanner.Text()
		listItems := strings.Split(line, ",")

		leftListItem, err := strconv.Atoi(listItems[0])
		if err != nil {
			panic(err)
		}
		leftList = append(leftList, leftListItem)

		rightListItem, err := strconv.Atoi(listItems[1])
		if err != nil {
			panic(err)
		}
		rightList = append(rightList, rightListItem)
	}

	return leftList, rightList
}

func getDistancePartOne(leftList *[]int, rightList *[]int) int {
	var distance int

	for i := range *leftList {
		distance += int(math.Abs(float64((*leftList)[i] - (*rightList)[i])))
	}

	return distance
}

func getDistancePartTwo(leftList *[]int, rightList *[]int) int {
	freq := make(map[int]int)
	distance := 0

	for _, num := range *rightList {
		freq[num]++
	}

	for _, num := range *leftList {
		distance += num * freq[num]
	}

	return distance
}
