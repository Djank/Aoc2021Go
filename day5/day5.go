package aoc

import (
	"log"
	"strings"

	"github.com/djank/aoc2021go/utils"
)

func Part1() int {
	input := utils.ReadFileToString("day5/input.txt")
	lines := utils.InputToStrings(input)
	return part1Solution(lines, 1000)
}

func Part2() int {
	input := utils.ReadFileToString("day5/input.txt")
	lines := utils.InputToStrings(input)
	return part2Solution(lines, 1000)
}

func part1Solution(lines []string, dimension int) int {
	lns := parseInput(lines)
	arr := createArray2D(dimension)
	for _, line := range lns {
		from := line.from
		dest := line.to
		if from.x != dest.x && from.y != dest.y {
			continue
		}
		for !from.equal(dest) {
			arr[from.x][from.y] += 1
			from = from.moveTo(dest)
		}
		arr[from.x][from.y] += 1
	}
	return countGreateThen2(arr)
}

func part2Solution(lines []string, dimension int) int {
	lns := parseInput(lines)
	arr := createArray2D(dimension)
	for _, line := range lns {
		from := line.from
		dest := line.to
		for !from.equal(dest) {
			arr[from.x][from.y] += 1
			from = from.moveTo(dest)
		}
		arr[from.x][from.y] += 1
	}
	return countGreateThen2(arr)
}

func (source Point) moveTo(dest Point) Point {
	dx := signOf(dest.x - source.x)
	dy := signOf(dest.y - source.y)
	return Point{source.x + dx, source.y + dy}
}

func (source Point) equal(dest Point) bool {
	return source.x == dest.x && source.y == dest.y
}

func signOf(num int) int {
	switch {
	case num == 0:
		return 0
	case num < 0:
		return -1
	default:
		return 1
	}
}

func countGreateThen2(arr [][]int) int {
	count := 0
	for _, xx := range arr {
		for _, x := range xx {
			if x > 1 {
				count += 1
			}
		}
	}
	return count
}

func createArray2D(dimension int) [][]int {
	arr := make([][]int, dimension)
	for i, _ := range arr {
		arr[i] = make([]int, dimension)
	}
	return arr
}

func parseInput(lines []string) []Line {
	var result []Line
	for i, v := range lines {
		ll := strings.Split(strings.Trim(v, "\r"), " -> ")
		if len(ll) != 2 {
			log.Panicln("Parse trouble. Line ", i)
		}
		from := utils.SplitLineToIntegers(ll[0], ",")
		to := utils.SplitLineToIntegers(ll[1], ",")
		line := Line{
			from: Point{x: from[0], y: from[1]},
			to:   Point{x: to[0], y: to[1]},
		}
		result = append(result, line)
	}
	return result
}

type Point struct {
	x int
	y int
}

type Line struct {
	from Point
	to   Point
}
