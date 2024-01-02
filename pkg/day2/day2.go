package day2

type game struct {
	red   int
	blue  int
	green int
}

func (g *game) apply(red int, blue int, green int) {
	if red > g.red {
		g.red = red
	}

	if blue > g.blue {
		g.blue = blue
	}
}

func parse_game(game string)

func Solve(part int) {
	switch part {
	case 1:
	case 2:
	}
}
