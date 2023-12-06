package day6

import "fmt"

type Race struct {
	time     int
	distance int
}

var example = []Race{
	{7, 9},
	{15, 40},
	{30, 200},
}
var input = []Race{
	{47, 207},
	{84, 1394},
	{74, 1209},
	{67, 1014},
}

func distance(holdTime int, totalTime int) int {
	return (totalTime - holdTime) * holdTime
}

func calculateMargin(rr []Race) int {
	var aa []int
	for _, r := range rr {
		rc := 0
		for i := 0; i < r.time; i++ {
			if distance(i, r.time) > r.distance {
				rc++
			}
		}

		aa = append(aa, rc)
	}
	ans := aa[0]
	for _, v := range aa[1:] {
		ans *= v
	}
	return ans
}

func Execute() {
	fmt.Printf("answer %d\n", calculateMargin(input))
}
