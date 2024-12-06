package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	inputArr, guardStartLocation := getInput()

	visitedPositions := walk(inputArr, guardStartLocation)
	fmt.Printf("Pt.1 Visisted positions: %d", visitedPositions)
}

func walk(field [][]string, startingPosition []int) int {
	hitWall := false
	currentPosition := startingPosition
	currentDirection := []int{-1, 0}
	visitedPositions := 0

	field[currentPosition[0]][currentPosition[1]] = "X"
	visitedPositions++

	for !hitWall {
		targetRowIndex := currentPosition[0] + currentDirection[0]
		targetColumnIndex := currentPosition[1] + currentDirection[1]

		if targetRowIndex < 0 || targetColumnIndex < 0 || targetRowIndex >= len(field) || targetColumnIndex >= len(field[0]) {
			hitWall = true
		} else {
			hitCharacter := field[targetRowIndex][targetColumnIndex]
			if hitCharacter == "#" {
				currentDirection = rotate90(currentDirection)
			} else if hitCharacter == "." || hitCharacter == "^" {
				visitedPositions++
				field[targetRowIndex][targetColumnIndex] = "X"
				currentPosition[0] = targetRowIndex
				currentPosition[1] = targetColumnIndex
			} else if hitCharacter == "X" {
				currentPosition[0] = targetRowIndex
				currentPosition[1] = targetColumnIndex
			}
		}
	}
	return visitedPositions
}

func rotate90(currentDirection []int) []int {
	if len(currentDirection) != 2 {
		return nil
	}

	x, y := currentDirection[0], currentDirection[1]
	return []int{y, -x}
}

func getInput() ([][]string, []int) {
	var startPosition []int = []int{0, 0}
	var inputArr [][]string

	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var i int = 0

	for scanner.Scan() {
		line := scanner.Text()

		chars := strings.Split(line, "")
		inputArr = append(inputArr, chars)
		guardIndex := findStringIndex(chars, "^")
		if guardIndex != -1 {
			startPosition[0] = i
			startPosition[1] = guardIndex
		}
		i++
	}

	return inputArr, startPosition
}

func findStringIndex(arr []string, target string) int {
	for i, str := range arr {
		if str == target {
			return i
		}
	}
	return -1
}
