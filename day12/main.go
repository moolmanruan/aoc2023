package day12

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

type springReport struct {
	report string
	groups []int
}

func Part1() string {
	res := process(1)
	fmt.Println("Part 1: 7705")
	return res
}

func Part2() string {
	res := process(5)
	fmt.Println("Part 2: 50338344809230")
	return res
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

func process(folds int) string {
	lines := strings.Split(input, "\n")
	var ans int
	numLines := len(lines)
	for li, line := range lines {
		fmt.Printf("%d/%d\n", li, numLines)
		sp := parseSpringReport(line)
		sp = unfold(sp, folds)
		ans += countVariations(sp.report, sp.groups)
	}
	return strconv.Itoa(ans)
}

func unfold(sp springReport, numFolds int) springReport {
	var newReportLines []string
	var newGroups []int
	for i := 0; i < numFolds; i++ {
		newReportLines = append(newReportLines, sp.report)
		newGroups = append(newGroups, sp.groups...)
	}
	return springReport{report: strings.Join(newReportLines, "?"), groups: newGroups}
}

const gap = 1

func leftString(s string, i int) string {
	if i <= 0 {
		return ""
	}
	return s[:i]
}
func rightString(s string, i int) string {
	if i >= len(s) {
		return ""
	}
	return s[i:]
}

func leftIntSlice(s []int, i int) []int {
	if i <= 0 {
		return nil
	}
	return s[:i]
}
func rightIntSlice(s []int, i int) []int {
	if i >= len(s) {
		return nil
	}
	return s[i:]
}

func countVariations(report string, groups []int) int {
	if len(groups) > 0 && len(report) == 0 {
		// invalid if we still have groups to place, but no space left
		return 0
	}
	if len(groups) == 0 {
		if len(report) == 0 {
			// if we have no space, but now groups as well, this is valid
			return 1
		} else if !strings.Contains(report, "#") {
			// if there are no forced placements left, this is valid
			return 1
		} else {
			// out of groups, but still need to place, thus invalid
			return 0
		}
	}

	// get the middle group to split the remaining groups
	groupIndex := len(groups) / 2
	group := groups[groupIndex]

	var branches int
	// for each valid position of the group, count up all the variations
	for _, p := range groupValidPos(report, group) {
		// left
		lC := countVariations(leftString(report, p-gap), leftIntSlice(groups, groupIndex))
		if lC == 0 {
			// don't do the right hand side if this is already 0
			continue
		}
		// right
		rC := countVariations(rightString(report, p+group+gap), rightIntSlice(groups, groupIndex+1))
		branches += lC * rC
	}
	return branches
}

func groupValidPos(report string, size int) []int {
	var pp []int
	if size > len(report) {
		return pp
	}
	last := len(report) - size

positionLoop:
	for i := 0; i <= last; i++ {
		// should be open before the group
		if i > 0 {
			if report[i-1] == '#' {
				continue positionLoop
			}
		}
		// should have values for each spot in the group
		for j := 0; j < size; j++ {
			if report[i+j] == '.' {
				continue positionLoop
			}
		}
		// should be open after the group
		iLast := i + size
		if iLast < len(report) {
			if report[iLast] == '#' {
				continue positionLoop
			}
		}
		pp = append(pp, i)
	}
	return pp
}
