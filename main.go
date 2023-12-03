package main

import (
	"log"
	"os"
	"strings"

	"github.com/flexearl/advent_2023.git/Day_2"
)

func main() {
	content, err := os.ReadFile("Day_2/Day_2.txt")
	lines := strings.Split(string(content), "\n")
	if err != nil {
		log.Fatal(err)
	}
	Day_2.Part1(lines)
}
