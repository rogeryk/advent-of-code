package main

import (
	"cmp"
	_ "embed"
	"math"
	"slices"
	"sort"
	"strings"

	"github.com/spf13/cast"
)

//go:embed input.txt
var input string

// func init() {
// 	input = `seeds: 79 14 55 13

// seed-to-soil map:
// 50 98 2
// 52 50 48

// soil-to-fertilizer map:
// 0 15 37
// 37 52 2
// 39 0 15

// fertilizer-to-water map:
// 49 53 8
// 0 11 42
// 42 0 7
// 57 7 4

// water-to-light map:
// 88 18 7
// 18 25 70

// light-to-temperature map:
// 45 77 23
// 81 45 19
// 68 64 13

// temperature-to-humidity map:
// 0 69 1
// 1 0 69

// humidity-to-location map:
// 60 56 37
// 56 93 4`
// }

func main() {
	solve1()
	solve2()
}

func solve1() {
	lines := strings.Split(input, "\n")
	lines = append(lines, "")
	_, content, _ := strings.Cut(lines[0], ": ")
	nums := []int{}
	for _, v := range strings.Split(content, " ") {
		nums = append(nums, cast.ToInt(v))
	}
	sort.Ints(nums)
	next := make([]int, 0, len(nums))
	visted := map[int]bool{}
	for i := 3; i < len(lines); i++ {
		if lines[i] == "" {
			for _, v := range nums {
				if !visted[v] {
					next = append(next, v)
				}
			}
			sort.Ints(next)
			nums = next
			next = make([]int, 0, len(nums))
			clear(visted)
			i++
			continue
		}
		mps := strings.Split(lines[i], " ")
		d, s, c := cast.ToInt(mps[0]), cast.ToInt(mps[1]), cast.ToInt(mps[2])
		p := sort.SearchInts(nums, s)
		for p < len(nums) && nums[p] < s+c {
			next = append(next, d+nums[p]-s)
			visted[nums[p]] = true
			p++
		}
	}
	println(slices.Min(nums))
}

func solve2() {
	lines := strings.Split(input, "\n")
	lines = append(lines, "")
	_, content, _ := strings.Cut(lines[0], ": ")
	seeds := strings.Split(content, " ")
	nums := [][]int{}
	for i := 0; i < len(seeds); i += 2 {
		nums = append(nums, []int{cast.ToInt(seeds[i]), cast.ToInt(seeds[i+1])})
	}
	slices.SortFunc(nums, func(a, b []int) int {
		return cmp.Compare(a[0], b[0])
	})
	next := make([][]int, 0, len(nums))
	for i := 3; i < len(lines); i++ {
		if lines[i] == "" {
			for _, v := range nums {
				if v[1] != 0 {
					next = append(next, v)
				}
			}
			slices.SortFunc(next, func(a, b []int) int {
				return cmp.Compare(a[0], b[0])
			})
			nums = next
			next = make([][]int, 0, len(nums))
			i++
			continue
		}
		mps := strings.Split(lines[i], " ")
		d, s, c := cast.ToInt(mps[0]), cast.ToInt(mps[1]), cast.ToInt(mps[2])
		p := sort.Search(len(nums), func(p int) bool {
			return nums[p][0] >= s || nums[p][0]+nums[p][1] >= s
		})
		for p < len(nums) && (nums[p][0] < s+c) {
			if nums[p][1] == 0 {
				p++
				continue
			}
			if nums[p][0] >= s {
				next = append(next, []int{d + nums[p][0] - s, min(nums[p][1], s+c-nums[p][0])})
				nums[p][1] -= min(nums[p][1], s+c-nums[p][0])
			} else {
				next = append(next, []int{d, nums[p][0] + nums[p][1] - s})
				nums[p][1] -= nums[p][0] + nums[p][1] - s
			}
			p++
		}
	}
	ans := math.MaxInt
	for _, v := range nums {
		ans = min(ans, v[0])
	}
	println(ans)
}
