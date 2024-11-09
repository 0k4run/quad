package piscine

import "github.com/01-edu/z01"

func QuadE(x, y int) {
	for ycount := 0; ycount < y; ycount++ {
		for xcount := 0; xcount < x; xcount++ {
			if (xcount == 0 && ycount == 0) || (xcount == x-1 && ycount == y-1) {
				z01.PrintRune('A')
			} else if (xcount == x-1 && ycount == 0) || (xcount == 0 && ycount == y-1) {
				z01.PrintRune('C')
			} else if xcount == 0 || xcount == x-1 || ycount == 0 || ycount == y-1 {
				z01.PrintRune('B')
			} else {
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune('\n')
	}
}
