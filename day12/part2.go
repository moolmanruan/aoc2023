package day12

import (
	"fmt"
	intmath "github.com/thomaso-mirodin/intmath/intgr"
	"strconv"
	"strings"
)

func unfold(sp springReport, numFolds int) springReport {
	var newReportLines []string
	var newGroups []int
	for i := 0; i < numFolds; i++ {
		newReportLines = append(newReportLines, sp.report)
		newGroups = append(newGroups, sp.groups...)
	}
	return springReport{report: strings.Join(newReportLines, "?"), groups: newGroups}
}

func Part2Old() string {
	lines := strings.Split(input, "\n")
	var ans int
	for _, line := range lines[:10] {
		sp := parseSpringReport(line)
		sp = unfold(sp, 5)
		c := countVariations(sp.report, sp.groups, 0)
		ans += c
	}
	return strconv.Itoa(ans)
}

func countVariations(report string, groups []int, nesting int) int {
	if len(groups) == 0 {
		panic("no groups")
	}

	minReportSize := sum(groups) + len(groups) - 1 // gaps
	if len(report) < minReportSize {
		panic("report shorter than min size required")
	}

	group := groups[0]
	pp := groupValidPosWithStop(report, group, len(report)-minReportSize)

	isLastGroup := len(groups) == 1

	var count int
	for _, pos := range pp {
		if isLastGroup {
			nextPos := pos + group
			if !strings.Contains(report[nextPos:], "#") {
				count++
			}
		} else {
			nextPos := pos + group + 1
			count += countVariations(report[nextPos:], groups[1:], nesting+1)
		}
	}
	return count
}

func groupValidPosWithStop(report string, size, stopIndex int) []int {
	var pp []int
	if size > len(report) {
		return pp
	}
	last := intmath.Min(len(report)-size, stopIndex)
	firstActual := strings.Index(report, "#")
	if firstActual >= 0 {
		last = intmath.Min(last, firstActual)
	}

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

func sliceMax(vv []int) int {
	v := vv[0]
	for _, val := range vv {
		if val > v {
			v = val
		}
	}
	return v
}

func Part2() string {
	lines := strings.Split(input, "\n")
	var ans int
	for li, line := range lines {
		fmt.Println(li)
		sp := parseSpringReport(line)
		sp = unfold(sp, 5)

		c := split(sp.report, sp.groups)

		ans += c
	}
	return strconv.Itoa(ans)
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

//func minGroups(input string) int {
//
//	for _, c := range input {
//
//	}
//}

func medianIndex(gg []int, val int) int {
	var ii []int
	for i, g := range gg {
		if g == val {
			ii = append(ii, i)
		}
	}
	return ii[len(ii)/2]
}

func split(report string, groups []int) int {
	if len(groups) > 0 && len(report) == 0 {
		return 0
	}
	if len(groups) == 0 {
		if len(report) == 0 {
			return 1
		} else if !strings.Contains(report, "#") {
			return 1
		} else {
			return 0
		}
	}

	largestGroup := sliceMax(groups)
	groupIndex := medianIndex(groups, largestGroup)

	pp := groupValidPos(report, largestGroup)
	var branches int
	// should at least have one side group...
	for _, p := range pp {
		c := 1
		// left
		lC := split(leftString(report, p-gap), leftIntSlice(groups, groupIndex))
		if lC == 0 {
			continue
		}
		c *= lC

		// right
		rC := split(rightString(report, p+largestGroup+gap), rightIntSlice(groups, groupIndex+1))
		if rC == 0 {
			continue
		}
		c *= rC
		branches += c
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
