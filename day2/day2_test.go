package aoc

import (
	"testing"

	"github.com/djank/aoc2021go/utils"
)

func Test_part1Solution(t *testing.T) {
	tests := []struct {
		name string
		args string
		want int
	}{
		{
			name: "Example",
			args: `forward 5
down 5
forward 8
up 3
down 8
forward 2`,
			want: 150,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1Solution(utils.InputToStrings(tt.args)); got != tt.want {
				t.Errorf("part1Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_part2Solution(t *testing.T) {
	tests := []struct {
		name string
		args string
		want int
	}{
		{
			name: "Example",
			args: `forward 5
down 5
forward 8
up 3
down 8
forward 2`,
			want: 900,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2Solution(utils.InputToStrings(tt.args)); got != tt.want {
				t.Errorf("part1Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
