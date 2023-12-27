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

	scanner := bufio.NewScanner(file)
	var sum int
	for scanner.Scan() {
		game, err := day02.ParseGame(scanner.Text())
		util.CheckError(err)
		sum += game.MaxRound().Power()
	}

	log.Printf("Answer: %v", sum)
}
