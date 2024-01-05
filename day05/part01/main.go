package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/rferg/aoc2023/day05"
	"github.com/rferg/aoc2023/util"
)

func main() {
	file, err := os.Open("../input")
	util.CheckError(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	answer, err := day05.Solve(scanner, FirstPartSeedParser{})
	util.CheckError(err)
	fmt.Println("Answer: ", answer)
}

type FirstPartSeedParser struct{}

func (parser FirstPartSeedParser) Parse(line string) ([]int, error) {
	seedStrings := strings.Fields(strings.Split(line, ":")[1])
	seeds := make([]int, 0)
	for _, s := range seedStrings {
		i, err := strconv.Atoi(s)
		if err != nil {
			return seeds, err
		}
		seeds = append(seeds, i)
	}

	return seeds, nil
}
