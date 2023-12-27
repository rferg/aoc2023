package main

import (
	"bufio"
	"errors"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/rferg/aoc2023/util"
)

func main() {
	file, err := os.Open("../input")
	util.CheckError(err)
	defer file.Close()

	maxRound := Round{red: 12, green: 13, blue: 14}
	scanner := bufio.NewScanner(file)
	var sum int
	for scanner.Scan() {
		game, err := parseGame(scanner.Text())
		util.CheckError(err)
		if game.IsPossible(maxRound) {
			sum += game.id
		}
	}

	log.Printf("Answer: %v", sum)
}

func parseGame(line string) (Game, error) {
	splitOnColon := strings.Split(line, ":")
	header := splitOnColon[0]
	id, err := strconv.Atoi(strings.Split(header, " ")[1])
	if err != nil {
		return Game{}, err
	}

	roundStrings := strings.Split(splitOnColon[1], ";")
	game := Game{id: id}
	for _, roundString := range roundStrings {
		parsed, err := parseRound(roundString)
		if err != nil {
			return game, err
		}
		game.rounds = append(game.rounds, parsed)
	}

	return game, nil
}

func parseRound(roundString string) (Round, error) {
	round := Round{}
	for _, countString := range strings.Split(roundString, ",") {
		trimmed := strings.Trim(countString, " ")
		split := strings.Split(trimmed, " ")
		count, err := strconv.Atoi(split[0])
		if err != nil {
			return round, err
		}

		color := split[1]
		switch color {
		case "red":
			round.red = count
		case "green":
			round.green = count
		case "blue":
			round.blue = count
		default:
			return round, errors.New("Unrecognized color: " + color)
		}
	}
	return round, nil
}

type Game struct {
	id     int
	rounds []Round
}

func (game Game) IsPossible(maxRound Round) bool {
	for _, round := range game.rounds {
		if round.red > maxRound.red || round.blue > maxRound.blue || round.green > maxRound.green {
			return false
		}
	}
	return true
}

type Round struct {
	red   int
	blue  int
	green int
}
