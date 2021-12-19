package aoc

import "testing"

func Test_part1Solution(t *testing.T) {
	type args struct {
		lines []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1Solution(tt.args.lines); got != tt.want {
				t.Errorf("part1Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
