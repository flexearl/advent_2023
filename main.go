package main

import (
	"log"
	"os"
	"strings"

	"github.com/flexearl/advent_2023.git/Day_3"
)

func main() {
	content, err := os.ReadFile("Day_3/Day_3.txt")
	lines := strings.Split(string(content), "\n")
	if err != nil {
		log.Fatal(err)
	}
	Day_3.Part1(lines)
}
