package day05

import (
	"bufio"
	"math"
)

type SeedParser interface {
	Parse(line string) ([]int, error)
}

func Solve(scanner *bufio.Scanner, seedParser SeedParser) (int, error) {
	var seeds []int
	var err error
	convLines := make([]string, 0)
	for i := 0; scanner.Scan(); i++ {
		line := scanner.Text()
		if i == 0 {
			seeds, err = seedParser.Parse(line)
			if err != nil {
				return 0, err
			}
		} else if len(line) == 0 {
			if len(convLines) > 0 {
				seeds, err = convert(convLines, seeds)
				if err != nil {
					return 0, err
				}
				convLines = make([]string, 0)
			}
		} else {
			convLines = append(convLines, line)
		}
	}

	if len(convLines) > 0 {
		seeds, err = convert(convLines, seeds)
		if err != nil {
			return 0, err
		}
	}

	return min(seeds), err
}

func convert(lines []string, seeds []int) ([]int, error) {
	newSeeds := make([]int, 0)
	conv, err := ParseConversion(lines)
	if err != nil {
		return newSeeds, err
	}

	for _, seed := range seeds {
		newSeeds = append(newSeeds, conv.Convert(seed))
	}
	return newSeeds, nil
}

func min(ns []int) int {
	m := math.MaxInt32
	for _, n := range ns {
		if n < m {
			m = n
		}
	}

	return m
}
