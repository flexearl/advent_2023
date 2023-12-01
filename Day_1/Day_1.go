package Day_1

import (
	"fmt"
	"log"
	"strconv"
)

func Part1(lines []string) {
	total := 0
	for i := 0; i < len(lines); i++ {
		foundFirst := false
		firstDigit := ""
		lastDigit := ""
		for j := 0; j < len(lines[i]) && !foundFirst; j++ {
			_, valid := strconv.Atoi(string(lines[i][j]))
			if valid == nil {
				firstDigit = string(lines[i][j])
				foundFirst = true
			}
		}
		foundLast := false
		for j := len(lines[i]) - 1; j >= 0 && !foundLast; j-- {
			_, valid := strconv.Atoi(string(lines[i][j]))
			if valid == nil {
				lastDigit = string(lines[i][j])
				foundLast = true
			}
		}
		sum := firstDigit + lastDigit
		convertedSum, err := strconv.Atoi(sum)
		if err != nil {
			log.Fatal(err)
		}
		total += convertedSum
	}
	fmt.Println(total)
}

func Part2(lines []string) {
	numbers := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	total := 0
	for i := 0; i < len(lines); i++ {
		foundFirst := false
		firstDigit := ""
		lastDigit := ""
		for j := 0; j < len(lines[i]) && !foundFirst; j++ {
			_, valid := strconv.Atoi(string(lines[i][j]))
			if valid == nil {
				firstDigit = string(lines[i][j])
				foundFirst = true
			}
		}
		foundLast := false
		for j := len(lines[i]) - 1; j >= 0 && !foundLast; j-- {
			_, valid := strconv.Atoi(string(lines[i][j]))
			if valid == nil {
				lastDigit = string(lines[i][j])
				foundLast = true
			}
		}
		sum := firstDigit + lastDigit
		convertedSum, err := strconv.Atoi(sum)
		if err != nil {
			log.Fatal(err)
		}
		total += convertedSum
	}
	fmt.Println(total)
}
