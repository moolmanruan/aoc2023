package day10

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed example1.txt
var example1 string

//go:embed example2.txt
var example2 string

//go:embed input.txt
var input string

type pipeMap struct {
	data []string
}

type dir string

const (
	UP    dir = "UP"
	DOWN  dir = "DOWN"
	LEFT  dir = "LEFT"
	RIGHT dir = "RIGHT"
)

type pipe string

func (p pipe) pointsDown() bool {
	return p == DR || p == DL || p == UD
}
func (p pipe) pointsUp() bool {
	return p == UR || p == UL || p == UD
}
func (p pipe) pointsLeft() bool {
	return p == UL || p == DL || p == LR
}
func (p pipe) pointsRight() bool {
	return p == UR || p == DR || p == LR
}

const (
	NONE  pipe = "."
	START pipe = "S"
	UL    pipe = "J"
	UR    pipe = "L"
	DL    pipe = "7"
	DR    pipe = "F"
	LR    pipe = "-"
	UD    pipe = "|"
)

type pos struct {
	x, y int
}

func (p pos) adjacent(d dir) pos {
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

func (p pos) right() pos {
	return pos{p.x + 1, p.y}
}
func (p pos) left() pos {
	return pos{p.x - 1, p.y}
}
func (p pos) up() pos {
	return pos{p.x, p.y - 1}
}
func (p pos) down() pos {
	return pos{p.x, p.y + 1}
}

func (p pos) adjacentPositions() []pos {
	return []pos{
		p.right(),
		p.down(),
		p.left(),
		p.up(),
	}
}

func NewPipeMap(data string) pipeMap {
	return pipeMap{strings.Split(data, "\n")}
}

func (m pipeMap) startPos() pos {
	for ri, r := range m.data {
		for ci, p := range r {
			if pipe(p) == START {
				return pos{ci, ri}
			}
		}
	}
	panic("No start found")
}

func (m pipeMap) connectedPipes(p pos) []pos {
	var pp []pos
	switch m.at(p) {
	case START:
		if m.atSafe(p.up()).pointsDown() {
			pp = append(pp, p.up())
		}
		if m.atSafe(p.down()).pointsUp() {
			pp = append(pp, p.down())
		}
		if m.atSafe(p.left()).pointsRight() {
			pp = append(pp, p.left())
		}
		if m.atSafe(p.right()).pointsLeft() {
			pp = append(pp, p.right())
		}
	case UD:
		pp = append(pp, p.up(), p.down())
	case LR:
		pp = append(pp, p.left(), p.right())
	case UL:
		pp = append(pp, p.up(), p.left())
	case UR:
		pp = append(pp, p.up(), p.right())
	case DL:
		pp = append(pp, p.down(), p.left())
	case DR:
		pp = append(pp, p.down(), p.right())
	}
	return pp
}

func (m pipeMap) loopLength() int {
	sp := m.startPos()
	pipePositions := make(map[pos]struct{})
	pipePositions[sp] = struct{}{}

	activePos := sp
mainLoop:
	for {
		for _, p := range m.connectedPipes(activePos) {
			if _, ok := pipePositions[p]; !ok {
				activePos = p
				pipePositions[p] = struct{}{}
				continue mainLoop
			}
		}
		break
	}
	return len(pipePositions)
}

func (m pipeMap) atSafe(p pos) pipe {
	if p.y < 0 || p.x < 0 || p.y > len(m.data) || p.x > len(m.data[0]) {
		return NONE
	}
	return m.at(p)
}

func (m pipeMap) at(p pos) pipe {
	return pipe(m.data[p.y][p.x])
}

func prettify(data string) string {
	s := strings.ReplaceAll(data, string(UL), "┘")
	s = strings.ReplaceAll(s, string(UR), "└")
	s = strings.ReplaceAll(s, string(DL), "┐")
	s = strings.ReplaceAll(s, string(DR), "┌")
	s = strings.ReplaceAll(s, string(LR), "─")
	s = strings.ReplaceAll(s, string(UD), "│")
	return s
}

func Part1() string {
	data := input
	fmt.Println(prettify(data))
	m := NewPipeMap(data)
	return strconv.Itoa(m.loopLength() / 2)
}
