package main

import (
	_ "embed"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

//go:embed input.txt
var input string

func main() {
	solve1()
	solve2()
}

func solve1() {
	lines := strings.Split(input, "\n")
	p := regexp.MustCompile(`\d+`)
	ans := 0
	for i, line := range lines {
		for _, pos := range p.FindAllStringIndex(line, -1) {
			valid := false
			if pos[0] > 0 && line[pos[0]-1] != '.' {
				valid = true
			}
			if !valid && pos[1] < len(line) && line[pos[1]] != '.' {
				valid = true
			}
			if !valid && i > 0 {
				valid = strings.IndexFunc(lines[i-1][max(0, pos[0]-1):min(len(lines[i-1]), pos[1]+1)], func(r rune) bool {
					return !unicode.IsDigit(r) && r != '.'
				}) != -1
			}
			if !valid && i < len(lines)-1 {
				valid = strings.IndexFunc(lines[i+1][max(0, pos[0]-1):min(len(lines[i+1]), pos[1]+1)], func(r rune) bool {
					return !unicode.IsDigit(r) && r != '.'
				}) != -1
			}
			if valid {
				num, _ := strconv.Atoi(line[pos[0]:pos[1]])
				ans += num
			}
		}
	}
	println(ans)
}

func solve2() {
	lines := strings.Split(input, "\n")
	p := regexp.MustCompile(`\d+`)
	ans := 0
	n := len(lines[0])
	m := map[int][]int{}
	for i, line := range lines {
		for _, pos := range p.FindAllStringIndex(line, -1) {
			valid := false
			gear := -1
			if pos[0] > 0 && line[pos[0]-1] != '.' {
				valid = true
				if line[pos[0]-1] == '*' {
					gear = i*n + pos[0] - 1
				}
			}
			if pos[1] < len(line) && line[pos[1]] != '.' {
				valid = true
				if line[pos[1]] == '*' {
					gear = i*n + pos[1]
				}
			}
			if i > 0 {
				valid = valid || strings.IndexFunc(lines[i-1][max(0, pos[0]-1):min(len(lines[i-1]), pos[1]+1)], func(r rune) bool {
					return !unicode.IsDigit(r) && r != '.'
				}) != -1
				gi := strings.IndexByte(lines[i-1][max(0, pos[0]-1):min(len(lines[i-1]), pos[1]+1)], '*')
				if gi != -1 {
					gear = (i-1)*n + gi + max(0, pos[0]-1)
				}
			}
			if i < len(lines)-1 {
				valid = valid || strings.IndexFunc(lines[i+1][max(0, pos[0]-1):min(len(lines[i+1]), pos[1]+1)], func(r rune) bool {
					return !unicode.IsDigit(r) && r != '.'
				}) != -1
				gi := strings.IndexByte(lines[i+1][max(0, pos[0]-1):min(len(lines[i+1]), pos[1]+1)], '*')
				if gi != -1 {
					gear = (i+1)*n + gi + max(0, pos[0]-1)
				}
			}
			if valid {
				num, _ := strconv.Atoi(line[pos[0]:pos[1]])
				if gear != -1 {
					m[gear] = append(m[gear], num)
				}
			}
		}
	}
	for _, nums := range m {
		if len(nums) == 2 {
			ans += nums[0] * nums[1]
		}
	}
	println(ans)
}
