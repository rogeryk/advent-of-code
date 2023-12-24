package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

// func init() {
// 	input = `...........
// .....###.#.
// .###.##..#.
// ..#.#...#..
// ....#.#....
// .##..S####.
// .##..#...#.
// .......##..
// .##.#.####.
// .##..##.##.
// ...........`
// }

func main() {
	solve()
}

func solve() {
	g := strings.Split(input, "\n")
	m, n := len(g), len(g[0])
	q := [][]int{}
	vis := map[string]bool{}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if g[i][j] == 'S' {
				q = append(q, []int{0, i, j})
				vis[fmt.Sprint(i, j)] = true
			}
		}
	}
	steps := [][]int{}
	for len(q) != 0 {
		p := q[0]
		steps = append(steps, p)
		q = q[1:]
		for _, d := range [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} {
			d, i, j := p[0]+1, p[1]+d[0], p[2]+d[1]
			if i >= 0 && i < m && j >= 0 && j < n {
				if g[i][j] != '#' && !vis[fmt.Sprint(i, j)] {
					vis[fmt.Sprint(i, j)] = true
					q = append(q, []int{d, i, j})
				}
			}
		}
	}
	ans := 0
	for _, step := range steps {
		if step[0] <= 64 && step[0]%2 == 0 {
			ans++
		}
	}

	println(ans)

	var odd, oddCorn, even, evenCorn int
	for _, v := range steps {
		if v[0]%2 == 0 {
			even++
			if v[0] > 65 {
				evenCorn++
			}
		} else {
			odd++
			if v[0] > 65 {
				oddCorn++
			}
		}
	}
	cnt := 26501365 / m
	println(cnt, odd, oddCorn, even, evenCorn)
	ans2 := cnt*cnt*even + (cnt+1)*(cnt+1)*odd - (cnt+1)*oddCorn + cnt*evenCorn
	println(ans2)
}
