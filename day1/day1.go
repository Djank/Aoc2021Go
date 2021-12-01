package aoc

import (
	"github.com/djank/aoc2021go/utils"
)

func Part1() int {
	input := utils.ReadFileToString("day1/input.txt")
	ints := utils.InputToIntegers(input)
	return part1Solution(ints)
}

func Part2() int {
	input := utils.ReadFileToString("day1/input.txt")
	ints := utils.InputToIntegers(input)

	return part2Solution(ints)

}

func part1Solution(ints []int) int {
	return increasedCount(ints)
}

func part2Solution(ints []int) int {
	// into slide window of 3
	sums := []int{}
	sums = append(sums, ints[0]+ints[1]+ints[2])
	for i := 1; i < len(ints)-2; i += 1 {
		sums = append(sums, sums[i-1]-ints[i-1]+ints[i+2])
	}
	return increasedCount(sums)
}

func increasedCount(ints []int) int {
	acc := 0
	for i := 1; i < len(ints); i += 1 {
		if ints[i] > ints[i-1] {
			acc += 1
		}
	}
	return acc
}
