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
	numbers := []Number{}
	symbols := make(SymbolCollection)
	y := 0
	for scanner.Scan() {
		newNumbers, newSymbolCoordinates, err := parseLine(scanner.Text(), y)
		util.CheckError(err)
		numbers = append(numbers, newNumbers...)
		symbols.Add(newSymbolCoordinates...)
		y++
	}

	var sum int
	for _, number := range numbers {
		if symbols.HasAdjacent(number) {
			sum += number.Value
		}
	}
	log.Printf("Answer: %v", sum)
}

func isInt(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

func parseLine(line string, y int) ([]Number, []Coordinate, error) {
	numbers := []Number{}
	coordinates := []Coordinate{}
	currentNumber := Number{Y: y}
	currentNumberString := ""
	for x, c := range line {
		char := string(rune(c))
		if char == "." {
			if len(currentNumberString) > 0 {
				err := currentNumber.SetValue(currentNumberString, x-1)
				if err != nil {
					return numbers, coordinates, err
				}
				numbers = append(numbers, currentNumber)

				currentNumber = Number{Y: y}
				currentNumberString = ""
			}
		} else if isInt(char) {
			currentNumberString += char
			if currentNumber.MinX == 0 {
				currentNumber.MinX = x
			}
		} else { // is symbol
			coordinates = append(coordinates, Coordinate{X: x, Y: y})

			if len(currentNumberString) > 0 {
				err := currentNumber.SetValue(currentNumberString, x-1)
				if err != nil {
					return numbers, coordinates, err
				}
				numbers = append(numbers, currentNumber)

				currentNumber = Number{Y: y}
				currentNumberString = ""
			}
		}
	}

	if len(currentNumberString) > 0 {
		err := currentNumber.SetValue(currentNumberString, len(line)-1)
		if err != nil {
			return numbers, coordinates, err
		}
		numbers = append(numbers, currentNumber)
	}

	return numbers, coordinates, nil
}

type Number struct {
	MinX  int
	MaxX  int
	Y     int
	Value int
}

func (number *Number) SetValue(str string, maxX int) error {
	number.MaxX = maxX
	num, err := strconv.Atoi(str)
	if err != nil {
		return err
	}
	number.Value = num
	return nil
}

type Coordinate struct {
	X int
	Y int
}

type SymbolCollection map[int][]int

func (collection SymbolCollection) Add(coordinates ...Coordinate) {
	for _, coordinate := range coordinates {
		collection[coordinate.X] = append(collection[coordinate.X], coordinate.Y)
	}
}

func (collection SymbolCollection) Contains(coordinate Coordinate) bool {
	ys, exists := collection[coordinate.X]
	if !exists {
		return false
	}
	for _, y := range ys {
		if y == coordinate.Y {
			return true
		}
	}
	return false
}

func (collection SymbolCollection) HasAdjacent(number Number) bool {
	// Check directly above and below the number
	for x := number.MinX; x <= number.MaxX; x++ {
		if collection.Contains(Coordinate{X: x, Y: number.Y + 1}) ||
			collection.Contains(Coordinate{X: x, Y: number.Y - 1}) {
			return true
		}
	}
	// Check left and right of the number and diagonals
	for _, coord := range []Coordinate{
		{X: number.MinX - 1, Y: number.Y},
		{X: number.MaxX + 1, Y: number.Y},
		{X: number.MinX - 1, Y: number.Y - 1},
		{X: number.MinX - 1, Y: number.Y + 1},
		{X: number.MaxX + 1, Y: number.Y - 1},
		{X: number.MaxX + 1, Y: number.Y + 1},
	} {
		if collection.Contains(coord) {
			return true
		}
	}

	return false
}
