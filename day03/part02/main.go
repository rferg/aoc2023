package main

import (
	"bufio"
	"log"
	"os"
	"strconv"

	"github.com/rferg/aoc2023/util"
)

func main() {
	file, err := os.Open("../input")
	util.CheckError(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	grid := [][]string{}
	gears := []Coordinate{}
	y := 0
	for scanner.Scan() {
		row := []string{}
		for x, c := range scanner.Text() {
			char := string(rune(c))
			if char == "*" {
				gears = append(gears, Coordinate{X: x, Y: y})
			}
			row = append(row, char)
		}
		grid = append(grid, row)
		y++
	}

	var sum int
	for _, gear := range gears {
		sum += multiplyAdjacentPair(gear, grid)
	}
	log.Printf("Answer: %v", sum)
}

func multiplyAdjacentPair(gear Coordinate, grid [][]string) int {
	numbers := []int{}
	for y := gear.Y - 1; y <= gear.Y+1 && y >= 0 && y < len(grid); y++ {
		numbers = append(numbers, findNumbersOnLine(gear.X-1, gear.X+1, y, grid)...)
	}

	if len(numbers) != 2 {
		return 0
	}

	return numbers[0] * numbers[1]
}

func findNumbersOnLine(minX int, maxX int, y int, grid [][]string) []int {
	numbers := []int{}
	coordinates := []Coordinate{}
	for x := minX; x <= maxX && x >= 0 && x < len(grid[y]); x++ {
		s := grid[y][x]
		if isInt(s) {
			result, initialCoordinate, err := parseFullInt(x, y, grid)
			util.CheckError(err)
			if !contains(coordinates, initialCoordinate) {
				numbers = append(numbers, result)
				coordinates = append(coordinates, initialCoordinate)
			}
		}
	}
	return numbers
}

func parseFullInt(x int, y int, grid [][]string) (int, Coordinate, error) {
	numString := grid[y][x]
	xCursor := x - 1
	for xCursor >= 0 && isInt(grid[y][xCursor]) {
		numString = grid[y][xCursor] + numString
		xCursor--
	}
	initialCoord := Coordinate{X: xCursor + 1, Y: y}

	xCursor = x + 1
	for xCursor < len(grid[y]) && isInt(grid[y][xCursor]) {
		numString += grid[y][xCursor]
		xCursor++
	}

	result, err := strconv.Atoi(numString)
	if err != nil {
		return result, initialCoord, err
	}

	return result, initialCoord, nil
}

func contains(coordinates []Coordinate, coord Coordinate) bool {
	for _, check := range coordinates {
		if check.X == coord.X && check.Y == coord.Y {
			return true
		}
	}

	return false
}

func isInt(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

type Coordinate struct {
	X int
	Y int
}
