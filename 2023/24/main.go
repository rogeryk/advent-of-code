package main

import (
	_ "embed"
	"strings"

	"github.com/spf13/cast"
)

//go:embed input.txt
var input string

func main() {
	solve1()
	solve2()
}

func solve2() {
	var hailstones []HailStone
	velX := map[int][]int{}
	velY := map[int][]int{}
	velZ := map[int][]int{}
	for _, line := range strings.Split(input, "\n") {
		position, velocity, _ := strings.Cut(line, " @ ")
		var pos, vel []int
		for _, v := range strings.Split(position, ", ") {
			pos = append(pos, cast.ToInt(strings.TrimSpace(v)))
		}
		for _, v := range strings.Split(velocity, ", ") {
			vel = append(vel, cast.ToInt(strings.TrimSpace(v)))
		}
		velX[vel[0]] = append(velX[vel[0]], pos[0])
		velY[vel[1]] = append(velY[vel[1]], pos[1])
		velZ[vel[2]] = append(velZ[vel[2]], pos[2])
		hailstones = append(hailstones, HailStone{pos, vel})
	}
	vx := getVelocity(velX)
	vy := getVelocity(velY)
	vz := getVelocity(velZ)

	for i, h1 := range hailstones {
		for _, h2 := range hailstones[i+1:] {
			a1, b1, c1 := float64(h1.vel[1]-vy), float64(-h1.vel[0]+vx), float64(h1.pos[0]*(h1.vel[1]-vy)-h1.pos[1]*(h1.vel[0]-vx))
			a2, b2, c2 := float64(h2.vel[1]-vy), float64(-h2.vel[0]+vx), float64(h2.pos[0]*(h2.vel[1]-vy)-h2.pos[1]*(h2.vel[0]-vx))
			if a1*b2 == a2*b1 {
				continue
			}
			px := (c1*b2 - c2*b1) / (a1*b2 - a2*b1)
			py := (c1*a2 - c2*a1) / (b1*a2 - b2*a1)

			a1, b1, c1 = float64(h1.vel[2]-vz), float64(-h1.vel[0]+vx), float64(h1.pos[0]*(h1.vel[2]-vz)-h1.pos[2]*(h1.vel[0]-vx))
			a2, b2, c2 = float64(h2.vel[2]-vz), float64(-h2.vel[0]+vx), float64(h2.pos[0]*(h2.vel[2]-vz)-h2.pos[2]*(h2.vel[0]-vx))

			px2 := (c1*b2 - c2*b1) / (a1*b2 - a2*b1)
			pz := (c1*a2 - c2*a1) / (b1*a2 - b2*a1)
			if px == px2 {
				ans := int(px) + int(py) + int(pz)
				println(ans)
				return
			}
		}
	}
}

func getVelocity(vel map[int][]int) int {
	cnt := 0
	res := 0
	for i := -1000; i <= 1000; i++ {
		if i == 0 {
			continue
		}
		flag := true
		for v, values := range vel {
			if len(values) < 2 {
				continue
			}
			if i == int(v) || int(values[0]-values[1])%int(i-int(v)) != 0 {
				flag = false
				break
			}
		}
		if flag {
			res = i
			cnt++
		}
	}
	if cnt != 1 {
		panic("can't get velocity")
	}
	return res
}

type HailStone struct {
	pos []int
	vel []int
}

func solve1() {
	var hailstones []HailStone
	for _, line := range strings.Split(input, "\n") {
		position, velocity, _ := strings.Cut(line, " @ ")
		var pos, vel []int
		for _, v := range strings.Split(position, ", ") {
			pos = append(pos, cast.ToInt(strings.TrimSpace(v)))
		}
		for _, v := range strings.Split(velocity, ", ") {
			vel = append(vel, cast.ToInt(strings.TrimSpace(v)))
		}
		hailstones = append(hailstones, HailStone{
			pos: pos,
			vel: vel,
		})
	}

	ans := 0
	for i, h1 := range hailstones {
		for _, h2 := range hailstones[i+1:] {
			a1, b1, c1 := float64(h1.vel[1]), float64(-h1.vel[0]), float64(h1.pos[0]*h1.vel[1]-h1.pos[1]*h1.vel[0])
			a2, b2, c2 := float64(h2.vel[1]), float64(-h2.vel[0]), float64(h2.pos[0]*h2.vel[1]-h2.pos[1]*h2.vel[0])
			if a1*b2 == a2*b1 {
				continue
			}
			x := float64(c1*b2-c2*b1) / float64(a1*b2-a2*b1)
			y := float64(c1*a2-c2*a1) / float64(b1*a2-b2*a1)
			if x < 200000000000000 || x > 400000000000000 || y < 200000000000000 || y > 400000000000000 {
				continue
			}
			if (x-float64(h1.pos[0]))*float64(h1.vel[0]) < 0 || (x-float64(h2.pos[0]))*float64(h2.vel[0]) < 0 ||
				(y-float64(h1.pos[1]))*float64(h1.vel[1]) < 0 || (y-float64(h2.pos[1]))*float64(h2.vel[1]) < 0 {
				continue
			}
			ans++
		}
	}
	println(ans)
}
