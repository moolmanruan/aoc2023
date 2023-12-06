package day6

import (
	"strconv"
)

type Race struct {
	time     int64
	distance int64
}

var example = []Race{
	{7, 9},
	{15, 40},
	{30, 200},
}
var example2 = Race{
	71530,
	940200,
}
var input = []Race{
	{47, 207},
	{84, 1394},
	{74, 1209},
	{67, 1014},
}
var input2 = Race{
	47847467,
	207139412091014,
}

func distance(holdTime int64, totalTime int64) int64 {
	return (totalTime - holdTime) * holdTime
}

func calculateRace(r Race) int {
	rc := 0
	for i := int64(0); i < r.time; i++ {
		if distance(i, r.time) > r.distance {
			rc++
		}
	}
	return rc
}

func calculateMargin(rr []Race) int {
	var aa []int
	for _, r := range rr {
		aa = append(aa, calculateRace(r))
	}
	ans := aa[0]
	for _, v := range aa[1:] {
		ans *= v
	}
	return ans
}

func Part2() string {
	return strconv.Itoa(calculateRace(input2))
}

func Part1() string {
	return strconv.Itoa(calculateMargin(input))
}
