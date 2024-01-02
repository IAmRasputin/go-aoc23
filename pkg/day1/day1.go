package day1

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/IAmRasputin/go-aoc23/internal/input"
	"github.com/IAmRasputin/go-aoc23/internal/util"
)

func parse_rune(r rune) int {
	return int(r) - '0'
}

func int_rune(i int) rune {
	return rune(i + '0')
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

func part2_decode(s string) int {
	numbers := []string{
		"zero",
		"one",
		"two",
		"three",
		"four",
		"five",
		"six",
		"seven",
		"eight",
		"nine",
	}

	digits := []rune{
		'0',
		'1',
		'2',
		'3',
		'4',
		'5',
		'6',
		'7',
		'8',
		'9',
	}

	first := 0
	second := 0

	// From the front!
toplevel:
	for i, c := range s {
		for j, n := range numbers {
			if j != 0 && (strings.HasPrefix(s[i:], n) || rune(c) == digits[j]) {
				first = j
				break toplevel
			}
		}
	}

	// From the back!
reverse_toplevel:
	for i := len(s) - 1; i >= 0; i-- {
		for j, n := range numbers {
			if j != 0 && (strings.HasPrefix(s[i:], n) || rune(s[i]) == digits[j]) {
				second = j
				break reverse_toplevel
			}
		}
	}

	return (10 * first) + second
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
			value := part2_decode(line)

			sum += value
		}
		fmt.Printf("%d\n", sum)
	}
}
