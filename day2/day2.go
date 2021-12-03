package aoc

import (
	"log"
	"strconv"
	"strings"

	"github.com/djank/aoc2021go/utils"
)

func Part1() int {
	input := utils.ReadFileToString("day2/input.txt")
	lines := utils.InputToStrings(input)
	return part1Solution(lines)
}

func Part2() int {
	input := utils.ReadFileToString("day2/input.txt")
	lines := utils.InputToStrings(input)
	return part2Solution(lines)
}

func part1Solution(lines []string) int {
	x, z := 0, 0
	for _, line := range lines {
		dx, dz := parseLine(line)
		x += dx
		z += dz
	}
	return x * z
}

func parseLine(line string) (int, int) {
	line = strings.Trim(line, "\r\n")
	parts := strings.Split(line, " ")
	if len(parts) != 2 {
		log.Fatalf("Error on parsing line %s", line)
	}

	value, err := strconv.Atoi(parts[1])
	if err != nil {
		log.Printf("Error on atoi of line %s", line)
		log.Fatal(err)
	}

	switch parts[0] {
	case "forward":
		return value, 0
	case "down":
		return 0, value
	case "up":
		return 0, -value
	default:
		log.Fatalf("%s not recogized. Line %s", parts[0], line)
	}
	return 0, 0
}

func part2Solution(lines []string) int {
	hor, aim, depth := 0, 0, 0
	for _, line := range lines {
		line = strings.Trim(line, "\r\n")
		parts := strings.Split(line, " ")
		if len(parts) != 2 {
			log.Fatalf("Error on parsing line %s", line)
		}
		value, err := strconv.Atoi(parts[1])
		if err != nil {
			log.Printf("Error on atoi of line %s", line)
			log.Fatal(err)
		}
		switch parts[0] {
		case "forward":
			hor += value
			depth += aim * value
		case "down":
			aim += value
		case "up":
			aim -= value
		default:
			log.Fatalf("%s not recogized. Line %s", parts[0], line)
		}
	}
	return hor * depth
}
