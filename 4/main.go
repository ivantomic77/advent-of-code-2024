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

	foundWords := checkForWordsPt1(input)
	fmt.Printf("Result Pt. 1: %d\n", foundWords)

	foundWords = checkForWordsPt2(input)
	fmt.Printf("Result Pt. 2: %d\n", foundWords)
}

// TODO: refactor to be O(1) Space
func checkForWordsPt1(input [][]string) int {
	x := [8]int{-1, -1, -1, 0, 0, 1, 1, 1}
	y := [8]int{-1, 0, 1, -1, 1, -1, 0, 1}

	totalWords := 0

	rows := len(input)
	cols := len(input[0])

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if input[i][j] != "X" {
				continue
			}

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

func checkForWordsPt2(input [][]string) int {
	totalWords := 0

	rows := len(input)
	cols := len(input[0])

	// Search 2D array excluding first and last rows/columns
	for i := 1; i < rows-1; i++ {
		for j := 1; j < cols-1; j++ {
			// if A is found try to find MAS in both diagonals
			if input[i][j] == "A" {
				// check \ diagonal for MAS in both directions
				if (input[i-1][j-1] == "M" && input[i+1][j+1] == "S") || input[i-1][j-1] == "S" && input[i+1][j+1] == "M" {
					// check / diagonal for MAS in both directions
					if (input[i+1][j-1] == "M" && input[i-1][j+1] == "S") || input[i+1][j-1] == "S" && input[i-1][j+1] == "M" {
						totalWords++
					}
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
