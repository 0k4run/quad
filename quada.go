package piscine

import (
	"github.com/01-edu/z01"
)

func QuadA(x, y int) {
	for mus := 1; mus <= y; mus++ {
		for avva := 1; avva <= x; avva++ {
			if (avva == 1 && mus == 1) || (avva == x && mus == 1) || (avva == 1 && mus == y) || (avva == x && mus == y) {
				z01.PrintRune('o')
			} else if mus == 1 || mus == y {
				z01.PrintRune('-')
			} else if avva == 1 || avva == x {
				z01.PrintRune('|')
			} else {
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune('\n')
	}
}
