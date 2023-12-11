package day10

import (
	_ "embed"
	"fmt"
	"ruan.moolman/aoc2023/set"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

type direction string

const (
	UP    direction = "UP"
	DOWN  direction = "DOWN"
	LEFT  direction = "LEFT"
	RIGHT direction = "RIGHT"
)

type object string

func (o object) String() string {
	switch o {
	case Start:
		return "*"
	case PipeUpLeft:
		return "┘"
	case PipeUpRight:
		return "└"
	case PipeDownLeft:
		return "┐"
	case PipeDownRight:
		return "┌"
	case PipeLeftRight:
		return "─"
	case PipeUpDown:
		return "│"
	default:
		return "."
	}
}

func (p object) pointsTo(dir direction) bool {
	switch dir {
	case DOWN:
		return p == PipeDownRight || p == PipeDownLeft || p == PipeUpDown
	case UP:
		return p == PipeUpRight || p == PipeUpLeft || p == PipeUpDown
	case LEFT:
		return p == PipeUpLeft || p == PipeDownLeft || p == PipeLeftRight
	case RIGHT:
		return p == PipeUpRight || p == PipeDownRight || p == PipeLeftRight
	}
	return false
}

const (
	Empty         object = "."
	Start         object = "S"
	PipeUpLeft    object = "J"
	PipeUpRight   object = "L"
	PipeDownLeft  object = "7"
	PipeDownRight object = "F"
	PipeLeftRight object = "-"
	PipeUpDown    object = "|"
)

type position struct {
	x, y int
}

func (p position) adjacent(d direction) position {
	switch d {
	case RIGHT:
		return p.right()
	case DOWN:
		return p.down()
	case LEFT:
		return p.left()
	case UP:
		return p.up()
	}
	panic("Invalid direction")
}

func (p position) right() position {
	return position{p.x + 1, p.y}
}
func (p position) left() position {
	return position{p.x - 1, p.y}
}
func (p position) up() position {
	return position{p.x, p.y - 1}
}
func (p position) down() position {
	return position{p.x, p.y + 1}
}

type pipeMap struct {
	data []string
}

func NewPipeMap(data string) pipeMap {
	return pipeMap{strings.Split(data, "\n")}
}

func (m pipeMap) String() string {
	var lines []string
	for _, r := range m.data {
		var line string
		for _, p := range r {
			line += object(p).String()
		}
		lines = append(lines, line)
	}
	return strings.Join(lines, "\n")
}

func (m pipeMap) StringMarkPositions(ii set.Set[position], value string) string {
	var lines []string
	for y, r := range m.data {
		var line string
		for x, p := range r {
			if ii.Contains(position{x, y}) {
				line += value
				continue
			}
			line += object(p).String()
		}
		lines = append(lines, line)
	}
	return strings.Join(lines, "\n")
}

func (m pipeMap) insidePipeLoop() set.Set[position] {
	inside := set.NewSet[position]()
	for y, row := range m.data {
		isInside := false
		startPipe := Empty
		for x, p := range row {
			obj := object(p)
			switch obj {
			case PipeUpDown:
				isInside = !isInside
			case PipeDownRight, PipeUpRight:
				// pipes that run horizontally start with a part going right
				startPipe = obj
			case PipeDownLeft, PipeUpLeft:
				// pipes that run horizontally end with a part going left

				// Only switch sides when the starting pipe is opposite to
				// the end one wrt up down pipes
				if startPipe.pointsTo(UP) && obj.pointsTo(DOWN) ||
					startPipe.pointsTo(DOWN) && obj.pointsTo(UP) {
					isInside = !isInside
				}
				startPipe = Empty
			case Empty:
				if isInside {
					inside.Add(position{x, y})
				}
			}
		}
	}
	return inside
}

// clean return a pipeMap where all pipes that don't belong to the
// main loop are replaces by empty spots
func (m pipeMap) clean() pipeMap {
	pp := m.pipePositions()
	var newData []string
	for y, r := range m.data {
		var line string
		for x := range r {
			p := position{x, y}
			if !pp.Contains(p) {
				line += string(Empty)
			} else {
				posPipe := m.at(p)
				if posPipe == Start {
					posPipe = m.startPipe()
				}
				line += string(posPipe)
			}
		}
		newData = append(newData, line)
	}
	return pipeMap{newData}
}

// startPos returns the position of the start object
func (m pipeMap) startPos() position {
	for ri, r := range m.data {
		for ci, p := range r {
			if object(p) == Start {
				return position{ci, ri}
			}
		}
	}
	panic("No start found")
}

// startPipe returns the pipe object type that is covered by the start object
func (m pipeMap) startPipe() object {
	p := m.startPos()
	up := m.atSafe(p.up()).pointsTo(DOWN)
	down := m.atSafe(p.down()).pointsTo(UP)
	left := m.atSafe(p.left()).pointsTo(RIGHT)
	right := m.atSafe(p.right()).pointsTo(LEFT)

	switch {
	case up && down:
		return PipeUpDown
	case left && right:
		return PipeLeftRight
	case up && left:
		return PipeUpLeft
	case up && right:
		return PipeUpRight
	case down && left:
		return PipeDownLeft
	case down && right:
		return PipeDownRight
	}
	return Empty
}

// leadsTo returns the positions a pipe at the given position `p` leads to
func (m pipeMap) leadsTo(p position) []position {
	var pp []position
	pipe := m.at(p)
	if pipe == Start {
		pipe = m.startPipe()
	}
	switch pipe {
	case PipeUpDown:
		pp = append(pp, p.up(), p.down())
	case PipeLeftRight:
		pp = append(pp, p.left(), p.right())
	case PipeUpLeft:
		pp = append(pp, p.up(), p.left())
	case PipeUpRight:
		pp = append(pp, p.up(), p.right())
	case PipeDownLeft:
		pp = append(pp, p.down(), p.left())
	case PipeDownRight:
		pp = append(pp, p.down(), p.right())
	}
	return pp
}

func (m pipeMap) pipePositions() set.Set[position] {
	startPos := m.startPos()
	pipePositions := set.NewSet[position]()
	pipePositions.Add(startPos)

	activePos := startPos
mainLoop:
	for {
		for _, p := range m.leadsTo(activePos) {
			if !pipePositions.Contains(p) {
				activePos = p
				pipePositions.Add(p)
				continue mainLoop
			}
		}
		break
	}
	return pipePositions
}

func (m pipeMap) atSafe(p position) object {
	if p.y < 0 || p.x < 0 || p.y > len(m.data) || p.x > len(m.data[0]) {
		return Empty
	}
	return m.at(p)
}

func (m pipeMap) at(p position) object {
	return object(m.data[p.y][p.x])
}

func Part1() string {
	data := input
	m := NewPipeMap(data)
	fmt.Println(m)
	return strconv.Itoa(m.pipePositions().Count() / 2)
}

func Part2() string {
	data := input
	m := NewPipeMap(data).clean()
	insidePos := m.insidePipeLoop()
	fmt.Println(m.StringMarkPositions(insidePos, "I"))
	return strconv.Itoa(insidePos.Count())
}
