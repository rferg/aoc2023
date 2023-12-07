package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/rferg/aoc2023/util"
)

const ASCIIIntOffset = 48

func main() {
	file, err := os.Open("input")
	util.CheckError(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var sum int
	for scanner.Scan() {
		firstDigit, lastDigit, err := extractDigits(scanner.Text())
		util.CheckError(err)
		sum += (firstDigit*10 + lastDigit)
	}

	log.Printf("Answer: %v", sum)
}

func extractDigits(line string) (int, int, error) {
	var digits []int
	for _, c := range line {
		noOffset := int(c) - ASCIIIntOffset
		if noOffset >= 0 && noOffset <= 9 {
			digits = append(digits, noOffset)
		}
	}

	if len(digits) == 0 {
		return 0, 0, fmt.Errorf("no digits found in line %v", line)
	}

	return digits[0], digits[len(digits)-1], nil
}
