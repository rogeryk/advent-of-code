package main

import (
	_ "embed"
	"maps"
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
	flowPart, ratingPart, _ := strings.Cut(input, "\n\n")
	flows := make(map[string][]string)
	for _, line := range strings.Split(flowPart, "\n") {
		label, content, _ := strings.Cut(line, "{")
		flows[label] = strings.Split(content[:len(content)-1], ",")
	}
	var checkFunc func(rating map[string]int, label string) bool
	checkFunc = func(rating map[string]int, label string) bool {
		if label == "A" {
			return true
		}
		if label == "R" {
			return false
		}
		for _, v := range flows[label] {
			if rule, d, ok := strings.Cut(v, ":"); ok {
				k := rule[0:1]
				v := rule[2:]
				if rule[1] == '>' && rating[k] > cast.ToInt(v) {
					return checkFunc(rating, d)
				}
				if rule[1] == '<' && rating[k] < cast.ToInt(v) {
					return checkFunc(rating, d)
				}
				continue
			}
			return checkFunc(rating, v)
		}
		return false
	}
	ans := 0
	for _, line := range strings.Split(ratingPart, "\n") {
		rating := map[string]int{}
		s := 0
		for _, pair := range strings.Split(line[1:len(line)-1], ",") {
			k, v, _ := strings.Cut(pair, "=")
			x := cast.ToInt(v)
			s += x
			rating[k] = x
		}
		if checkFunc(rating, "in") {
			ans += s
		}
	}
	println(ans)
}

func solve2() {
	flowPart, _, _ := strings.Cut(input, "\n\n")
	flows := make(map[string][]string)
	for _, line := range strings.Split(flowPart, "\n") {
		label, content, _ := strings.Cut(line, "{")
		flows[label] = strings.Split(content[:len(content)-1], ",")
	}
	var count func(map[string][]int, string)
	ans := 0
	count = func(ranges map[string][]int, label string) {
		if label == "A" {
			v := 1
			for _, r := range ranges {
				v *= (r[1] - r[0] + 1)
			}
			ans += v
			return
		}
		for _, v := range flows[label] {
			if rule, d, ok := strings.Cut(v, ":"); ok {
				k := rule[0:1]
				v := cast.ToInt(rule[2:])
				if rule[1] == '>' && ranges[k][1] > v {
					copy := maps.Clone(ranges)
					copy[k] = []int{max(v+1, copy[k][0]), copy[k][1]}
					ranges[k] = []int{ranges[k][0], v}
					count(copy, d)
				} else if rule[1] == '<' && ranges[k][0] < v {
					copy := maps.Clone(ranges)
					copy[k] = []int{copy[k][0], min(v-1, copy[k][1])}
					ranges[k] = []int{v, ranges[k][1]}
					count(copy, d)
				}
				if ranges[k][0] > ranges[k][1] {
					return
				}
				continue
			}
			count(ranges, v)
			return
		}
	}
	ranges := map[string][]int{}
	fields := "xmsa"
	for i := range fields {
		ranges[fields[i:i+1]] = []int{1, 4000}
	}
	count(ranges, "in")
	println(ans)
}
