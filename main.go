package main

import (
	"fmt"
)

type Space [][]int

// directions in which to move
const (
	kN = iota
	kE
	kS
	kW
	kNESW						// for mod operations
)
var dTag = map[int]string{
	kN: "N",
	kE: "E",
	kS: "S",
	kW: "W",
}
func opposite(d int) int {
	return (d + 2) % kNESW
}

func next(r int, c int, dir int) (nr int, nc int) {
	switch dir {
	case kN:
		nr = r - 1
		nc = c
	case kE:
		nr = r
		nc = c + 1
	case kS:
		nr = r + 1
		nc = c
	case kW:
		nr = r
		nc = c - 1
	default:
		nr = r
		nc = c
	}
	return
}


func isOpen(s Space, nr int, nc int) (o bool) {
	o = nr >= 0 && nc >= 0 && nr < len(s) && nc < len(s[nr]) && s[nr][nc] == 0
	return
}

func clean(s Space, r int, c int, moves chan<- string) {
	s[r][c] = 1
	for dir, tag := range dTag {
		nr, nc := next(r, c, dir)
		if isOpen(s, nr, nc) {
			moves <- fmt.Sprintf("(%d, %d) %s (%d, %d)\n", r, c, tag, nr, nc)
			clean(s, nr, nc, moves)
			back := opposite(dir)
			moves <- fmt.Sprintf("(%d, %d) %s (%d, %d)\n", nr, nc, dTag[back], r, c)
		}
	}
}

func startClean(s Space, r int, c int, moves chan<- string) {
	clean(s, r, c, moves)
	close(moves)
}

func main() {
	room := Space{
		{0, 0, 0, 1, 1},
		{0, 1, 0, 0, 0},
		{0, 1, 1, 0, 1},
	}
	moves := make(chan string)
	go startClean(room, 1, 2, moves)
	for m := range moves {
		fmt.Print(m)
	}
}
