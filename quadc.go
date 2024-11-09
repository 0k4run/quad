package piscine

import (
	"github.com/01-edu/z01"
)

func QuadC(x, y int) {
	for ycout := 0; ycout < y; ycout++ {
		for xcout := 0; xcout < x; xcout++ {
			if (ycout == 0 && xcout == 0) || (ycout == 0 && xcout == x-1) {
				z01.PrintRune('A')
			} else if (ycout == y-1 && xcout == 0) || (ycout == y-1 && xcout == x-1) {
				z01.PrintRune('C')
			} else if xcout == 0 || xcout == x-1 || ycout == 0 || ycout == y-1 {
				z01.PrintRune('B')
			} else {
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune('\n')
	}
}
