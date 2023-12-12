package day12

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCountVariations(t *testing.T) {
	cases := []struct {
		name   string
		input  string
		groups []int
		want   int
	}{
		{
			"single with single option",
			"?????..#.#?",
			[]int{3, 1, 1},
			3,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.want, countVariations(c.input, c.groups, 0))
		})
	}
}

func TestGroupValidPos(t *testing.T) {
	cases := []struct {
		name  string
		input string
		size  int
		want  []int
	}{
		{
			"single with single option",
			"?",
			1,
			[]int{0},
		},
		{
			"single with multiple options",
			"???",
			1,
			[]int{0, 1, 2},
		},
		{
			"double with multiple options",
			"???",
			2,
			[]int{0, 1},
		},
		{
			"single fixed position",
			".#.",
			1,
			[]int{1},
		},
		{
			"double fixed position",
			".?..##.",
			2,
			[]int{4},
		},
		{
			"mixed optional and and fixed position",
			"..?..#.",
			1,
			[]int{2, 5},
		},
		{
			"only returns first valid position",
			"..#..#.",
			1,
			[]int{2},
		},
		{
			"only returns first valid position",
			"??###???",
			3,
			[]int{2},
		},
		{
			"first group bigger than expected",
			".?###...#?????",
			2,
			nil, // shouldn't find anything
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.want, groupValidPosWithStop(c.input, c.size, len(c.input)))
		})
	}
}

func TestSplit(t *testing.T) {
	cases := []struct {
		name   string
		input  string
		groups []int
		want   int
	}{
		{
			"single with single option",
			"?",
			[]int{1},
			1,
		},
		{
			"single with two option",
			"??",
			[]int{1},
			2,
		},
		{
			"single with three option",
			"???",
			[]int{1},
			3,
		},
		{
			"double multi",
			"??.??",
			[]int{1, 1},
			4,
		},
		{
			"example line 2",
			".??..??...?##.",
			[]int{1, 1, 3},
			4,
		},
		{
			"line 993",
			".##????###??#..#????",
			[]int{10, 1, 3},
			1,
		},
		{
			"less groups than needed",
			"#.#",
			[]int{1},
			0,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := split(c.input, c.groups)
			assert.Equal(t, c.want, got)
		})
	}
}
