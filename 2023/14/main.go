package main

import (
	"bytes"
	_ "embed"
)

//go:embed input.txt
var input []byte

func main() {
	solve1()
	solve2()
}

func solve1() {
	dish := bytes.Split(input, []byte{'\n'})
	dish = move(dish)
	println(calcLoad(dish))
}

func solve2() {
	dish := bytes.Split(input, []byte{'\n'})
	dish = cycle(dish)
	println(calcLoad(dish))
}

func cycle(dish [][]byte) [][]byte {
	count := 1000_000_000 * 4
	seens := [][][]byte{}
	m := map[string]int{}
	for i := 0; i < count; i++ {
		seens = append(seens, dish)
		m[dishStr(dish)] = i
		dish = move(dish)
		dish = transpose(dish)
		if last, ok := m[dishStr(dish)]; ok {
			remain := (count - last) % (i + 1 - last)
			dish = seens[last+remain]
			break
		}
	}
	return dish
}

func dishStr(dish [][]byte) string {
	return string(bytes.Join(dish, []byte("\n")))
}

func transpose(dish [][]byte) [][]byte {
	m, n := len(dish), len(dish[0])
	res := make([][]byte, n)
	for i := 0; i < n; i++ {
		for j := m - 1; j >= 0; j-- {
			res[i] = append(res[i], dish[j][i])
		}
	}
	return res
}

func move(dish [][]byte) [][]byte {
	m, n := len(dish), len(dish[0])
	res := make([][]byte, m)
	for i := 0; i < n; i++ {
		start := 0
		for j := 0; j < m; j++ {
			res[j] = append(res[j], '.')
			switch dish[j][i] {
			case 'O':
				res[start][i] = 'O'
				start++
			case '#':
				res[j][i] = '#'
				start = j + 1
			}
		}
	}
	return res
}

func calcLoad(dish [][]byte) int {
	m, n := len(dish), len(dish[0])
	ans := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if dish[j][i] == 'O' {
				ans += (m - j)
			}
		}
	}
	return ans
}
