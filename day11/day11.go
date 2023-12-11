package day11

import (
	_ "embed"
	"fmt"
	intmath "github.com/thomaso-mirodin/intmath/intgr"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

const GALAXY string = "#"
const EMPTY string = "."

type position struct {
	x, y int
}

func (p position) L1Distance(o position) int {
	return intmath.Abs(o.x-p.x) + intmath.Abs(o.y-p.y)
}

func expand(space string) string {
	rowsWithoutGalaxy := make(map[int]struct{})
	colsWithGalaxy := make(map[int]struct{})
	lines := strings.Split(space, "\n")
	for y, line := range lines {
		rowHasGalaxy := false
		for x, val := range line {
			isGalaxy := string(val) == GALAXY
			if isGalaxy {
				rowHasGalaxy = true
				colsWithGalaxy[x] = struct{}{}
			}
		}
		if !rowHasGalaxy {
			rowsWithoutGalaxy[y] = struct{}{}
		}
	}
	colsWithoutGalaxy := make(map[int]struct{})
	for x := range lines[0] {
		if _, ok := colsWithGalaxy[x]; !ok {
			colsWithoutGalaxy[x] = struct{}{}
		}
	}

	var newLines []string
	for y, line := range lines {
		var newLine string
		for x, v := range line {
			newLine += string(v)
			if _, ok := colsWithoutGalaxy[x]; ok {
				newLine += string(v)
			}
		}
		newLines = append(newLines, newLine)
		if _, ok := rowsWithoutGalaxy[y]; ok {
			newLines = append(newLines, newLine)
		}
	}

	return strings.Join(newLines, "\n")
}

func galaxyPositions(space string) []position {
	var pp []position
	for x, line := range strings.Split(space, "\n") {
		for y, val := range line {
			if string(val) == GALAXY {
				pp = append(pp, position{x, y})
			}
		}
	}
	return pp
}

func Part1() string {
	space := expand(input)
	fmt.Println(space)
	pp := galaxyPositions(space)
	var ans int
	for i := 0; i < len(pp)-1; i++ {
		for j := i + 1; j < len(pp); j++ {
			ans += pp[i].L1Distance(pp[j])
		}
	}
	return strconv.Itoa(ans)
}

func Part2() string {
	return "-"
}
