package day7

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	cases := []struct {
		name string
		hand []int
		typ  int
	}{
		{"three of a kind", []int{3, 12, 12, 1, 2}, 4},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			handType := HandTypeWithJokers(Hand{c.hand})
			if handType != c.typ {
				t.Error(fmt.Sprintf("Expected type %d, got %d", c.typ, handType))
			}
		})
	}

}
