package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Bruteforced TODO: Refactor
	safeLevels := loadSafeData(false)
	fmt.Printf("Number of safe reports: %d\n", len(safeLevels))

	safeLevels1 := loadSafeData(true)
	fmt.Printf("Number of safe reports with Safety: %d", len(safeLevels1))
}

func loadSafeData(hasSafety bool) [][]int {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var safeLists [][]int

	for scanner.Scan() {
		line := scanner.Text()
		stringListItems := strings.Split(line, " ")
		var listItems []int
		for _, str := range stringListItems {
			num, err := strconv.Atoi(str)
			if err != nil {
				fmt.Println("Error converting string to integer:", err)
				continue
			}
			listItems = append(listItems, num)
		}

		isListSafe, errIndex := isSafe(listItems)

		if isListSafe {
			safeLists = append(safeLists, listItems)
			continue
		}

		if hasSafety {
			for i := errIndex - 1; i < errIndex+2; i++ {
				if i < 0 || i > len(listItems)-1 {
					continue
				}
				tempList := make([]int, len(listItems))
				copy(tempList, listItems)
				tempList = removeElement(tempList, i)

				if safe, _ := isSafe(tempList); safe {
					safeLists = append(safeLists, listItems)
					break
				}
			}
		}
	}
	return safeLists
}

func removeElement(slice []int, index int) []int {
	return append(slice[:index], slice[index+1:]...)
}

func isSafe(listItems []int) (bool, int) {
	if len(listItems) == 2 {
		return true, -1
	}

	direction := getDirection(listItems[0], listItems[1])

	for i := range listItems {
		if i == len(listItems)-1 {
			break
		}

		currentDirection := getDirection(listItems[i], listItems[i+1])

		if currentDirection != direction {
			return false, i
		}

		diff := int(math.Abs(float64(listItems[i] - listItems[i+1])))
		if diff < 1 || diff > 3 {
			return false, i
		}
	}

	return true, -1
}

func getDirection(a, b int) int {
	if a > b {
		return -1
	}
	return 1
}
