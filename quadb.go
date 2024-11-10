package piscine

import (
	"github.com/01-edu/z01"
)

func QuadB(x, y int) {
	for ligne := 1; ligne <= y; ligne++ {
		for colonne := 1; colonne <= x; colonne++ {
			if ligne == 1 && colonne == 1 {
				z01.PrintRune('/')
			} else if (ligne == 1 && colonne == x) || (ligne == y && colonne == 1) {
				z01.PrintRune('\\')
			} else if ligne == y && colonne == x {
				z01.PrintRune('/')
			} else if ligne == 1 || ligne == y || colonne == 1 || colonne == x {
				z01.PrintRune('*')
			} else {
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune('\n')
	}
}
