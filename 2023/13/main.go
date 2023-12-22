package main

import (
	_ "embed"
	"math/bits"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	solve1()
	solve2()
}

func solve1() {
	patterns := strings.Split(input, "\n\n")
	ans := 0
	for _, pattern := range patterns {
		lines := strings.Split(pattern, "\n")
		rows := make([]int, len(lines))
		cols := make([]int, len(lines[0]))
		for i, line := range lines {
			for j, v := range line {
				if v == '#' {
					rows[i] |= 1 << j
					cols[j] |= 1 << i
				}
			}
		}
		for _, v := range mirrors(rows) {
			ans += v * 100
		}
		for _, v := range mirrors(cols) {
			ans += v
		}
	}
	println(ans)
}

func solve2() {
	patterns := strings.Split(input, "\n\n")
	ans := 0
	for _, pattern := range patterns {
		lines := strings.Split(pattern, "\n")
		rows := make([]int, len(lines))
		cols := make([]int, len(lines[0]))
		for i, line := range lines {
			for j, v := range line {
				if v == '#' {
					rows[i] |= 1 << j
					cols[j] |= 1 << i
				}
			}
		}
		for _, v := range mirrors2(rows) {
			ans += v * 100
		}
		for _, v := range mirrors2(cols) {
			ans += v
		}
	}
	println(ans)
}

func mirrors(pattern []int) []int {
	var res []int
	for i := 1; i < len(pattern); i++ {
		mirror := true
		for j := max(0, 2*i-len(pattern)); j < i; j++ {
			if pattern[j] != pattern[2*i-j-1] {
				mirror = false
				break
			}
		}
		if mirror {
			res = append(res, i)
		}
	}
	return res
}

func mirrors2(pattern []int) []int {
	var res []int
	for i := 1; i < len(pattern); i++ {
		mirror := true
		smudge := 0
		for j := max(0, 2*i-len(pattern)); j < i; j++ {
			if bits.OnesCount(uint(pattern[j]^pattern[2*i-j-1])) == 1 {
				smudge++
				continue
			} else if pattern[j] != pattern[2*i-j-1] {
				mirror = false
				break
			}
		}
		if mirror && smudge == 1 {
			res = append(res, i)
		}
	}
	return res
}
