package day11

import (
	_ "embed"
	intmath "github.com/thomaso-mirodin/intmath/intgr"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

const GALAXY string = "#"

type position struct {
	x, y int
}

func (p position) L1DistanceThroughSpace(o position, cols, rows map[int]struct{}, factor int64) int64 {
	minX := intmath.Min(o.x, p.x)
	maxX := intmath.Max(o.x, p.x)
	minY := intmath.Min(o.y, p.y)
	maxY := intmath.Max(o.y, p.y)
	var emptyCols, emptyRows int64
	for x := minX + 1; x < maxX; x++ {
		if _, ok := cols[x]; ok {
			emptyCols++
		}
	}
	for y := minY + 1; y < maxY; y++ {
		if _, ok := rows[y]; ok {
			emptyRows++
		}
	}
	return int64(intmath.Abs(o.x-p.x)) + emptyCols*(factor-1) + int64(intmath.Abs(o.y-p.y)) + emptyRows*(factor-1)
}

func emptySpace(space string) (map[int]struct{}, map[int]struct{}) {
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

	return colsWithoutGalaxy, rowsWithoutGalaxy
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

func Part1() string {
	space := input
	colSet, rowSet := emptySpace(space)
	pp := galaxyPositions(space)
	var ans int64
	for i := 0; i < len(pp)-1; i++ {
		for j := i + 1; j < len(pp); j++ {
			ans += pp[i].L1DistanceThroughSpace(pp[j], colSet, rowSet, 2)
		}
	}
	return strconv.FormatInt(ans, 10)
}

func Part2() string {
	space := input
	colSet, rowSet := emptySpace(space)
	pp := galaxyPositions(space)
	var ans int64
	for i := 0; i < len(pp)-1; i++ {
		for j := i + 1; j < len(pp); j++ {
			ans += pp[i].L1DistanceThroughSpace(pp[j], colSet, rowSet, 1000000)
		}
	}
	return strconv.FormatInt(ans, 10)
}
