package main

import (
	"cmp"
	_ "embed"
	"slices"
	"strconv"
	"strings"
)

func main() {
	solve1()
	solve2()
}

//go:embed input.txt
var input string

type Card struct {
	card string
	bid  int
	kind int
}

func solve1() {
	cs := "23456789TJQKA"
	order := make(map[rune]int, len(cs))
	for i, v := range cs {
		order[v] = i
	}
	cards := make([]*Card, 0)
	for _, line := range strings.Split(string(input), "\n") {
		card, num, _ := strings.Cut(line, " ")
		kind := getKind(card)
		bid, _ := strconv.Atoi(num)
		cards = append(cards, &Card{
			card: card,
			bid:  bid,
			kind: kind,
		})

	}
	slices.SortFunc(cards, func(a, b *Card) int {
		if a.kind != b.kind {
			return cmp.Compare(a.kind, b.kind)
		}
		for i := 0; i < 5; i++ {
			if a.card[i] != b.card[i] {
				return cmp.Compare(order[rune(a.card[i])], order[rune(b.card[i])])
			}
		}
		return 0
	})
	sum := 0
	for i, card := range cards {
		sum += card.bid * (i + 1)
	}
	println(sum)
}

func getKind(s string) int {
	mp := make(map[rune]int, len(s))
	mx := 1
	for _, c := range s {
		cnt := mp[c] + 1
		mx = max(mx, cnt)
		mp[c] = cnt
	}
	switch {
	case mx == 5:
		return 7
	case mx == 4:
		return 6
	case mx == 3:
		if len(mp) == 2 {
			return 5
		} else {
			return 4
		}
	case mx == 2:
		if len(mp) == 3 {
			return 3
		} else {
			return 2
		}

	default:
		return 1
	}
}

func solve2() {
	cs := "J23456789TQKA"
	order := make(map[rune]int, len(cs))
	for i, v := range cs {
		order[v] = i
	}
	cards := make([]*Card, 0)
	for _, line := range strings.Split(string(input), "\n") {
		card, num, _ := strings.Cut(line, " ")
		kind := getKind2(card)
		bid, _ := strconv.Atoi(num)
		cards = append(cards, &Card{
			card: card,
			bid:  bid,
			kind: kind,
		})

	}
	slices.SortFunc(cards, func(a, b *Card) int {
		if a.kind != b.kind {
			return cmp.Compare(a.kind, b.kind)
		}
		for i := 0; i < 5; i++ {
			if a.card[i] != b.card[i] {
				return cmp.Compare(order[rune(a.card[i])], order[rune(b.card[i])])
			}
		}
		return 0
	})
	sum := 0
	for i, card := range cards {
		sum += card.bid * (i + 1)
	}
	println(sum)
}

func getKind2(s string) int {
	mp := make(map[rune]int, len(s))
	mx := 1
	jcnt := 0
	for _, c := range s {
		if c == 'J' {
			jcnt++
			continue
		}
		cnt := mp[c] + 1
		mx = max(mx, cnt)
		mp[c] = cnt
	}
	if len(mp) == 0 {
		mp['J'] = jcnt
		mx = jcnt
	} else {
		mx += jcnt
	}

	switch {
	case mx == 5:
		return 7
	case mx == 4:
		return 6
	case mx == 3:
		if len(mp) == 2 {
			return 5
		} else {
			return 4
		}
	case mx == 2:
		if len(mp) == 3 {
			return 3
		} else {
			return 2
		}

	default:
		return 1
	}
}
