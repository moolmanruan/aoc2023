package set_test

import (
	"fmt"
	"ruan.moolman/aoc2023/set"
	"testing"
)

func TestSetContains(t *testing.T) {
	s := set.NewSet[int]()
	s.Add(2)
	if !s.Contains(2) {
		t.Error("Expected Set to contain 2")
	}
	if s.Contains(3) {
		t.Error("Expected Set to NOT contain 3")
	}
}

func TestSetValues(t *testing.T) {
	s := set.NewSet[int]()
	s.Add(2)
	s.Add(2)
	s.Add(5)
	want := []int{2, 5}
	for i, v := range s.Values() {
		if want[i] != v {
			t.Error(fmt.Sprintf("Expected Values to return %v", want))
		}
	}
}
