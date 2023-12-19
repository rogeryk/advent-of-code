package main

import (
	"bytes"
	_ "embed"
	"strings"
)

//go:embed input.txt
var input []byte
var n int

func init() {
	n = bytes.IndexByte(input, '\n')
}

func main() {
	solve()
}

var moves = [][]int{
	{0, 1},
	{1, 0},
	{0, -1},
	{-1, 0},
}

func move(x, y int, d int) (int, int, int) {
	x, y = x+moves[d][0], y+moves[d][1]
	switch input[x*(n+1)+y] {
	case 'F':
		if d == 3 {
			d = 0
		} else {
			d = 1
		}
	case 'J':
		if d == 0 {
			d = 3
		} else {
			d = 2
		}
	case 'L':
		if d == 1 {
			d = 0
		} else {
			d = 3
		}
	case '7':
		if d == 0 {
			d = 1
		} else {
			d = 2
		}
	}
	return x, y, d
}

func getStartDirection(x, y int, n int) (byte, int) {
	d := []int{}
	if y < n-1 && strings.ContainsRune("-J7", rune(input[x*(n+1)+y+1])) {
		d = append(d, 0)
	}
	if x < n-1 && strings.ContainsRune("|JL", rune(input[(x+1)*(n+1)+y])) {
		d = append(d, 1)
	}
	if y > 0 && strings.ContainsRune("-FL", rune(input[x*(n+1)+y-1])) {
		d = append(d, 2)
	}
	if x > 0 && strings.ContainsRune("|F7", rune(input[(x-1)*(n+1)+y])) {
		d = append(d, 3)
	}
	v := byte(0)
	if d[0] == 0 {
		switch d[1] {
		case 1:
			v = 'F'
		case 2:
			v = '-'
		case 3:
			v = 'L'
		}
	} else if d[0] == 1 {
		switch d[1] {
		case 2:
			v = '7'
		case 3:
			v = '|'
		}
	} else if d[0] == 2 {
		v = 'J'
	}
	return v, d[0]
}

func solve() {
	start := bytes.IndexByte(input, 'S')
	startX := start / (n + 1)
	startY := start % (n + 1)
	v, d := getStartDirection(startX, startY, n)
	input[start] = v
	steps := 0
	x, y := startX, startY
	loops := make([][]bool, n)
	for i := range loops {
		loops[i] = make([]bool, n)
	}
	for {
		loops[x][y] = true
		x, y, d = move(x, y, d)
		steps++
		if x == startX && y == startY {
			break
		}
	}
	println(steps / 2)

	enclosed := 0
	for i := 0; i < n; i++ {
		in := false
		down := false
		up := false
		for j := 0; j < n; j++ {
			if loops[i][j] {
				switch input[i*(n+1)+j] {
				case '|':
					in = !in
				case '-':
				case 'F', '7':
					down = !down
				case 'J', 'L':
					up = !up

				}
				if up && down {
					in = !in
					up = false
					down = false
				}
			}
			if in && !loops[i][j] {
				enclosed++
			}
		}
	}
	println(enclosed)
}
