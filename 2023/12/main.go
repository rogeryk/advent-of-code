package main

import (
	_ "embed"
	"fmt"
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
	for _, line := range strings.Split(input, "\n") {
		springs, groups, _ := strings.Cut(line, " ")
		var numbers []int
		for _, v := range strings.Split(groups, ",") {
			numbers = append(numbers, cast.ToInt(v))
		}
		ans += possiable(springs, numbers)
	}
	println(ans)
}

func solve2() {
	ans := 0
	for _, line := range strings.Split(input, "\n") {
		springs, groups, _ := strings.Cut(line, " ")
		springs = strings.Repeat(springs+"?", 5)
		springs = springs[:len(springs)-1]
		groups = strings.Repeat(groups+",", 5)
		groups = groups[:len(groups)-1]
		var numbers []int
		for _, v := range strings.Split(groups, ",") {
			numbers = append(numbers, cast.ToInt(v))
		}
		ans += possiable(springs, numbers)
	}
	println(ans)
}

func possiable(s string, nums []int) int {
	var dfs func(int, int, int) int
	m, n := len(s), len(nums)
	memo := make(map[string]int)
	dfs = func(i, j, c int) int {
		key := fmt.Sprintf("%d-%d-%d", i, j, c)
		if _, ok := memo[key]; ok {
			return memo[key]
		}
		if i == m {
			if j == n || (j == n-1 && c == nums[j]) {
				return 1
			}
			return 0
		}
		ans := 0
		if s[i] == '.' || s[i] == '?' {
			if c != 0 {
				if j < n && c == nums[j] {
					ans += dfs(i+1, j+1, 0)
				}
			} else {
				ans += dfs(i+1, j, 0)
			}
		}
		if s[i] == '#' || s[i] == '?' {
			if j < n && c < nums[j] {
				ans += dfs(i+1, j, c+1)
			}
		}
		memo[key] = ans
		return ans
	}

	return dfs(0, 0, 0)
}
