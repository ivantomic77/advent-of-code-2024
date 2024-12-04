package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	input := getInput()

	foundWords := checkForWords(input)
	fmt.Printf("Result Pt. 1: %d\n", foundWords)
}

// refactor to be O(1) Space
func checkForWords(input [][]string) int {
	x := [8]int{-1, -1, -1, 0, 0, 1, 1, 1}
	y := [8]int{-1, 0, 1, -1, 1, -1, 0, 1}

	totalWords := 0

	rows := len(input)
	cols := len(input[0])

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			for d := 0; d < 8; d++ {
				var word []string

				for step := 0; step < 4; step++ {
					newX := i + step*x[d]
					newY := j + step*y[d]

					if newX < 0 || newX >= rows || newY < 0 || newY >= cols {
						break
					}

					word = append(word, input[newX][newY])
				}

				if len(word) == 4 {
					totalWords += findWords(word)
				}
			}
		}
	}

	return totalWords
}

func findWords(input []string) int {
	inputString := strings.Join(input, "")
	re := regexp.MustCompile(`XMAS`)
	return len(re.FindStringSubmatch(inputString))
}

func getInput() [][]string {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var result [][]string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		chars := strings.Split(line, "")
		result = append(result, chars)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return result
}
