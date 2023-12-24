package main

import (
	_ "embed"
	"slices"
	"strings"

	"github.com/zyedidia/generic/queue"
)

//go:embed input.txt
var input string

func main() {
	solve1()
	solve2()
}

type pulse struct {
	from  string
	to    string
	state int
}

type module struct {
	Name    string
	Outputs []string
	Inputs  map[string]int
	Type    string
	State   int
}

func (m *module) accept(p *pulse) []*pulse {
	out := p.state
	switch m.Type {
	case "%":
		if p.state != 0 {
			return nil
		}
		if m.State == 0 {
			out = 1
			m.State = 1
		} else {
			out = 0
			m.State = 0
		}
	case "&":
		m.Inputs[p.from] = p.state
		out = 0
		for _, v := range m.Inputs {
			if v == 0 {
				out = 1
			}
		}
	}
	var res []*pulse
	for _, name := range m.Outputs {
		res = append(res, &pulse{from: m.Name, to: name, state: out})
	}
	return res
}

func solve1() {
	modules := map[string]*module{}
	for _, line := range strings.Split(input, "\n") {
		in, out, _ := strings.Cut(line, " -> ")
		name := in
		t := ""
		if in[0] == '&' || in[0] == '%' {
			name = in[1:]
			t = in[0:1]
		}
		outputs := strings.Split(out, ", ")
		modules[name] = &module{
			Name:    name,
			Outputs: outputs,
			Inputs:  map[string]int{},
			Type:    t,
			State:   0,
		}
	}
	for _, m := range modules {
		for _, name := range m.Outputs {
			if out, ok := modules[name]; ok {
				out.Inputs[m.Name] = 0
			}
		}
	}
	low, high := 0, 0
	for i := 0; i < 1000; i++ {
		low++
		q := queue.New[*pulse]()
		q.Enqueue(&pulse{from: "button", to: "broadcaster", state: 0})
		for !q.Empty() {
			p := q.Dequeue()
			if _, ok := modules[p.to]; !ok {
				continue
			}
			outputs := modules[p.to].accept(p)
			for _, out := range outputs {
				if out.state == 0 {
					low++
				} else {
					high++
				}
				q.Enqueue(out)
			}
		}
	}
	println(low * high)
}

func solve2() {
	modules := map[string]*module{}
	rxSource := ""
	for _, line := range strings.Split(input, "\n") {
		in, out, _ := strings.Cut(line, " -> ")
		name := in
		t := ""
		if in[0] == '&' || in[0] == '%' {
			name = in[1:]
			t = in[0:1]
		}
		outputs := strings.Split(out, ", ")
		if slices.Contains(outputs, "rx") {
			rxSource = name
		}
		modules[name] = &module{
			Name:    name,
			Outputs: outputs,
			Inputs:  map[string]int{},
			Type:    t,
			State:   0,
		}
	}
	for _, m := range modules {
		for _, name := range m.Outputs {
			if out, ok := modules[name]; ok {
				out.Inputs[m.Name] = 0
			}
		}
	}
	triggerHigh := map[string]int{}
	for i := 0; ; i++ {
		q := queue.New[*pulse]()
		q.Enqueue(&pulse{from: "button", to: "broadcaster", state: 0})
		for !q.Empty() {
			p := q.Dequeue()
			if _, ok := modules[p.to]; !ok {
				continue
			}
			outputs := modules[p.to].accept(p)
			for _, out := range outputs {
				if out.to == rxSource && out.state == 1 {
					if _, ok := triggerHigh[out.from]; !ok {
						triggerHigh[out.from] = i + 1
					}
				}
				q.Enqueue(out)
			}
		}
		if len(triggerHigh) == len(modules[rxSource].Inputs) {
			break
		}
	}
	ans := 1
	for _, v := range triggerHigh {
		ans = lcm(ans, v)
	}
	println(ans)
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
