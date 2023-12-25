package main

import (
	"cmp"
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
	bricks, supportTo, supportBy := getBricks()
	for i := range bricks {
		flag := true
		for _, j := range supportTo[i] {
			if len(supportBy[j]) < 2 {
				flag = false
			}
		}
		if flag {
			ans++
		}
	}
	println(ans)
}

func solve2() {
	bricks, supportTo, supportBy := getBricks()
	ans := 0
	for i := range bricks {
		q := []int{i}
		falling := map[int]int{}
		for len(q) != 0 {
			v := q[0]
			q = q[1:]
			for _, j := range supportTo[v] {
				falling[j]++
				if len(supportBy[j]) <= falling[j] {
					q = append(q, j)
					ans++
				}
			}
		}
	}
	println(ans)
}

func getBricks() (bricks [][]int, supportTo, supportBy map[int][]int) {
	for _, line := range strings.Split(input, "\n") {
		start, end, _ := strings.Cut(line, "~")
		var startPos, endPos []int
		for _, v := range strings.Split(start, ",") {
			startPos = append(startPos, cast.ToInt(v))
		}
		for _, v := range strings.Split(end, ",") {
			endPos = append(endPos, cast.ToInt(v))
		}
		if startPos[2] <= endPos[2] {
			bricks = append(bricks, append(startPos, endPos...))
		} else {
			bricks = append(bricks, append(endPos, startPos...))
		}
	}
	slices.SortFunc(bricks, func(a, b []int) int {
		return cmp.Compare(a[2], b[2])
	})
	for i, a := range bricks {
		minZ := 1
		for _, b := range bricks[:i] {
			if overlap(a, b) {
				minZ = max(minZ, b[5]+1)
			}
		}
		fall := a[2] - minZ
		a[2] -= fall
		a[5] -= fall
	}
	slices.SortFunc(bricks, func(a, b []int) int {
		return cmp.Compare(a[2], b[2])
	})

	supportTo = map[int][]int{}
	supportBy = map[int][]int{}
	for i := 0; i < len(bricks); i++ {
		for j := i + 1; j < len(bricks); j++ {
			if support(bricks[i], bricks[j]) {
				supportTo[i] = append(supportTo[i], j)
				supportBy[j] = append(supportBy[j], i)
			}
		}
	}

	return
}

func overlap(a, b []int) bool {
	return max(a[0], b[0]) <= min(a[3], b[3]) && max(a[1], b[1]) <= min(a[4], b[4])
}

func support(a, b []int) bool {
	return overlap(a, b) && a[5]+1 == b[2]
}
