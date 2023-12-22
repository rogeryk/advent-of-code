package main

import (
	_ "embed"
	"slices"
	"strings"

	"github.com/spf13/cast"
)

//go:embed input.txt
var input string

func main() {
	solve1()
	solve2()
}

func solve1() {
	ans := 0
	for _, v := range strings.Split(input, ",") {
		ans += hash(v)
	}
	println(ans)
}

func hash(s string) int {
	h := 0
	for _, c := range s {
		h = (h + int(c)) * 17 % 256
	}
	return h
}

func solve2() {
	type lens struct {
		label string
		focal int
	}
	boxes := make([][]lens, 256)
	for _, v := range strings.Split(input, ",") {
		if v[len(v)-1] == '-' {
			label := v[:len(v)-1]
			h := hash(label)
			i := slices.IndexFunc(boxes[h], func(l lens) bool {
				return l.label == label
			})
			if i != -1 {
				boxes[h][i].label = ""
			}
			continue
		}
		label, focal, _ := strings.Cut(v, "=")
		h := hash(label)
		i := slices.IndexFunc(boxes[h], func(l lens) bool {
			return l.label == label
		})
		if i == -1 {
			boxes[h] = append(boxes[h], lens{label: label, focal: cast.ToInt(focal)})
		} else {
			boxes[h][i].focal = cast.ToInt(focal)
		}
	}
	ans := 0
	for i, box := range boxes {
		j := 0
		for _, l := range box {
			if l.label == "" {
				continue
			}
			j++
			ans += (i + 1) * j * l.focal
		}
	}
	println(ans)
}
