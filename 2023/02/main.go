package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	solve()
}

func solve() {
	ans1 := 0
	ans2 := 0
	for _, line := range strings.Split(input, "\n") {
		pre, content, _ := strings.Cut(line, ": ")
		var id int
		fmt.Sscanf(pre, "Game %d:", &id)
		possible := true
		m := make(map[string]int)
		for _, round := range strings.Split(content, "; ") {
			for _, cubes := range strings.Split(round, ", ") {
				var (
					cnt   int
					color string
				)
				fmt.Sscanf(cubes, "%d %s", &cnt, &color)
				m[color] = max(m[color], cnt)
				if !check(color, cnt) {
					possible = false
				}
			}
		}
		if possible {
			ans1 += id
		}
		value := 1
		for _, v := range m {
			value *= v
		}
		ans2 += value
	}
	println(ans1)
	println(ans2)
}

func check(color string, cnt int) bool {
	switch color {
	case "red":
		return cnt <= 12
	case "blue":
		return cnt <= 14
	case "green":
		return cnt <= 13
	}
	return true
}
