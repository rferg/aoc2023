package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/rferg/aoc2023/util"
)

const ASCIIIntOffset = 48

func main() {
	numbers := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	file, err := os.Open("input")
	util.CheckError(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var sum int
	for scanner.Scan() {
		firstDigit, lastDigit, err := extractDigits(scanner.Text(), numbers)
		util.CheckError(err)
		sum += (firstDigit*10 + lastDigit)
	}

	log.Printf("Answer: %v", sum)
}

func extractDigits(line string, numbers []string) (int, int, error) {
	var digits []int
	current := ""
	for _, c := range line {
		noOffset := int(c) - ASCIIIntOffset
		if noOffset >= 0 && noOffset <= 9 {
			digits = append(digits, noOffset)
			current = ""
		} else {
			if len(current) == 5 { // no number word is longer than 5
				current = current[1:5] + string(c)
			} else {
				current += string(c)
			}
			// This is slow, but whatever.
			foundDigit := checkForNumbers(current, numbers)
			if foundDigit > 0 {
				digits = append(digits, foundDigit)
			}
		}
	}

	if len(digits) == 0 {
		return 0, 0, fmt.Errorf("no digits found in line %v", line)
	}

	return digits[0], digits[len(digits)-1], nil
}

func checkForNumbers(s string, numbers []string) int {
	for i, word := range numbers {
		found := strings.Contains(s, word)
		if found {
			return i + 1
		}
	}
	return 0
}
