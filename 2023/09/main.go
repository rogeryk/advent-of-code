package main

import (
	_ "embed"
	"strconv"
	"strings"
)

func main() {
	solve1()
	solve2()
}

//go:embed input.txt
var input string

//func init() {
//	input = `0 3 6 9 12 15
//1 3 6 10 15 21
//10 13 16 21 30 45`
//}

func solve1() {
	seqs := parseInput()
	sum := 0
	for _, seq := range seqs {
		diffs := make([][]int, 0, len(seq))
		diffs = append(diffs, seq)
		prevDiff := seq
		for {
			diff := make([]int, 0, len(prevDiff))
			allZero := true
			for i := 1; i < len(prevDiff); i++ {
				d := prevDiff[i] - prevDiff[i-1]
				diff = append(diff, d)
				if d != 0 {
					allZero = false
				}
			}
			prevDiff = diff
			diffs = append(diffs, diff)
			if allZero {
				break
			}
		}
		v := 0
		for _, diff := range diffs {
			v += diff[len(diff)-1]
		}
		sum += v
	}
	println(sum)
}

func parseInput() [][]int {
	lines := strings.Split(input, "\n")
	arrs := make([][]int, len(lines))
	for i, line := range lines {
		for _, v := range strings.Split(line, " ") {
			num, _ := strconv.Atoi(v)
			arrs[i] = append(arrs[i], num)
		}
	}
	return arrs
}

func solve2() {
	seqs := parseInput()
	sum := 0
	for _, seq := range seqs {
		diffs := make([][]int, 0, len(seq))
		diffs = append(diffs, seq)
		prevDiff := seq
		for {
			diff := make([]int, 0, len(prevDiff))
			allZero := true
			for i := 1; i < len(prevDiff); i++ {
				d := prevDiff[i] - prevDiff[i-1]
				diff = append(diff, d)
				if d != 0 {
					allZero = false
				}
			}
			prevDiff = diff
			diffs = append(diffs, diff)
			if allZero {
				break
			}
		}
		flag := true
		v := 0
		for _, diff := range diffs {
			if flag {
				v += diff[0]
			} else {
				v -= diff[0]
			}
			flag = !flag
		}
		sum += v
	}
	println(sum)
}
