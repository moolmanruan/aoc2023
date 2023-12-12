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
			got := countVariations(c.input, c.groups)
			assert.Equal(t, c.want, got)
		})
	}
}
