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
	layout := strings.Split(input, "\n")
	println(energized(layout, 0, 0, 0))
}

func solve2() {
	layout := strings.Split(input, "\n")
	m, n := len(layout), len(layout[0])
	ans := 0
	for i := 0; i < m; i++ {
		ans = max(ans, energized(layout, i, 0, 0))
		ans = max(ans, energized(layout, i, n-1, 2))
	}
	for i := 0; i < n; i++ {
		ans = max(ans, energized(layout, 0, i, 1))
		ans = max(ans, energized(layout, m-1, i, 3))
	}
	println(ans)
}

func energized(layout []string, r, c, d int) int {
	m, n := len(layout), len(layout[0])
	visited := make([][]uint, m)
	for i := range visited {
		visited[i] = make([]uint, n)
	}
	var dfs func(int, int, int)
	moves := [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	ans := 0
	dfs = func(r, c, d int) {
		if r < 0 || r >= m || c < 0 || c >= n {
			return
		}
		if (visited[r][c] & (1 << d)) != 0 {
			return
		}
		if visited[r][c] == 0 {
			ans++
		}
		visited[r][c] |= (1 << d)
		if layout[r][c] == '|' && (d == 0 || d == 2) {
			dfs(r+1, c, 1)
			dfs(r-1, c, 3)
		} else if layout[r][c] == '-' && (d == 1 || d == 3) {
			dfs(r, c+1, 0)
			dfs(r, c-1, 2)
		} else {
			if layout[r][c] == '/' {
				if d%2 != 0 {
					d = (d + 1) % 4
				} else {
					d = (d + 3) % 4
				}
			} else if layout[r][c] == '\\' {
				if d%2 != 0 {
					d = (d + 3) % 4
				} else {
					d = (d + 1) % 4
				}
			}
			r += moves[d][0]
			c += moves[d][1]
			dfs(r, c, d)
		}
	}
	dfs(r, c, d)
	return ans
}
