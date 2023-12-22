package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	solve1()
	solve2()
}

func solve1() {
	lines := strings.Split(input, "\n")
	path := lines[0]
	g := make(map[string][]string)
	for i := 2; i < len(lines); i++ {
		g[lines[i][:3]] = append(g[lines[i][:3]], lines[i][7:10], lines[i][12:15])
	}
	node := "AAA"
	steps := 0
	for ; node != "ZZZ"; steps++ {
		d := path[steps%len(path)]
		if d == 'L' {
			node = g[node][0]
		} else {
			node = g[node][1]
		}
	}
	println(steps)
}

func solve2() {
	lines := strings.Split(input, "\n")
	path := lines[0]
	g := make(map[string][]string)
	nodes := make([]string, 0)
	for i := 2; i < len(lines); i++ {
		node := lines[i][:3]
		g[node] = append(g[node], lines[i][7:10], lines[i][12:15])
		if node[2] == 'A' {
			nodes = append(nodes, node)
		}
	}
	nodeSteps := make([]int, len(nodes))
	for i, node := range nodes {
		steps := 0
		for ; node[2] != 'Z'; steps++ {
			d := path[steps%len(path)]
			if d == 'L' {
				node = g[node][0]
			} else {
				node = g[node][1]
			}
		}
		nodeSteps[i] = steps
	}
	ans := nodeSteps[0]
	for _, step := range nodeSteps[1:] {
		ans = lcm(ans, step)
	}

	fmt.Println(ans)
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}
