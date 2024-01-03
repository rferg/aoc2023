package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/rferg/aoc2023/day04"
	"github.com/rferg/aoc2023/util"
)

func main() {
	file, err := os.Open("../input")
	util.CheckError(err)
	defer file.Close()

	cardMap := make(map[int]day04.Card)
	cardStack := make([]int, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		card, err := day04.ParseCard(scanner.Text())
		util.CheckError(err)
		cardMap[card.Id] = card
		cardStack = append(cardStack, card.Id)
	}

	count := 0
	for ; len(cardStack) > 0; count++ {
		// pop
		id := cardStack[len(cardStack)-1]
		cardStack = cardStack[:len(cardStack)-1]

		card, ok := cardMap[id]
		if !ok {
			log.Fatal("Invalid card id ", id)
		}

		for i := range card.Winners() {
			cardStack = append(cardStack, id+i+1)
		}
	}

	fmt.Println("Answer: ", count)
}
