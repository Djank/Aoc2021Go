package aoc

import (
	"github.com/djank/aoc2021go/utils"
)

func Part1() int {
	input := utils.ReadFileToString("day4/input.txt")
	lines := utils.InputToStrings(input)
	return part1Solution(lines)
}

func Part2() int {
	input := utils.ReadFileToString("day4/input.txt")
	lines := utils.InputToStrings(input)
	return part2Solution(lines)
}

func part1Solution(lines []string) int {
	numbers, boards := ParseInput(lines)

	for _, number := range numbers {
		for _, board := range boards {
			if cell, ok := board.Cells[number]; ok {
				cell.Visit = true
				board.Columns[cell.Column] += 1
				if board.Columns[cell.Column] >= 5 {
					// bingo!
					return calcResult(board, number)
				}
				board.Rows[cell.Row] += 1
				if board.Rows[cell.Row] >= 5 {
					// bingo!
					return calcResult(board, number)
				}
			}
		}
	}
	return -1
}

func part2Solution(lines []string) int {
	numbers, boards := ParseInput(lines)
	winBoard := make([]bool, len(boards))
	var lastWinBoard *Board
	var lastWinNumber int
	for _, number := range numbers {
		for numBoard, board := range boards {
			if winBoard[numBoard] {
				continue
			}
			if cell, ok := board.Cells[number]; ok {
				cell.Visit = true
				board.Columns[cell.Column] += 1
				board.Rows[cell.Row] += 1
				if board.Columns[cell.Column] >= 5 || board.Rows[cell.Row] >= 5 {
					// bingo!
					lastWinBoard = board
					lastWinNumber = number
					winBoard[numBoard] = true
				}
			}
		}
	}
	return calcResult(lastWinBoard, lastWinNumber)
}

func calcResult(winBoard *Board, winNumber int) int {
	sumOfUnmarked := 0
	for _, c := range winBoard.Cells {
		if !c.Visit {
			sumOfUnmarked += c.Number
		}
	}
	return sumOfUnmarked * winNumber
}

type Numbers []int

type Cell struct {
	Number int
	Column int
	Row int
	Visit bool
}

type Board struct {
	Columns [5]int
	Rows [5]int
	Cells map[int] *Cell
}

func ParseInput(lines []string) (Numbers, []*Board) {
	numbers := Numbers(utils.SplitLineToIntegers(lines[0], ","))
	var boards []*Board
	numLine := 1
	for numLine < len(lines) {
		board := Board {
			Cells: map[int] *Cell {},
		}
		for i, row := range lines[numLine+1: numLine+6] {
			columns := utils.SplitLineToIntegers(row, " ")
			for j, v := range columns {
				board.Cells[v] = &Cell {
					Number: v,
					Row: i,
					Column: j,
				}
			}
		}
		boards = append(boards, &board)
		numLine = numLine + 6
	}
	return numbers, boards
}
