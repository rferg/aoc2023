package day04

import (
	"math"
	"strconv"
	"strings"
)

func ParseCard(line string) (Card, error) {
	grossSplit := strings.Split(line, ":")
	idSection := grossSplit[0]
	id, err := strconv.Atoi(strings.Fields(idSection)[1])
	if err != nil {
		return Card{}, err
	}

	numbersSection := grossSplit[1]
	splitNumbers := strings.Split(numbersSection, "|")
	winningNumbers, err := parseIntSet(splitNumbers[0])
	if err != nil {
		return Card{}, err
	}

	cardNumbers, err := parseIntSet(splitNumbers[1])
	if err != nil {
		return Card{}, err
	}

	return Card{Id: id, Numbers: cardNumbers, WinningNumbers: winningNumbers}, nil
}

func parseIntSet(spaceDelimited string) (IntSet, error) {
	intStrings := strings.Fields(spaceDelimited)
	set := make(IntSet)
	for _, intString := range intStrings {
		parsed, err := strconv.Atoi(intString)
		if err != nil {
			return set, err
		}

		set[parsed] = true
	}
	return set, nil
}

type IntSet map[int]bool

type Card struct {
	Id             int
	Numbers        IntSet
	WinningNumbers IntSet
}

func (card Card) Winners() []int {
	return card.WinningNumbers.Intersection(card.Numbers)
}

func (card Card) Score() int {
	winnersCount := len(card.Winners())
	if winnersCount == 0 {
		return 0
	}

	exponent := float64(winnersCount - 1)
	return int(math.Pow(2, exponent))
}

func (set IntSet) Intersection(otherSet IntSet) []int {
	result := []int{}
	for n := range set {
		if otherSet[n] {
			result = append(result, n)
		}
	}
	return result
}
