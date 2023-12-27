package main

import (
	"cmp"
	_ "embed"
	"slices"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	g := map[string][]string{}
	for _, line := range strings.Split(input, "\n") {
		node, content, _ := strings.Cut(line, ": ")
		for _, next := range strings.Split(content, " ") {
			g[node] = append(g[node], next)
			g[next] = append(g[next], node)
		}
	}
	edgeUses := map[string]int{}
	for node := range g {
		q := [][]any{{node, 0}}
		vis := map[string]bool{}
		prev := map[string]string{}
		for len(q) > 0 {
			n, d := q[0][0].(string), q[0][1].(int)
			q = q[1:]
			for _, next := range g[n] {
				if !vis[next] {
					vis[next] = true
					prev[next] = n
					q = append(q, []any{next, d + 1})
				}
			}
		}
		for _, n := range prev {
			for n != node {
				nx := prev[n]
				key := []string{n, nx}
				slices.Sort(key)
				edgeUses[strings.Join(key, ",")]++
				n = nx
			}
		}
	}

	edges := [][]string{}
	for k := range edgeUses {
		edges = append(edges, strings.Split(k, ","))
	}
	slices.SortFunc(edges, func(a, b []string) int {
		return cmp.Compare(edgeUses[strings.Join(a, ",")], edgeUses[strings.Join(b, ",")])
	})
	for _, e := range edges[len(edges)-3:] {
		i := slices.Index(g[e[0]], e[1])
		j := slices.Index(g[e[1]], e[0])
		g[e[0]] = slices.Delete(g[e[0]], i, i+1)
		g[e[1]] = slices.Delete(g[e[1]], j, j+1)
	}

	var node string
	for n := range g {
		node = n
		break
	}
	q := []string{node}
	a := 1
	vis := map[string]bool{}
	vis[node] = true
	for len(q) > 0 {
		n := q[0]
		q = q[1:]
		for _, next := range g[n] {
			if !vis[next] {
				a++
				vis[next] = true
				q = append(q, next)
			}
		}
	}
	println((len(g) - a) * a)
}
