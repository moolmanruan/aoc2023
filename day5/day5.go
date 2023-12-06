package day5

import (
	"cmp"
	"fmt"
	intmath "github.com/thomaso-mirodin/intmath/intgr"
	"regexp"
	"slices"
	"strconv"
	"strings"
)
import _ "embed"

//go:embed example.txt
var example string

//go:embed input.txt
var input string

type MappingRange struct {
	source      int
	destination int
	count       int
}

func (r MappingRange) String() string {
	return fmt.Sprintf("[%d, %d)->[%d, %d)", r.source, r.sourceEnd(), r.destination, r.destEnd())
}

func (r MappingRange) diff() int {
	return r.destination - r.source
}

func (r MappingRange) sourceRangeContains(value int) bool {
	return value >= r.source && value < r.sourceEnd()
}

func (r MappingRange) sourceEnd() int {
	return r.source + r.count
}

func (r MappingRange) destEnd() int {
	return r.destination + r.count
}

type Mapping struct {
	from   string
	to     string
	ranges []MappingRange
}

var mapStartRegex, _ = regexp.Compile(`([a-z]+)-to-([a-z]+) map:`)

func mustAtoi(in string) int {
	val, err := strconv.Atoi(in)
	if err != nil {
		panic(err)
	}
	return val
}

func parseRange(input string) MappingRange {
	parts := strings.Split(input, " ")
	return MappingRange{
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
		slices.SortFunc(m.ranges, func(a, b MappingRange) int {
			return cmp.Compare(a.source, b.source)
		})
		mm = append(mm, m)
	}
	return mm
}

type seedMap map[string]int

func mapID(id int, ranges []MappingRange) int {
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

type IDRange struct {
	start int
	count int
}

func (r IDRange) String() string {
	return fmt.Sprintf("[%d, %d)", r.start, r.end())
}

func (r IDRange) end() int {
	return r.start + r.count
}

func parseSeedRanges(input string) []IDRange {
	line := strings.Split(input, "\n")[0]
	vvStr := strings.Split(strings.Trim(strings.Split(line, ":")[1], " "), " ")
	var srs []IDRange
	for i := 0; i < len(vvStr); i += 2 {
		srs = append(srs, IDRange{
			start: mustAtoi(vvStr[i]),
			count: mustAtoi(vvStr[i+1]),
		})
	}

	slices.SortFunc(srs, func(a, b IDRange) int {
		return cmp.Compare(a.start, b.start)
	})
	return srs
}

func rangeForIndex(mapping Mapping, index int) (MappingRange, bool) {
	for _, r := range mapping.ranges {
		if r.sourceRangeContains(index) {
			return r, true
		}
	}
	return MappingRange{}, false
}
func nextRangeAfterIndex(mapping Mapping, index int) (MappingRange, bool) {
	for _, r := range mapping.ranges {
		if r.source >= index {
			return r, true
		}
	}
	return MappingRange{}, false
}

func applyMappingToRange(idRange IDRange, mapping Mapping) []IDRange {
	var rr []IDRange
	i := idRange.start
	for i < idRange.end() {
		if r, ok := rangeForIndex(mapping, i); ok {
			d := r.diff()
			count := intmath.Min(r.sourceEnd(), idRange.end()) - i
			rr = append(rr, IDRange{start: i + d, count: count})
			i += count
		} else {
			if nextR, ok := nextRangeAfterIndex(mapping, i); ok {
				count := intmath.Min(nextR.source, idRange.end()) - i
				rr = append(rr, IDRange{start: i, count: count})
				i += count
			} else {
				newRange := IDRange{start: i, count: idRange.end() - i}
				rr = append(rr, newRange)
				i += newRange.count
			}
		}
	}
	return rr
}

func applyMappingsToRange(idRange IDRange, mappings []Mapping) []IDRange {
	rr := []IDRange{idRange}
	for _, m := range mappings {
		var rrNew []IDRange
		for _, r := range rr {
			rrNew = append(rrNew, applyMappingToRange(r, m)...)
		}
		rr = rrNew
	}
	return rr
}

func mapIDRanges(rr []IDRange, mappings []Mapping) []IDRange {
	var mm []IDRange
	for _, s := range rr {
		mm = append(mm, applyMappingsToRange(s, mappings)...)
	}
	return mm
}

func Execute() {
	inputString := input
	seeds := parseSeedRanges(inputString)
	mm := parseMappings(inputString)
	rr := mapIDRanges(seeds, mm)
	ans := -1
	for _, s := range rr {
		if s.start < ans || ans == -1 {
			ans = s.start
		}
	}
	fmt.Printf("answer %d\n", ans)
}

func parseSeeds(input string) []int {
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
	seeds := parseSeeds(inputString)
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
