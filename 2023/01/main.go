package main

import (
	_ "embed"
	"regexp"
	"slices"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	solve1()
	solve2()
}

func solve1() {
	sum := 0
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}

		digitRegex := regexp.MustCompile(`\d`)
		all := digitRegex.FindAllString(line, -1)
		d1 := toNumber(all[0])
		d2 := toNumber(all[len(all)-1])
		sum += d1*10 + d2
	}
	println(sum)
}

func solve2() {
	sum := 0
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}

		letterDigits := "one|two|three|four|five|six|seven|eight|nine"
		digitRegex := regexp.MustCompile(`\d|` + letterDigits)
		reverseRegex := regexp.MustCompile(`\d|` + reverse(letterDigits))
		d1 := toNumber(digitRegex.FindString(line))
		d2 := toNumber(reverse(reverseRegex.FindString(reverse(line))))
		sum += d1*10 + d2
	}
	println(sum)
}
func reverse(s string) string {
	bs := []byte(s)
	slices.Reverse(bs)
	return string(bs)
}

func toNumber(digits string) int {
	switch digits {
	case "one": return 1
	case "two": return 2
	case "three": return 3
	case "four": return 4
	case "five": return 5
	case "six": return 6
	case "seven": return 7
	case "eight": return 8
	case "nine": return 9
	default:
		return int(digits[0] - '0')
	}
}
