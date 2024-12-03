package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	correctSum := getCorrectSum(false)
	fmt.Printf("Result Pt. 1: %d\n", correctSum)

	correctSum = getCorrectSum(true)
	fmt.Printf("Result Pt. 2: %d", correctSum)
}

func getCorrectSum(enableFlags bool) int {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	sums := 0
	enabled := true

	for scanner.Scan() {
		line := scanner.Text()
		potentialCommands := strings.SplitAfter(line, ")")
		for i := range potentialCommands {
			num1, num2, newEnabled := findNumbers(potentialCommands[i], enableFlags)
			if num1 == -1 && num2 == -1 && enableFlags {
				enabled = newEnabled
				continue
			}
			if num1 == -2 && num2 == -2 {
				continue
			}
			if enabled {
				sums += num1 * num2
			}
		}
	}

	return sums
}

func findNumbers(input string, enableFlags bool) (int, int, bool) {
	if enableFlags {
		re := regexp.MustCompile(`don't()`)
		if len(re.FindStringSubmatch(input)) > 0 {
			return -1, -1, false
		}

		re = regexp.MustCompile(`do()`)
		if len(re.FindStringSubmatch(input)) > 0 {
			return -1, -1, true
		}
	}

	re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	matches := re.FindStringSubmatch(input)

	if len(matches) == 3 {
		num1, err1 := strconv.Atoi(matches[1])
		num2, err2 := strconv.Atoi(matches[2])
		if err1 == nil && err2 == nil {
			return num1, num2, true
		}
	}

	return -2, -2, false
}
