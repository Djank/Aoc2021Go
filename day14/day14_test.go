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
			name: "Small example",
			args: `NNCB

CH -> B
HH -> N
CB -> H
NH -> C
HB -> C
HC -> B
HN -> C
NN -> C
BH -> H
NC -> B
NB -> B
BN -> B
BB -> N
BC -> B
CC -> N
CN -> C`,
			want: 1588,
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

func Benchmark_part1(b *testing.B) {
	tests := []struct {
		name string
		args string
		want int
	}{
		{
			name: "Small example",
			args: `NNCB

CH -> B
HH -> N
CB -> H
NH -> C
HB -> C
HC -> B
HN -> C
NN -> C
BH -> H
NC -> B
NB -> B
BN -> B
BB -> N
BC -> B
CC -> N
CN -> C`,
			want: 1588,
		},
	}
	for n := 0; n < b.N; n++ {
		part1Solution(utils.InputToStrings(tests[0].args))
	}
}
