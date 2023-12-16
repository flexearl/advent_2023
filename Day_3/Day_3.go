package Day_3

import (
	"fmt"
	"log"
	"strconv"
)

func Part1(lines []string) {

	total := 0
	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[i]); j++ {
			character := lines[i][j]
			if isNumber(character) {
				number := getNumber(lines, []int{i, j})

				if isPartNumber(lines, []int{i, j}, number) {

					convertedNumber, err := strconv.Atoi(number)
					if err != nil {
						log.Fatal(err)
					}
					//fmt.Println("Converted Number", convertedNumber)
					total += convertedNumber
				}

				j += len(number)
			}
		}
	}
	fmt.Println(total)
}

func isNumber(character byte) bool {
	return character >= 48 && character <= 57
}

func isPartNumber(lines []string, index []int, number string) bool {
	isPartNumber := false

	for i := 0; i < len(number) && !isPartNumber; i++ {
		isPartNumber = isAdjacent(lines, []int{index[0], index[1] + i})
	}
	return isPartNumber
}

func getNumber(lines []string, index []int) string {
	number := ""
	isCharacter := true
	for i := 0; i+index[1] < len(lines[index[0]]) && isCharacter; i++ {
		character := lines[index[0]][index[1]+i]
		if isNumber(character) {
			number += string(character)
		} else {
			isCharacter = false
		}
	}
	return number
}

func isAdjacent(grid []string, index []int) bool {
	isAdjacent := false

	surrounding := [][]int{[]int{index[0] - 1, index[1]}, {index[0] - 1, index[1] + 1},
		[]int{index[0], index[1] + 1},
		[]int{index[0] + 1, index[1] + 1},
		[]int{index[0] + 1, index[1]}, []int{index[0] + 1, index[1] - 1},
		[]int{index[0] - 1, index[1] - 1}, []int{index[0], index[1] - 1}}
	for i := 0; i < len(surrounding) && !isAdjacent; i++ {

		if surrounding[i][0] >= 0 && surrounding[i][0] < len(grid) &&
			surrounding[i][1] >= 0 && surrounding[i][1] < len(grid[0]) {
			character := grid[surrounding[i][0]][surrounding[i][1]]

			if character != '.' && (!isNumber(character)) && character != '\n' && character != '\r' {

				isAdjacent = true
			}

		}
	}
	return isAdjacent
}
