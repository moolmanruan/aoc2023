package day2

import (
	"fmt"
	"strconv"
	"strings"
)

const inputExample = "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green\nGame 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue\nGame 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red\nGame 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red\nGame 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green"
const input = "Game 1: 10 red, 7 green, 3 blue; 5 blue, 3 red, 10 green; 4 blue, 14 green, 7 red; 1 red, 11 green; 6 blue, 17 green, 15 red; 18 green, 7 red, 5 blue\nGame 2: 13 green, 10 red; 11 green, 1 blue, 7 red; 5 red, 12 green, 1 blue; 12 green, 6 red; 8 green, 5 red; 12 green, 1 red\nGame 3: 7 green, 1 blue; 1 blue, 3 green, 1 red; 1 green, 1 blue; 2 green; 1 blue, 7 green, 2 red; 2 green\nGame 4: 7 green, 11 blue; 12 blue, 7 green; 1 green, 7 blue; 5 blue, 2 green; 5 red, 9 green, 14 blue\nGame 5: 2 red, 6 blue, 6 green; 2 red, 6 green; 12 blue, 5 red, 3 green; 12 green, 5 red, 8 blue; 10 blue, 5 green; 2 red, 4 green\nGame 6: 8 blue, 1 red, 17 green; 7 blue; 10 green, 6 blue; 5 blue, 1 red, 11 green\nGame 7: 1 blue, 2 red, 2 green; 1 blue, 3 green; 3 green, 1 red, 3 blue; 2 blue, 3 green, 1 red\nGame 8: 3 green, 10 red, 15 blue; 1 green, 9 red; 9 blue, 2 green, 12 red\nGame 9: 4 green, 10 blue, 13 red; 16 red, 7 blue; 14 red, 1 green, 1 blue; 14 red, 4 blue, 1 green\nGame 10: 6 blue, 9 red, 3 green; 9 green, 7 blue, 9 red; 2 red, 4 blue, 6 green; 12 green, 7 blue, 5 red\nGame 11: 1 green, 6 blue, 6 red; 7 red, 1 blue; 1 green, 6 blue; 4 red, 1 green, 1 blue; 6 red, 9 green, 4 blue; 5 green, 7 red, 4 blue\nGame 12: 18 green, 4 red, 12 blue; 7 green, 5 blue, 3 red; 7 green, 3 red; 8 green, 7 blue; 4 red, 7 green, 10 blue\nGame 13: 1 red, 2 blue; 1 red, 6 green; 5 blue, 2 red, 12 green; 1 red, 11 green, 2 blue; 2 red, 8 green, 1 blue; 3 blue, 16 green, 1 red\nGame 14: 3 blue, 2 green; 4 green, 1 red; 1 green, 1 red, 3 blue; 4 blue, 3 green; 5 blue, 1 green; 4 green, 2 blue, 1 red\nGame 15: 12 blue, 3 red; 5 blue, 2 red, 1 green; 12 blue, 3 red, 2 green; 1 green, 5 red, 6 blue; 1 green, 5 blue, 3 red\nGame 16: 8 red, 4 blue, 1 green; 15 blue, 5 red, 4 green; 3 green, 13 blue, 4 red; 4 red, 1 blue, 3 green; 1 green, 13 blue, 6 red\nGame 17: 8 red, 7 green, 2 blue; 6 green, 1 blue, 12 red; 3 red; 4 green, 1 red; 7 red, 1 blue, 9 green\nGame 18: 7 blue, 10 red, 3 green; 3 green, 1 blue; 7 red, 1 green, 7 blue; 7 blue, 4 red, 1 green; 2 green, 1 blue, 10 red; 3 blue, 11 red, 1 green\nGame 19: 10 red, 10 blue; 13 red; 4 blue, 15 red, 3 green; 6 green, 11 red, 11 blue; 4 blue, 8 red\nGame 20: 1 blue, 9 green, 2 red; 2 blue, 4 red, 4 green; 4 green, 2 red\nGame 21: 13 green, 1 red; 3 red, 5 green, 11 blue; 1 blue, 2 red, 4 green; 7 blue, 3 red; 2 red, 1 blue, 3 green\nGame 22: 2 red, 2 blue, 3 green; 10 red, 4 blue; 8 blue, 8 green, 11 red\nGame 23: 1 red, 2 blue; 1 blue, 1 green; 1 green; 3 red, 1 blue, 1 green\nGame 24: 12 green, 4 red, 2 blue; 8 green, 5 blue; 8 green, 2 blue, 2 red\nGame 25: 3 red, 8 green; 1 red, 4 blue, 1 green; 6 green; 3 blue, 5 green, 3 red; 9 green, 3 blue, 5 red\nGame 26: 1 green, 3 red, 2 blue; 7 red, 2 green, 11 blue; 7 blue, 4 red; 11 blue, 1 red, 1 green; 2 green, 10 blue, 1 red; 1 green, 7 red, 7 blue\nGame 27: 5 green, 2 red, 4 blue; 5 red, 4 blue, 3 green; 5 green, 2 red, 7 blue; 7 red, 15 green, 5 blue\nGame 28: 1 green, 7 blue, 14 red; 7 green, 6 blue, 3 red; 7 blue, 4 red, 10 green; 9 red, 11 green, 5 blue\nGame 29: 4 red, 6 blue, 5 green; 12 red, 3 green, 1 blue; 6 blue, 11 red, 6 green; 2 green, 2 blue, 12 red\nGame 30: 13 green, 11 red, 11 blue; 7 green, 9 blue, 7 red; 11 red, 1 blue, 11 green\nGame 31: 14 green, 1 blue, 8 red; 1 green, 2 blue; 1 green, 1 red, 1 blue\nGame 32: 7 blue, 2 green; 12 blue, 7 green; 4 red, 14 blue, 2 green; 14 green, 4 blue\nGame 33: 5 blue, 12 red; 3 blue, 4 red, 1 green; 9 red, 2 blue; 11 red\nGame 34: 1 blue; 3 blue; 1 blue, 1 red; 5 red, 2 blue; 4 red, 1 blue, 1 green\nGame 35: 3 green, 2 blue, 1 red; 2 red, 8 green, 3 blue; 7 green, 2 red, 8 blue; 3 blue, 4 green\nGame 36: 10 green, 9 blue, 2 red; 3 green, 7 blue, 7 red; 14 green, 13 blue; 8 green, 8 red, 2 blue\nGame 37: 3 red, 1 blue, 14 green; 1 blue, 1 green; 5 red, 9 green; 1 red, 2 blue, 13 green; 11 red, 14 green, 2 blue\nGame 38: 4 green, 3 red, 6 blue; 18 red, 15 blue, 1 green; 17 blue, 6 green, 19 red; 18 red, 15 blue; 1 green, 12 blue, 18 red\nGame 39: 1 red; 10 blue, 6 red, 1 green; 1 green, 1 red, 9 blue; 17 red, 10 blue\nGame 40: 5 red, 3 green, 9 blue; 8 red, 4 blue; 2 green, 3 blue, 4 red; 3 blue, 4 red, 6 green; 4 blue, 5 red, 2 green; 4 blue\nGame 41: 6 green, 1 blue; 5 blue, 3 green, 6 red; 10 red, 1 blue; 6 green, 1 blue, 9 red\nGame 42: 1 red, 5 green, 7 blue; 7 red, 4 blue, 4 green; 5 red, 2 green, 6 blue\nGame 43: 1 green, 18 red, 8 blue; 7 red, 4 green, 5 blue; 1 blue, 18 red; 5 red, 8 blue\nGame 44: 3 blue, 10 green; 5 green, 2 red, 1 blue; 6 blue, 14 green; 3 green, 5 blue, 5 red\nGame 45: 12 red, 1 blue, 16 green; 1 red, 6 blue, 3 green; 5 red, 5 blue, 7 green; 8 red, 15 green; 3 green, 12 red, 7 blue\nGame 46: 3 red, 1 green; 1 green, 17 blue, 10 red; 2 green, 17 blue; 3 green, 17 blue, 12 red; 2 green, 12 red\nGame 47: 3 green, 9 red; 3 red, 1 blue, 6 green; 10 red, 9 green, 1 blue; 2 blue, 15 green; 7 red, 12 green, 3 blue\nGame 48: 4 green, 13 red, 14 blue; 8 red, 8 green; 15 blue, 4 red, 11 green; 3 blue, 3 red, 4 green; 2 blue, 6 red, 4 green; 13 green, 12 blue, 11 red\nGame 49: 15 blue, 2 green, 7 red; 1 green, 7 red, 7 blue; 13 blue; 3 blue, 2 red, 1 green\nGame 50: 9 red; 5 green, 2 blue, 10 red; 5 red, 1 green\nGame 51: 3 green, 1 blue, 3 red; 4 blue, 4 red; 4 green, 6 red, 5 blue; 4 red, 7 blue\nGame 52: 10 green, 12 red, 2 blue; 2 green, 7 red; 18 green, 3 red, 3 blue; 6 red, 13 green, 2 blue\nGame 53: 13 blue, 2 green; 2 green, 12 blue; 1 green, 11 blue, 1 red; 11 blue, 2 green, 8 red\nGame 54: 5 red; 15 green, 17 red, 7 blue; 14 green, 5 red, 15 blue; 2 red, 10 blue, 16 green\nGame 55: 1 blue, 1 red, 2 green; 5 green, 3 blue, 8 red; 6 red, 4 blue, 7 green; 2 blue, 10 green, 7 red\nGame 56: 1 blue, 8 red, 7 green; 3 green, 7 blue, 5 red; 5 green, 7 blue; 3 blue, 12 red, 8 green; 3 blue; 2 blue, 3 green, 10 red\nGame 57: 5 red, 13 green, 3 blue; 19 green, 7 red, 8 blue; 1 red, 12 green, 3 blue; 4 green, 10 blue, 4 red; 3 blue, 7 red, 20 green\nGame 58: 8 blue, 5 red, 2 green; 4 red, 11 blue; 9 blue, 6 green, 8 red; 7 green, 11 blue\nGame 59: 7 red, 7 green, 9 blue; 5 red, 4 green, 5 blue; 1 red, 2 blue, 6 green; 10 green, 12 blue, 3 red; 7 green, 18 blue, 4 red\nGame 60: 12 blue, 7 red, 12 green; 18 green, 9 red; 13 green, 13 red, 12 blue; 14 red, 5 green, 13 blue; 17 green, 7 red, 13 blue\nGame 61: 5 blue; 2 blue, 10 green, 2 red; 12 green, 2 red, 1 blue; 4 blue, 2 green; 2 red, 6 green; 6 green, 2 blue, 2 red\nGame 62: 2 blue, 5 red, 4 green; 3 green, 6 blue, 7 red; 13 red, 5 blue, 1 green; 3 red, 3 blue, 1 green; 17 blue, 4 green, 3 red; 5 red, 13 blue, 3 green\nGame 63: 1 red, 6 blue, 10 green; 1 red, 8 blue, 6 green; 7 red, 11 blue\nGame 64: 11 blue, 13 red; 12 blue, 6 red; 1 green, 2 blue, 4 red\nGame 65: 1 green, 9 red, 4 blue; 11 blue, 3 green; 2 blue, 1 green; 3 red, 2 green, 10 blue\nGame 66: 8 red, 1 blue, 3 green; 1 green, 3 blue, 1 red; 2 blue, 9 green; 8 green, 3 blue, 6 red; 2 blue, 12 green, 7 red\nGame 67: 5 green, 5 red, 10 blue; 12 blue, 13 green, 4 red; 6 red, 11 green, 3 blue; 8 blue, 4 red; 4 red, 14 green; 1 red, 1 blue, 14 green\nGame 68: 7 green, 17 red; 14 green, 1 blue, 1 red; 11 green, 1 blue, 16 red\nGame 69: 11 red, 2 green, 2 blue; 4 blue, 14 red; 2 red, 6 blue, 3 green; 6 red, 2 green; 5 red, 1 green, 4 blue; 7 red, 3 blue\nGame 70: 18 blue, 4 red; 5 red, 14 blue; 17 blue, 9 red; 13 red, 17 blue, 1 green; 2 blue, 9 red\nGame 71: 1 green, 6 red, 6 blue; 6 green, 4 blue, 5 red; 8 red, 3 blue, 7 green; 7 red, 2 blue, 1 green; 3 blue, 2 green, 3 red\nGame 72: 11 green, 4 red, 2 blue; 2 blue, 6 green, 1 red; 3 red, 1 blue, 9 green; 4 blue, 12 green, 3 red; 2 red, 3 green, 1 blue\nGame 73: 1 blue, 12 red; 14 green, 2 blue, 10 red; 6 blue, 8 red, 8 green; 7 green; 6 red, 10 green, 4 blue; 4 green, 9 red\nGame 74: 5 green, 6 blue; 1 green, 12 blue; 2 blue, 2 green, 5 red; 5 green, 9 blue, 2 red\nGame 75: 11 red, 7 blue, 12 green; 7 blue, 8 red, 9 green; 3 red, 17 green, 3 blue\nGame 76: 1 green, 12 blue; 11 blue, 7 green, 10 red; 10 green, 12 blue, 1 red; 10 green, 12 red, 1 blue\nGame 77: 2 blue, 17 green, 3 red; 10 red, 13 green; 12 green, 2 blue, 13 red; 12 green, 2 blue, 8 red; 14 green, 10 red, 1 blue\nGame 78: 3 red, 8 green, 5 blue; 8 green, 3 blue; 2 green, 6 red; 4 red, 1 green, 4 blue; 4 red, 8 green, 6 blue; 1 red, 1 blue, 8 green\nGame 79: 1 green, 2 blue, 2 red; 1 blue, 19 red, 1 green; 18 red; 1 green, 3 red, 5 blue; 15 red, 1 blue; 2 blue, 17 red, 1 green\nGame 80: 13 red, 1 green; 15 red, 1 blue; 8 red, 1 green\nGame 81: 1 blue, 1 red, 2 green; 1 red, 3 green, 2 blue; 1 blue, 4 green; 2 green, 2 blue\nGame 82: 8 red, 4 green, 8 blue; 4 green, 6 red, 3 blue; 3 red, 3 blue; 2 blue, 1 green, 11 red; 2 green, 1 blue, 4 red\nGame 83: 1 red, 15 green; 2 red, 6 blue, 12 green; 3 green, 10 blue, 14 red; 6 blue, 7 red, 1 green\nGame 84: 2 blue, 12 red, 4 green; 1 red, 3 blue, 5 green; 6 blue, 5 green, 12 red; 2 red, 1 green; 2 red, 5 blue, 5 green\nGame 85: 4 red; 3 red, 15 green, 2 blue; 15 green, 1 red, 2 blue; 4 green, 4 red, 2 blue\nGame 86: 1 green, 3 red, 4 blue; 2 green, 7 red, 4 blue; 7 red, 4 green, 4 blue; 1 blue, 11 red, 4 green\nGame 87: 5 green, 5 red, 15 blue; 4 blue, 12 red, 10 green; 3 green, 11 blue, 9 red; 3 red, 4 green, 16 blue; 3 red, 10 blue, 10 green; 15 blue, 9 green, 12 red\nGame 88: 2 green, 10 blue; 4 blue, 8 green; 8 green, 1 blue; 13 blue, 1 red, 2 green; 2 green, 16 blue\nGame 89: 5 blue, 7 red; 10 red, 11 blue, 6 green; 6 green, 3 red, 7 blue; 5 green, 3 red, 20 blue; 8 red, 6 green, 10 blue; 7 blue, 5 green, 10 red\nGame 90: 4 red, 1 green, 4 blue; 9 red, 9 blue, 9 green; 4 green, 11 red; 9 red, 5 green, 3 blue; 9 red, 2 green, 2 blue\nGame 91: 13 green, 13 blue; 3 red, 11 green, 5 blue; 10 blue, 3 green, 1 red; 3 blue, 10 green, 2 red; 5 blue, 2 green\nGame 92: 8 blue, 1 green, 4 red; 3 blue, 6 red; 3 red, 1 green, 14 blue; 6 blue, 8 red; 15 blue, 9 red; 4 blue, 2 red\nGame 93: 3 blue, 17 red, 2 green; 9 blue, 6 red; 6 blue, 2 green, 16 red; 1 green, 5 blue, 15 red; 3 blue, 2 green, 14 red\nGame 94: 7 blue, 19 green, 1 red; 4 blue; 8 blue, 3 red, 4 green\nGame 95: 2 green, 6 red, 13 blue; 5 red, 12 green, 12 blue; 18 blue, 8 red, 4 green; 7 red, 6 green, 17 blue; 4 green, 9 red, 6 blue; 10 red, 1 green, 4 blue\nGame 96: 8 blue, 9 red; 9 red, 10 blue; 5 blue, 1 green, 2 red; 2 blue, 2 red\nGame 97: 4 red, 1 blue, 2 green; 2 green, 11 red, 1 blue; 8 red, 1 green; 7 red, 3 green, 1 blue; 5 red, 1 green, 1 blue\nGame 98: 6 green, 4 blue, 12 red; 3 blue, 13 red, 1 green; 2 blue, 12 green, 2 red; 13 green, 2 red, 1 blue; 10 red, 7 green, 1 blue\nGame 99: 6 blue, 3 green, 5 red; 3 green, 6 red, 8 blue; 3 green, 11 blue, 14 red; 14 red, 5 green, 1 blue\nGame 100: 16 red, 3 blue; 2 red, 5 green; 9 red; 1 blue, 3 green, 10 red; 1 red, 5 blue, 3 green; 12 blue, 9 red"

