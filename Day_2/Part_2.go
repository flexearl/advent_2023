package Day_2

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func Part2(lines []string) {
	total := 0
	for i := 0; i < len(lines); i++ {
		valid := true
		colonSplitted := strings.Split(lines[i], ": ")
		semiColonSplitted := strings.Split(colonSplitted[1], "; ")
		minimum := make(map[string]int)
		minimum["red"] = 0
		minimum["green"] = 0
		minimum["blue"] = 0
		for j := 0; j < len(semiColonSplitted) && valid; j++ {
			minimum = getPower(semiColonSplitted[j], minimum)
		}

		total += (minimum["green"] * minimum["red"] * minimum["blue"])
	}
	fmt.Println(total)
}

func getPower(line string, minimum map[string]int) map[string]int {

	valid := true

	commaSplitted := strings.Split(line, ", ")

	for i := 0; i < len(commaSplitted) && valid; i++ {
		spaceSplitted := strings.Split(commaSplitted[i], " ")

		amount, err := strconv.Atoi(spaceSplitted[0])

		if err != nil {
			log.Fatal(err)
		}
		key := spaceSplitted[1]
		key = strings.TrimSpace(key)
		if amount > minimum[key] {
			minimum[key] = amount
		}
	}
	return minimum
}
