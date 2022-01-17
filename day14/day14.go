package aoc

import (
	//"log"
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
	return len(lines)
	//return part2Solution(lines)
}

func part1Solution(lines []string) int {
	list, transform := ParseInput(lines)
	
	for i := 0; i < 20; i++ {
		list = Step(list, transform)
	}

	countOfRunes := make(map[rune]int)
	for _, sym := range list {
		countOfRunes[sym.char] = countOfRunes[sym.char] + 1
	}
	min, max := 10000000, 0
	for _, count := range countOfRunes {
		if count < min {
			min = count
		}
		if count > max {
			max = count
		}
	}

	return max - min
}

func Step(list []Sym, transform map[string]rune) []Sym {
	curIndex := 0
	c := list[curIndex]
	for c.next != -1 {
		n := list[c.next]
		pair := string([]rune {c.char, n.char})
		newrune, exist := transform[pair]
		if !exist {
			continue
		}
		middle := Sym { char: newrune, next: c.next}
		// middle у нас будет под индексом len(list)
		c.next = len(list)
		list[curIndex] = c
		list = append(list, middle)
		curIndex = middle.next
		c = n
	}
	return list
}

func ParseInput(lines []string) ([]Sym, map[string]rune) {
	firstLine := strings.Trim(lines[0], "\r")
	var syms []Sym
	oldIndex := -1
	for _,ch := range firstLine {
		cur := Sym { char: ch, next: -1 }
		if oldIndex >= 0 {
			old := syms[oldIndex]
			old.next = len(syms)
			syms[oldIndex] = old	
		}
		syms = append(syms, cur)
		oldIndex += 1
	}

	trans := make(map[string]rune) 
	for _, line := range lines[2:] {
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
