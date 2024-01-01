package day1

import (
	"fmt"
	"strconv"

	"github.com/IAmRasputin/go-aoc23/internal/input"
	"github.com/IAmRasputin/go-aoc23/internal/util"
)

func parse_rune(r rune) int {
	return int(r) - '0'
}

func part1_decode(s string) (int, error) {
	first := -1

	for _, c := range s {
		if c >= '0' && c <= '9' {
			first = parse_rune(c)
			break
		}
	}

	if first == -1 {
		// Just return 0?
		return 0, nil
	}

	second := -1

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] >= '0' && s[i] <= '9' {
			second = parse_rune(rune(s[i]))
			break
		}
	}

	if second == -1 {
		panic("Uh oh! didn't find last number!")
	}

	return strconv.Atoi(fmt.Sprintf("%d%d", first, second))
}

func part2_decode(s string) (int, error) {
	return 0, nil
}

func Solve(part int) {
	switch part {
	case 1:
		sum := 0

		for _, line := range input.InputLines(1) {
			value, err := part1_decode(line)
			util.Check(err)

			sum += value
		}

		fmt.Printf("%d\n", sum)
	case 2:
		sum := 0

		for _, line := range input.InputLines(1) {
			value, err := part2_decode(line)
			util.Check(err)

			sum += value
		}
		fmt.Printf("%d\n", 0)
	}
}
