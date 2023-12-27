package day02

import (
	"errors"
	"strconv"
	"strings"
)

type Round struct {
	Red   int
	Blue  int
	Green int
}

func (round Round) Power() int {
	return round.Red * round.Blue * round.Green
}

func ParseRound(roundString string) (Round, error) {
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
			round.Red = count
		case "green":
			round.Green = count
		case "blue":
			round.Blue = count
		default:
			return round, errors.New("Unrecognized color: " + color)
		}
	}
	return round, nil
}
