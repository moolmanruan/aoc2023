package day12

import (
	_ "embed"
	"gonum.org/v1/gonum/stat/combin"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

type springReport struct {
	report string
	groups []int
}

func parseSpringReport(line string) springReport {
	parts := strings.Split(line, " ")
	var groups []int
	for _, v := range strings.Split(parts[1], ",") {
		g, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		groups = append(groups, g)
	}
	return springReport{report: parts[0], groups: groups}
}

func sum(vv []int) int {
	var r int
	for _, v := range vv {
		r += v
	}
	return r
}

func groupsAdditional(groups []int) int {
	return sum(groups) - len(groups)
}

func generateReport(size int, combination, groups []int) string {
	var c int
	var out string
	for ci, combo := range combination {
		for c < combo {
			out += "."
			c++
		}
		for i := 0; i < groups[ci]; i++ {
			out += "#"
		}
	}
	for len(out) < size {
		out += "."
	}
	return out[:size]
}

func valid(report, variation string) bool {
	if len(report) != len(variation) {
		panic("variation and report size mismatch")
	}
	for i := 0; i < len(report); i++ {
		if report[i] == '?' {
			continue
		}
		if report[i] != variation[i] {
			return false
		}
	}
	return true
}

func numVariations(report springReport) int {
	additional := groupsAdditional(report.groups)
	gaps := len(report.groups) - 1
	maxLen := len(report.report) - gaps - additional
	cc := combin.Combinations(maxLen, len(report.groups))
	s := len(report.report)
	var count int
	for _, combo := range cc {
		if valid(report.report, generateReport(s, combo, report.groups)) {
			count++
		}
	}
	return count
}

func Part1() string {
	lines := strings.Split(input, "\n")
	var ans int
	for _, line := range lines {
		sp := parseSpringReport(line)
		nv := numVariations(sp)
		ans += nv
	}
	return strconv.Itoa(ans)
}
