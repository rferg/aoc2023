package main

import (
	"bufio"
	"log"
	"os"

	"github.com/rferg/aoc2023/day02"
	"github.com/rferg/aoc2023/util"
)

func main() {
	file, err := os.Open("../input")
	util.CheckError(err)
	defer file.Close()

	maxRound := day02.Round{Red: 12, Green: 13, Blue: 14}
	scanner := bufio.NewScanner(file)
	var sum int
	for scanner.Scan() {
		game, err := day02.ParseGame(scanner.Text())
		util.CheckError(err)
		if game.IsPossible(maxRound) {
			sum += game.Id
		}
	}

	log.Printf("Answer: %v", sum)
}
