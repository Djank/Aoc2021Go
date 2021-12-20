package aoc

import (
	"testing"

	"github.com/djank/aoc2021go/utils"
)

func Test_part1Solution(t *testing.T) {
	tests := []struct {
		name      string
		args      string
		dimension int
		want      int
	}{
		{
			name: "Example",
			args: `0,9 -> 5,9
8,0 -> 0,8
9,4 -> 3,4
2,2 -> 2,1
7,0 -> 7,4
6,4 -> 2,0
0,9 -> 2,9
3,4 -> 1,4
0,0 -> 8,8
5,5 -> 8,2`,
			dimension: 10,
			want:      5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1Solution(utils.InputToStrings(tt.args), tt.dimension); got != tt.want {
				t.Errorf("part1Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_part2Solution(t *testing.T) {
	tests := []struct {
		name      string
		args      string
		dimension int
		want      int
	}{
		{
			name: "Example",
			args: `0,9 -> 5,9
8,0 -> 0,8
9,4 -> 3,4
2,2 -> 2,1
7,0 -> 7,4
6,4 -> 2,0
0,9 -> 2,9
3,4 -> 1,4
0,0 -> 8,8
5,5 -> 8,2`,
			dimension: 10,
			want:      12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2Solution(utils.InputToStrings(tt.args), tt.dimension); got != tt.want {
				t.Errorf("part2Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
