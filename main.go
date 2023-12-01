package main

import (
	"log"
	"os"
	"strings"

	"github.com/flexearl/advent_2023.git/Day_1"
)

func main() {
	content, err := os.ReadFile("Day_1/Day_1.txt")
	lines := strings.Split(string(content), "\n")
	if err != nil {
		log.Fatal(err)
	}
	Day_1.Part1(lines)
}
