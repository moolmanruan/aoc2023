package day9

import (
	_ "embed"
	"strconv"
	"strings"
)

//go:embed example.txt
var example string

//go:embed input.txt
var input string

type sequence []int

func (s sequence) allZeroes() bool {
	for _, v := range s {
		if v != 0 {
			return false
		}
	}
	return true
}

func mustAtoi(val string) int {
	i, err := strconv.Atoi(val)
	if err != nil {
		panic(err)
	}
	return i
}

func parseSequences(text string) []sequence {
	lines := strings.Split(text, "\n")
	var ss []sequence
	for _, l := range lines {
		nums := strings.Split(l, " ")
		var nn []int
		for _, n := range nums {
			nn = append(nn, mustAtoi(n))
		}
		ss = append(ss, nn)
	}
	return ss
}

func derive(s sequence) sequence {
	var res sequence
	for i := 0; i < len(s)-1; i++ {
		res = append(res, s[i+1]-s[i])
	}
	return res
}

func deriveAll(s sequence) []sequence {
	res := []sequence{s}
	for {
		s = derive(s)
		if s.allZeroes() {
			break
		}
		res = append(res, s)
	}
	return res
}

func last[T any](ll []T) T {
	return ll[len(ll)-1]
}

func extrapolateValue(ss []sequence) int {
	val := last(last(ss))
	for i := len(ss) - 2; i >= 0; i-- {
		val += last(ss[i])
	}
	return val
}

func Part1() string {
	ss := parseSequences(input)

	var ans int
	for _, s := range ss {
		ans += extrapolateValue(deriveAll(s))
	}
	return strconv.Itoa(ans)
}
