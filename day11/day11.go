package day11

import (
	_ "embed"
	intmath "github.com/thomaso-mirodin/intmath/intgr"
	"ruan.moolman/aoc2023/set"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

const GALAXY string = "#"

type position struct {
	x, y int
}

func spaceDist(a, b int, s set.Set[int], factor int64) int64 {
	minV := intmath.Min(a, b)
	maxV := intmath.Max(a, b)
	var c int64
	for i := minV + 1; i < maxV; i++ {
		if s.Contains(i) {
			c++
		}
	}
	return int64(maxV-minV) + c*(factor-1)
}

func (p position) SpaceDistance(o position, cols, rows set.Set[int], factor int64) int64 {
	return spaceDist(o.x, p.x, cols, factor) + spaceDist(o.y, p.y, rows, factor)
}

func emptySpace(space string) (set.Set[int], set.Set[int]) {
	rows := set.NewSet[int]()
	colsWithGalaxy := set.NewSet[int]()
	lines := strings.Split(space, "\n")
	for y, line := range lines {
		rowHasGalaxy := false
		for x, val := range line {
			isGalaxy := string(val) == GALAXY
			if isGalaxy {
				rowHasGalaxy = true
				colsWithGalaxy.Add(x)
			}
		}
		if !rowHasGalaxy {
			rows.Add(y)
		}
	}
	cols := set.NewSet[int]()
	for x := range lines[0] {
		if !colsWithGalaxy.Contains(x) {
			cols.Add(x)
		}
	}

	return cols, rows
}

func galaxyPositions(space string) []position {
	var pp []position
	for y, line := range strings.Split(space, "\n") {
		for x, val := range line {
			if string(val) == GALAXY {
				pp = append(pp, position{x, y})
			}
		}
	}
	return pp
}

func distanceBetweenGalaxies(space string, factor int64) int64 {
	colSet, rowSet := emptySpace(space)
	pp := galaxyPositions(space)
	var ans int64
	for i := 0; i < len(pp)-1; i++ {
		for j := i + 1; j < len(pp); j++ {
			ans += pp[i].SpaceDistance(pp[j], colSet, rowSet, factor)
		}
	}
	return ans
}

func Part1() string {
	return strconv.FormatInt(distanceBetweenGalaxies(input, 2), 10)
}

func Part2() string {
	return strconv.FormatInt(distanceBetweenGalaxies(input, 1000000), 10)
}
