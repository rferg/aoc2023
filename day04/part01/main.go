package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/rferg/aoc2023/day04"
	"github.com/rferg/aoc2023/util"
)

func main() {
	file, err := os.Open("../input")
	util.CheckError(err)
	defer file.Close()

	var sum int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		card, err := day04.ParseCard(scanner.Text())
		util.CheckError(err)
		sum += card.Score()
	}

	fmt.Println("Answer: ", sum)
}