func Execute() {
	lines := strings.Split(input, "\n")
	sum := 0
	for _, line := range lines {
		pp := strings.Split(line, ":")
		game := strings.Replace(pp[0], "Game ", "", 1)
		pwr := power(pp[1])
		fmt.Printf("game %s pwr %d\n", game, pwr)
		sum += pwr
	}
	fmt.Printf("Answer: %d\n", sum)
}

func power(pulls string) int {
	pp := parsePulls(pulls)
	maxes := [3]int{0, 0, 0}
	for _, p := range pp {
		for i, c := range p {
			if c > maxes[i] {
				maxes[i] = c
			}
		}
	}
	return maxes[0] * maxes[1] * maxes[2]
}

var rgbMaxes = [3]int{12, 13, 14}

func parsePull(pull string) [3]int {
	pullColours := strings.Split(pull, ",")
	var res = [3]int{0, 0, 0}
	for _, pullColour := range pullColours {
		clean := strings.Trim(pullColour, " ")
		parts := strings.Split(clean, " ")
		amount, err := strconv.Atoi(parts[0])
		if err != nil {
			panic(err)
		}
		colour := parts[1]
		switch colour {
		case "red":
			res[0] = amount
		case "green":
			res[1] = amount
		case "blue":
			res[2] = amount
		}
	}
	return res
}

func parsePulls(pullsStr string) [][3]int {
	pullStr := strings.Split(pullsStr, ";")
	var rgbs [][3]int
	for _, pull := range pullStr {
		rgbs = append(rgbs, parsePull(pull))
	}
	return rgbs
}

func valid(pullsStr string) bool {
	pulls := parsePulls(pullsStr)
	for _, pull := range pulls {
		for i, c := range pull {
			if c > rgbMaxes[i] {
				return false
			}
		}
	}
	return true
}

func ExecutePart1() {
	lines := strings.Split(input, "\n")
	sum := 0
	for _, line := range lines {
		pp := strings.Split(line, ":")
		game := strings.Replace(pp[0], "Game ", "", 1)

		fmt.Printf("game %s\n", game)
		if valid(pp[1]) {
			val, err := strconv.Atoi(game)
			if err != nil {
				panic(err)
			}
			sum += val
		}
	}
	fmt.Printf("Answer: %d\n", sum)
}
