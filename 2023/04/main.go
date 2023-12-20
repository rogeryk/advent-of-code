package main

import (
	_ "embed"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	solve1()
	solve2()
}

func solve1() {
	ans := 0
	for _, line := range strings.Split(input, "\n") {
		_, card, _ := strings.Cut(line, ": ")
		wins, nums, _ := strings.Cut(card, " | ")
		m := map[string]bool{}
		for _, v := range strings.Split(wins, " ") {
			if v != "" {
				m[v] = true
			}
		}
		cnt := 0
		for _, v := range strings.Split(nums, " ") {
			if m[v] {
				cnt++
			}
		}
		if cnt > 0 {
			ans += 1 << (cnt - 1)
		}
	}
	println(ans)
}

func solve2() {

	lines := strings.Split(input, "\n")
	n := len(lines)
	cnts := make([]int, n)
	for i, line := range lines {
		_, card, _ := strings.Cut(line, ": ")
		wins, nums, _ := strings.Cut(card, " | ")
		m := map[string]bool{}
		for _, v := range strings.Split(wins, " ") {
			if v != "" {
				m[v] = true
			}
		}
		cnt := 0
		for _, v := range strings.Split(nums, " ") {
			if m[v] {
				cnt++
			}
		}
		cnts[i]++
		for j := 1; j <= cnt; j++ {
			cnts[i+j] += cnts[i]
		}
	}
	ans := 0
	for _, v := range cnts {
		ans += v
	}
	println(ans)
}
