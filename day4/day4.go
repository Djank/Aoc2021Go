package aoc

import (
	"strings"

	"github.com/djank/aoc2021go/utils"
)

func Part1() int {
	input := utils.ReadFileToString("day4/input.txt")
	lines := utils.InputToStrings(input)
	return part1Solution(lines)
}

func Part2() int {
	return 0
}

func part1Solution(lines []string) int {
	numbers, boards := LinesToTypes(lines)

	return 0
}

type Numbers []int

type Cell struct {
	Number int
	Column int
	Row int
}

type Board struct {
	Columns []int
	Row []int
	Cells map[int] Cell
}

func LinesToTypes(lines []string) (Numbers, []Board) {
	numbers := Numbers(utils.SplitLineToIntegers(lines[0], ","))

	return numbers, nil
}