package main

import (
	_ "embed"
	"strings"

	"github.com/zyedidia/generic/heap"
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
	moves := [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	q := heap.New(func(a, b []int) bool {
		return a[4] < b[4]
	})
	vis := make(map[[4]int]bool, m*n)
	q.Push([]int{0, 0, 0, 0, 0})
	vis[[4]int{0, 0, 0, 0}] = true
	ans := 0
	for ans == 0 && q.Size() != 0 {
		p, _ := q.Pop()
		r, c, d, s, t := p[0], p[1], p[2], p[3], p[4]
		for i, v := range moves {
			if ((i+2)%4 == d) || (i == d && s == 3) {
				continue
			}
			rr, cc := r+v[0], c+v[1]
			if rr < 0 || rr >= m || cc < 0 || cc >= n {
				continue
			}
			if rr == m-1 && cc == n-1 {
				ans = t + int(g[rr][cc]-'0')
				break
			}
			ss := 1
			if i == d {
				ss = s + 1
			}
			k := [4]int{rr, cc, i, ss}
			if vis[k] {
				continue
			}
			vis[k] = true
			q.Push([]int{rr, cc, i, ss, t + int(g[rr][cc]-'0')})
		}
	}
	println(ans)
}
func solve2() {
	g := strings.Split(input, "\n")
	m, n := len(g), len(g[0])
	moves := [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	q := heap.New(func(a, b []int) bool {
		return a[4] < b[4]
	})
	vis := make(map[[4]int]bool, m*n)
	q.Push([]int{0, 0, 0, 0, 0})
	vis[[4]int{0, 0, 0, 0}] = true
	ans := 0
	for ans == 0 && q.Size() != 0 {
		p, _ := q.Pop()
		r, c, d, s, t := p[0], p[1], p[2], p[3], p[4]
		for i, v := range moves {
			if ((i+2)%4 == d) || (s < 4 && i != d) || (i == d && s == 10) {
				continue
			}
			rr, cc := r+v[0], c+v[1]
			if rr < 0 || rr >= m || cc < 0 || cc >= n {
				continue
			}
			if rr == m-1 && cc == n-1 {
				ans = t + int(g[rr][cc]-'0')
				break
			}
			ss := 1
			if i == d {
				ss = s + 1
			}
			k := [4]int{rr, cc, i, ss}
			if vis[k] {
				continue
			}
			vis[k] = true
			q.Push([]int{rr, cc, i, ss, t + int(g[rr][cc]-'0')})
		}
	}
	print(ans)
}
