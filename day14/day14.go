package aoc

import (
	"log"
	"strings"

	"github.com/djank/aoc2021go/utils"
)

func Part1() int {
	input := utils.ReadFileToString("day14/input.txt")
	lines := utils.InputToStrings(input)
	return part1Solution(lines)
}

func Part2() int {
	input := utils.ReadFileToString("day14/input.txt")
	lines := utils.InputToStrings(input)
	return part2Solution(lines)
}

func part1Solution(lines []string) int {
	list, transform := ParseInput(lines)
	log.Println("input list of ", len(list))
	log.Println("Transforms of ", len(transform))
	panic("notimplemented")
}

func ParseInput(lines []string) ([]Sym, map[string]rune) {
	firstLine := strings.Trim(lines[0], "\r")
	var syms []Sym
	var old Sym
	for _,ch := range firstLine {
		cur := Sym { char: ch, next: -1 }
		old.next = len(syms)
		syms = append(syms, cur)
		old = cur
	}

	var trans map[string]rune
	for _, line := range lines[3:] {
		parts := strings.Split(line," -> ")
		trans[parts[0]] = rune(parts[1][0])
	}
	return syms, trans
}

type Sym struct {
	char rune
	next int
}

type SymList struct {
	syms []Sym
}

func part2Solution(lines []string) int {
	panic("unimplemented")
}
