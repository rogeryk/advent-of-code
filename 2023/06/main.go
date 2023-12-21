package main

import (
	_ "embed"
	"math"
	"regexp"
	"strings"

	"github.com/spf13/cast"
)

//go:embed input.txt
var input string

func main() {
	solve1()
	solve2()
}

func solve1() {
	lines := strings.Split(input, "\n")
	r := regexp.MustCompile(`\d+`)
	var times, distances []int
	for _, v := range r.FindAllString(lines[0], -1) {
		times = append(times, cast.ToInt(v))
	}
	for _, v := range r.FindAllString(lines[1], -1) {
		distances = append(distances, cast.ToInt(v))
	}
	ans := 1
	for i, t := range times {
		c := 0
		for x := 1; x < t; x++ {
			if x*(t-x) > distances[i] {
				c++
			}
		}
		ans *= c
	}
	println(ans)
}

func solve2() {
	lines := strings.Split(input, "\n")
	r := regexp.MustCompile(`\d+`)
	t := cast.ToFloat64(strings.Join(r.FindAllString(lines[0], -1), ""))
	d := cast.ToFloat64(strings.Join(r.FindAllString(lines[1], -1), ""))
	bound := math.Pow(math.Pow(t, 2)/4-d, 0.5)
	h := math.Floor(bound + t/2)
	l := math.Ceil(-bound + t/2)
	ans := h - l + 1
	println(int(ans))
}
