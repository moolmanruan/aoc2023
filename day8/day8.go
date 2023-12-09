package day8

import (
	_ "embed"
	"regexp"
	"strconv"
	"strings"
)

//go:embed example.txt
var example string

//go:embed example2.txt
var example2 string

//go:embed input.txt
var input string

type Instruction string

const (
	Left  Instruction = "L"
	Right             = "R"
)

func parseInstructions(s string) []Instruction {
	lines := strings.Split(s, "\n")
	var ii []Instruction
	for _, r := range []rune(lines[0]) {
		if r == 'L' {
			ii = append(ii, Left)
		} else {
			ii = append(ii, Right)
		}
	}
	return ii
}

type Node struct {
	left  string
	right string
}

var nodeRegex, _ = regexp.Compile(`(\w+) = \((\w+), (\w+)\)`)

func parseNodes(s string) map[string]Node {
	lines := strings.Split(s, "\n")
	nn := make(map[string]Node)
	for _, l := range lines[2:] {
		mm := nodeRegex.FindStringSubmatch(l)
		nn[mm[1]] = Node{left: mm[2], right: mm[3]}
	}
	return nn
}

func walk(start string, ii []Instruction, nn map[string]Node, end func(string) bool) int {
	var c int
	var i int
	n := start
	for {
		node := nn[n]
		if ii[i] == Left {
			n = node.left
		} else {
			n = node.right
		}
		c++
		if end(n) {
			break
		}
		i++
		if i >= len(ii) {
			i = 0
		}
	}
	return c
}

func Part1() string {
	data := input
	ii := parseInstructions(data)
	nn := parseNodes(data)
	return strconv.Itoa(walk("AAA", ii, nn, func(i string) bool {
		return i == "ZZZ"
	}))
}

func reachedEnd(nn []string) bool {
	for _, n := range nn {
		if !strings.HasSuffix(n, "Z") {
			return false
		}
	}
	return true
}

func endOnZ(i string) bool {
	return strings.HasSuffix(i, "Z")
}

func ghostWalk(ii []Instruction, nn map[string]Node) int {
	var counts []int
	for k := range nn {
		if strings.HasSuffix(k, "A") {
			count := walk(k, ii, nn, endOnZ)
			counts = append(counts, count)
		}
	}
	return LCM(counts[0], counts[1], counts[2:]...)
}

// GCD returns the greatest common divisor via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// LCM returns the Least Common Multiple via GCD
func LCM(a, b int, vv ...int) int {
	result := a * b / GCD(a, b)
	for i := 0; i < len(vv); i++ {
		result = LCM(result, vv[i])
	}
	return result
}

func Part2() string {
	data := input
	ii := parseInstructions(data)
	nn := parseNodes(data)
	return strconv.Itoa(ghostWalk(ii, nn))
}
