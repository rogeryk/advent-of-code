package main

import (
	_ "embed"
	"strconv"
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
	points := [][2]int{}
	b := 0
	p := [2]int{0, 0}
	moves := map[string][]int{"R": {0, 1}, "D": {1, 0}, "L": {0, -1}, "U": {-1, 0}}
	for _, line := range strings.Split(input, "\n") {
		values := strings.Split(line, " ")
		n := cast.ToInt(values[1])
		b += n
		p[0] += moves[values[0]][0] * n
		p[1] += moves[values[0]][1] * n
		points = append(points, p)
	}
	println(getArea(points, b))
}

func solve2() {
	points := [][2]int{}
	b := 0
	p := [2]int{0, 0}
	moves := map[string][]int{"0": {0, 1}, "1": {1, 0}, "2": {0, -1}, "3": {-1, 0}}
	for _, line := range strings.Split(input, "\n") {
		valies := strings.Split(line, " ")
		color := valies[2]
		n, _ := strconv.ParseInt(color[2:7], 16, 0)
		b += int(n)
		p[0] += moves[color[7:8]][0] * int(n)
		p[1] += moves[color[7:8]][1] * int(n)
		points = append(points, p)
	}
	println(getArea(points, b))
}

func getArea(points [][2]int, l int) int {
	A := 0
	a := points[len(points)-1]
	for _, b := range points {
		A += (a[0] - b[0]) * (a[1] + b[1])
		a = b
	}
	A /= 2
	if A < 0 {
		A = -A
	}
	i := A - l/2 + 1
	return i + l
}
