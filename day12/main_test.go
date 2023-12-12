package day12

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNumVariations(t *testing.T) {
	cases := []struct {
		name   string
		report springReport
		want   int
	}{
		{
			"only option",
			springReport{"?", []int{1}},
			1,
		},
		{
			"two options",
			springReport{"??", []int{1}},
			2,
		},
		{
			"two groups",
			springReport{"???", []int{1, 1}},
			1,
		},
		{
			"two bigger groups",
			springReport{"??????", []int{2, 3}},
			1,
		},
		{
			"with fixed values",
			springReport{"#.??", []int{1, 1}},
			2,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.want, numVariations(c.report))
		})
	}
}

func TestGenerateReport(t *testing.T) {
	cases := []struct {
		name  string
		size  int
		combo []int
		group []int
		want  string
	}{
		{
			"single a",
			2,
			[]int{0},
			[]int{1},
			"#.",
		},
		{
			"single b",
			2,
			[]int{1},
			[]int{1},
			".#",
		},
		{
			"double a",
			6,
			[]int{0, 1},
			[]int{2, 2},
			"##.##.",
		},
		{
			"double b",
			6,
			[]int{0, 2},
			[]int{2, 2},
			"##..##",
		},
		{
			"double c",
			6,
			[]int{1, 2},
			[]int{2, 2},
			".##.##",
		},
		{
			"multi a",
			12,
			[]int{0, 1, 2},
			[]int{2, 3, 2},
			"##.###.##...",
		},
		{
			"multi b",
			12,
			[]int{1, 3, 4},
			[]int{2, 3, 2},
			".##..###.##.",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.want, generateReport(c.size, c.combo, c.group))
		})
	}
}
