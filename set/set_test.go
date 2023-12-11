package set_test

import (
	"github.com/stretchr/testify/assert"
	"ruan.moolman/aoc2023/set"
	"testing"
)

func TestSetContains(t *testing.T) {
	s := set.NewSet[int]()
	s.Add(2)
	assert.True(t, s.Contains(2))
	assert.False(t, s.Contains(3))
}

func TestSetValues(t *testing.T) {
	s := set.NewSet[int]()
	s.Add(2)
	s.Add(2)
	s.Add(5)
	assert.Equal(t, []int{2, 5}, s.Values())
}

func TestSetCount(t *testing.T) {
	s := set.NewSet[int]()
	assert.Equal(t, s.Count(), 0)
	s.Add(2)
	assert.Equal(t, s.Count(), 1)
	s.Add(2)
	s.Add(5)
	assert.Equal(t, s.Count(), 2)
}
