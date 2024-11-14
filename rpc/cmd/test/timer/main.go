package main

import (
	"fmt"
)

type Nut struct {
	id int // unique identifier of the nut
}

type Bolt struct {
	id int // unique identifier of the bolt
}

func match(n Nut, b Bolt) int { // when return 0 is match; return 1 for n > b; return -1 for n < b
	if n.id == b.id {
		return 0
	} else if n.id > b.id {
		return 1
	} else {
		return -1
	}
}

func matchMap(nuts []Nut, bolts []Bolt) map[Nut]Bolt {
	matches := make(map[Nut]Bolt)
	for _, nut := range nuts {
		for _, bolt := range bolts {
			result := match(nut, bolt)
			if result == 0 {
				matches[nut] = bolt
				break
			}
		}
	}
	return matches
}

func main() {
	nuts := []Nut{{1}, {2}, {3}}
	bolts := []Bolt{{1}, {2}, {3}}

	matches := matchMap(nuts, bolts)

	fmt.Println(matches)
}
