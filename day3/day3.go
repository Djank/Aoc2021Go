package aoc

import (
	"log"
	"strings"

	"github.com/djank/aoc2021go/utils"
)

func Part1() int {
	input := utils.ReadFileToString("day3/input.txt")
	lines := utils.InputToStrings(input)
	return part1Solution(lines)
}

func Part2() int {
	input := utils.ReadFileToString("day3/input.txt")
	lines := utils.InputToStrings(input)
	return part2Solution(lines)
}

func part1Solution(lines []string) int {
	linelen := len(strings.TrimRight(lines[0], "\r\n"))

	log.Printf("Length of line is %d", linelen)
	
	counter := make([]int, linelen)
	for _, line := range lines {
		line = strings.TrimRight(line, "\r\n")
		for i, r := range line {
			c := string(r)
			if c == "1" {
				counter[i] += 1
			}
		}
	}

	half := len(lines) / 2
	gamma, epsilon := 0, 0
	for _, v := range counter {
		gamma = gamma << 1
		epsilon = epsilon << 1
		if v > half {
			gamma += 1
		} else {
			epsilon += 1
		}
	}
	log.Print(gamma, epsilon)
	log.Print(counter)
	return gamma * epsilon
}

func part2Solution(lines []string) int {
	return 0
}
