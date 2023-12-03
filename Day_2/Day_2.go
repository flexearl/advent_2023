package Day_2

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func Part1(lines []string) {
	total := 0
	for i := 0; i < len(lines); i++ {
		valid := true
		colonSplitted := strings.Split(lines[i], ": ")
		semiColonSplitted := strings.Split(colonSplitted[1], "; ")
		for j := 0; j < len(semiColonSplitted) && valid; j++ {
			valid = checkRightAmount(semiColonSplitted[j])
		}
		fmt.Println(valid)
		if valid {
			spaceSplitted := strings.Split(colonSplitted[0], " ")
			id, err := strconv.Atoi(spaceSplitted[1])
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("id", id)
			total += id
		}
	}
	fmt.Println(total)
}

func checkRightAmount(line string) bool {
	maxes := make(map[string]int)
	maxes["red"] = 12
	maxes["green"] = 13
	maxes["blue"] = 14
	valid := true

	commaSplitted := strings.Split(line, ", ")
	fmt.Println(commaSplitted)
	for i := 0; i < len(commaSplitted) && valid; i++ {
		spaceSplitted := strings.Split(commaSplitted[i], " ")
		fmt.Println(spaceSplitted[0], spaceSplitted[1])
		amount, err := strconv.Atoi(spaceSplitted[0])
		fmt.Println(amount)
		if err != nil {
			log.Fatal(err)
		}
		key := spaceSplitted[1]
		key = strings.ReplaceAll(key, " ", "")
		fmt.Println(len(key))
		fmt.Println(maxes[key])
		valid = maxes[spaceSplitted[1]] >= amount

	}
	return valid
}
