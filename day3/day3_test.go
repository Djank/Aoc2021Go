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
			name: "Example 1",
			args: `00100
11110
10110
10111
10101
01111
00111
11100
10000
11001
00010
01010`,
			want: 198,
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
