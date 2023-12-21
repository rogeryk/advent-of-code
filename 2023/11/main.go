package main

import (
	_ "embed"
	"strings"
)

//go:embed input.txt
var input string

// func init() {
// 	input = `...#......
// .......#..
// #.........
// ..........
// ......#...
// .#........
// .........#
// ..........
// .......#..
// #...#.....`
// }

func main() {
	solve1()
	solve2()
}

func solve1() {
	lines := strings.Split(input, "\n")
	m, n := len(lines), len(lines[0])
	rows := make([]int, m)
	cols := make([]int, n)
	for i, line := range lines {
		for j, v := range line {
			if v == '#' {
				rows[i]++
				cols[j]++
			}
		}
	}
	poses := make([][]int, 0)
	r, c := 0, 0
	for i, line := range lines {
		if rows[i] == 0 {
			r += 2
			continue
		}
		r++
		for j, v := range line {
			if cols[j] == 0 {
				c += 2
				continue
			}
			c++
			if v == '#' {
				poses = append(poses, []int{r, c})
			}
		}
		c = 0
	}
	ans := 0
	for i := range poses {
		for j := i + 1; j < len(poses); j++ {
			ans += abs(poses[j][0]-poses[i][0]) + abs(poses[j][1]-poses[i][1])
		}
	}
	println(ans)
}

func solve2() {
	lines := strings.Split(input, "\n")
	m, n := len(lines), len(lines[0])
	rows := make([]int, m)
	cols := make([]int, n)
	for i, line := range lines {
		for j, v := range line {
			if v == '#' {
				rows[i]++
				cols[j]++
			}
		}
	}
	poses := make([][]int, 0)
	r, c := 0, 0
	for i, line := range lines {
		if rows[i] == 0 {
			r += 1000_000
			continue
		}
		r++
		for j, v := range line {
			if cols[j] == 0 {
				c += 1000_000
				continue
			}
			c++
			if v == '#' {
				poses = append(poses, []int{r, c})
			}
		}
		c = 0
	}
	ans := 0
	for i := range poses {
		for j := i + 1; j < len(poses); j++ {
			ans += abs(poses[j][0]-poses[i][0]) + abs(poses[j][1]-poses[i][1])
		}
	}
	println(ans)
}

func abs(v int) int {
	if v < 0 {
		return v * -1
	}
	return v
}
