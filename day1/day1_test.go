package aoc

import (
	"testing"
)

func Test_part1Solution(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want int
	}{
		{
			name: "Example 1",
			args: []int{
				199,
				200,
				208,
				210,
				200,
				207,
				240,
				269,
				260,
				263,
			},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1Solution(tt.args); got != tt.want {
				t.Errorf("increasedCount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_part2Solution(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want int
	}{
		{
			name: "Example 1",
			args: []int{
				199,
				200,
				208,
				210,
				200,
				207,
				240,
				269,
				260,
				263,
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2Solution(tt.args); got != tt.want {
				t.Errorf("increasedCount() = %v, want %v", got, tt.want)
			}
		})
	}
}
