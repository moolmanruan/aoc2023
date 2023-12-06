package day5

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)
import _ "embed"

//go:embed example.txt
var example string

//go:embed input.txt
var input string

type Range struct {
	destination int
	source      int
	count       int
}

type Mapping struct {
	from   string
	to     string
	ranges []Range
}

var mapStartRegex, _ = regexp.Compile(`([a-z]+)-to-([a-z]+) map:`)

func mustAtoi(in string) int {
	val, err := strconv.Atoi(in)
	if err != nil {
		panic(err)
	}
	return val
}

func parseRange(input string) Range {
	parts := strings.Split(input, " ")
	return Range{
		destination: mustAtoi(parts[0]),
		source:      mustAtoi(parts[1]),
		count:       mustAtoi(parts[2]),
	}
}

func parseMappings(input string) []Mapping {
	lines := strings.Split(input, "\n")
	var mm []Mapping
	for li := 0; li < len(lines); li++ {
		line := lines[li]
		res := mapStartRegex.FindStringSubmatch(line)
		if len(res) == 0 {
			continue
		}

		m := Mapping{
			from: res[1],
			to:   res[2],
		}
		for {
			li++ // next line
			if li >= len(lines) {
				break
			}
			line = lines[li]
			if line == "" {
				break
			}
			m.ranges = append(m.ranges, parseRange(line))
		}
		mm = append(mm, m)
	}
	return mm
}

type seedMap map[string]int

func mapID(id int, ranges []Range) int {
	for _, r := range ranges {
		if id >= r.source && id < r.source+r.count {
			return r.destination + (id - r.source)
		}
	}
	return id
}

func mapSeed(seed int, mappings []Mapping) seedMap {
	sm := seedMap{"seed": seed}
	for _, m := range mappings {
		fromID := sm[m.from]
		sm[m.to] = mapID(fromID, m.ranges)
	}
	return sm
}

func mapSeeds(seeds []int, mappings []Mapping) []seedMap {
	var mm []seedMap
	for _, s := range seeds {
		mm = append(mm, mapSeed(s, mappings))
	}
	return mm
}

func parseSeedsPart2(input string) []int {
	line := strings.Split(input, "\n")[0]
	vvStr := strings.Split(strings.Trim(strings.Split(line, ":")[1], " "), " ")
	var vv []int
	for i := 0; i < len(vvStr); i += 2 {
		from := mustAtoi(vvStr[i])
		count := mustAtoi(vvStr[i+1])
		for s := from; s < from+count; s++ {
			vv = append(vv, s)
		}
	}
	return vv
}

func Execute() {
	inputString := input
	seeds := parseSeedsPart2(inputString)
	mm := parseMappings(inputString)
	ss := mapSeeds(seeds, mm)
	ans := -1
	for _, s := range ss {
		locationID := s["location"]
		if locationID < ans || ans == -1 {
			ans = locationID
		}
	}
	fmt.Printf("answer %d\n", ans)
}

func parseSeedsPart1(input string) []int {
	line := strings.Split(input, "\n")[0]
	valuesString := strings.Trim(strings.Split(line, ":")[1], " ")
	var vv []int
	for _, v := range strings.Split(valuesString, " ") {
		vv = append(vv, mustAtoi(v))
	}
	return vv
}

func ExecutePart1() {
	inputString := input
	seeds := parseSeedsPart1(inputString)
	mm := parseMappings(inputString)
	ss := mapSeeds(seeds, mm)
	ans := -1
	for _, s := range ss {
		locationID := s["location"]
		if locationID < ans || ans == -1 {
			ans = locationID
		}
	}
	fmt.Printf("answer %d\n", ans)
}
