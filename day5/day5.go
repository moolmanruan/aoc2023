package day5

import (
	"cmp"
	"fmt"
	intmath "github.com/thomaso-mirodin/intmath/intgr"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"time"
)
import _ "embed"

//go:embed example.txt
var example string

//go:embed input.txt
var input string

type Range struct {
	start int
	count int
}

func (r Range) String() string {
	return fmt.Sprintf("[%d, %d)", r.start, r.end())
}

func (r Range) end() int {
	return r.start + r.count
}

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

func parseRanges(input string) []Range {
	line := strings.Split(input, "\n")[0]
	vvStr := strings.Split(strings.Trim(strings.Split(line, ":")[1], " "), " ")
	var srs []Range
	for i := 0; i < len(vvStr); i += 2 {
		srs = append(srs, Range{
			start: mustAtoi(vvStr[i]),
			count: mustAtoi(vvStr[i+1]),
		})
	}

	slices.SortFunc(srs, func(a, b Range) int {
		return cmp.Compare(a.start, b.start)
	})
	return srs
}

// Return the range that contains the index value.
// The second return value indicates if a range could be found or not.
func rangeForIndex(mapping Mapping, index int) (MappingRange, bool) {
	for _, r := range mapping.ranges {
		if r.sourceRangeContains(index) {
			return r, true
		}
	}
	return MappingRange{}, false
}

// Return the first range comes after the given index value.
// The second return value indicates if a range could be found or not.
func nextRangeAfterIndex(mapping Mapping, index int) (MappingRange, bool) {
	for _, r := range mapping.ranges {
		if r.source >= index {
			return r, true
		}
	}
	return MappingRange{}, false
}

// Convert the source `inputRange` to a list of destination ranges using the `mapping` provided
func applyMappingToRange(inputRange Range, mapping Mapping) []Range {
	var outputRanges []Range
	i := inputRange.start
	for i < inputRange.end() {
		if r, ok := rangeForIndex(mapping, i); ok {
			d := r.diff()
			count := intmath.Min(r.sourceEnd(), inputRange.end()) - i
			outputRanges = append(outputRanges, Range{start: i + d, count: count})
			i += count
		} else {
			if nextR, ok := nextRangeAfterIndex(mapping, i); ok {
				count := intmath.Min(nextR.source, inputRange.end()) - i
				outputRanges = append(outputRanges, Range{start: i, count: count})
				i += count
			} else {
				newRange := Range{start: i, count: inputRange.end() - i}
				outputRanges = append(outputRanges, newRange)
				i += newRange.count
			}
		}
	}
	return outputRanges
}

// Apply a list of `mappings` to the source `inputRange` range to get a list of destination output ranges.
// The first mapping is applied to the input range, while the following mappings are applied to result of the resulting
// ranges of the previous mapping.
func applyMappingsToRange(inputRange Range, mappings []Mapping) []Range {
	rr := []Range{inputRange}
	for _, m := range mappings {
		var rrNew []Range
		for _, r := range rr {
			rrNew = append(rrNew, applyMappingToRange(r, m)...)
		}
		rr = rrNew
	}
	return rr
}

// Apply multiple mappings to multiple input ranges to get the destination ranges.
func applyMappingsToRanges(rr []Range, mappings []Mapping) []Range {
	var mm []Range
	for _, s := range rr {
		mm = append(mm, applyMappingsToRange(s, mappings)...)
	}
	return mm
}

func Execute() {
	t := time.Now()
	inputString := input
	srcRanges := parseRanges(inputString)
	mm := parseMappings(inputString)
	destRanges := applyMappingsToRanges(srcRanges, mm)
	ans := -1
	for _, s := range destRanges {
		if s.start < ans || ans == -1 {
			ans = s.start
		}
	}

	fmt.Printf("Answer: %d\n(%d μs)\n", ans, time.Now().Sub(t).Microseconds())
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
	var idRanges []Range
	for _, s := range parseSeeds(inputString) {
		idRanges = append(idRanges, Range{start: s, count: 1})
	}
	mm := parseMappings(inputString)
	ss := applyMappingsToRanges(idRanges, mm)
	ans := -1
	for _, s := range ss {
		if s.start < ans || ans == -1 {
			ans = s.start
		}
	}
	fmt.Printf("Answer %d\n", ans)
}
