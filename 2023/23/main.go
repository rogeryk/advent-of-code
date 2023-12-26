package main

import (
	_ "embed"
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
	g := strings.Split(input, "\n")
	m, n := len(g), len(g[0])
	start := strings.Index(g[0], ".")
	end := strings.Index(g[m-1], ".")
	vis := map[[2]int]bool{}
	ans := 0
	moves := map[byte][][]int{
		'.': {
			{1, 0},
			{0, 1},
			{-1, 0},
			{0, -1},
		},
		'>': {
			{0, 1},
		},
		'<': {
			{0, -1},
		},
		'^': {
			{-1, 0},
		},
		'v': {
			{1, 0},
		},
	}
	var dfs func(int, int, int)
	dfs = func(r, c, d int) {
		if r == m-1 && c == end {
			ans = max(ans, d)
		}
		vis[[2]int{r, c}] = true
		defer delete(vis, [2]int{r, c})
		for _, move := range moves[g[r][c]] {
			nr, nc := r+move[0], c+move[1]
			if nr < 0 || nr >= m || nc < 0 || nc >= n || g[nr][nc] == '#' || vis[[2]int{nr, nc}] {
				continue
			}
			dfs(nr, nc, d+1)
		}
	}
	dfs(0, start, 0)
	println(ans)
}

func solve2() {
	moves := [][]int{
		{1, 0},
		{0, 1},
		{-1, 0},
		{0, -1},
	}
	lines := strings.Split(input, "\n")
	m, n := len(lines), len(lines[0])
	points := make([][2]int, 0)
	for r, line := range lines {
		for c, v := range line {
			if v == '#' {
				continue
			}
			degree := 0
			for _, move := range moves {
				nr, nc := r+move[0], c+move[1]
				if nr >= 0 && nr < m && nc >= 0 && nc < n && lines[nr][nc] != '#' {
					degree++
				}
			}
			if degree >= 3 {
				points = append(points, [2]int{r, c})
			}
		}
	}
	start := strings.Index(lines[0], ".")
	end := strings.Index(lines[m-1], ".")
	points = append(points, [2]int{0, start})
	points = append(points, [2]int{m - 1, end})
	g := make([][][]int, len(points))
	for i, p := range points {
		q := [][]int{{p[0], p[1], 0}}
		vis := make(map[[2]int]bool)
		vis[p] = true
		for len(q) != 0 {
			v := q[0]
			q = q[1:]
			for _, move := range moves {
				nr, nc, nd := v[0]+move[0], v[1]+move[1], v[2]+1
				np := [2]int{nr, nc}
				if nr < 0 || nr >= m || nc < 0 || nc >= n || lines[nr][nc] == '#' || vis[np] {
					continue
				}
				vis[np] = true
				if j := slices.Index(points, np); j != -1 {
					g[i] = append(g[i], []int{j, nd})
					continue
				}
				q = append(q, []int{nr, nc, nd})
			}
		}
	}
	ans := 0
	var dfs func(int, int)
	vis := map[int]bool{}
	dfs = func(p, d int) {
		if p == len(points)-1 {
			ans = max(ans, d)
		}
		vis[p] = true
		defer delete(vis, p)
		for _, v := range g[p] {
			np, nd := v[0], v[1]+d
			if vis[np] {
				continue
			}
			dfs(np, nd)
		}
	}
	dfs(len(points)-2, 0)
	println(ans)
}
