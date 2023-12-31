package day02

import (
	"strconv"
	"strings"
)

type Game struct {
	Id     int
	Rounds []Round
}

func (game Game) IsPossible(maxRound Round) bool {
	for _, round := range game.Rounds {
		if round.Red > maxRound.Red || round.Blue > maxRound.Blue || round.Green > maxRound.Green {
			return false
		}
	}
	return true
}

func (game Game) MaxRound() Round {
	round := Round{}
	for _, currentRound := range game.Rounds {
		if currentRound.Red > round.Red {
			round.Red = currentRound.Red
		}

		if currentRound.Blue > round.Blue {
			round.Blue = currentRound.Blue
		}

		if currentRound.Green > round.Green {
			round.Green = currentRound.Green
		}
	}

	return round
}

func ParseGame(line string) (Game, error) {
	splitOnColon := strings.Split(line, ":")
	header := splitOnColon[0]
	id, err := strconv.Atoi(strings.Split(header, " ")[1])
	if err != nil {
		return Game{}, err
	}

	roundStrings := strings.Split(splitOnColon[1], ";")
	game := Game{Id: id}
	for _, roundString := range roundStrings {
		parsed, err := ParseRound(roundString)
		if err != nil {
			return game, err
		}
		game.Rounds = append(game.Rounds, parsed)
	}

	return game, nil
}
