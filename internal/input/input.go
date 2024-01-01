package input

import (
	"fmt"
	"os"
	"strings"

	"github.com/IAmRasputin/go-aoc23/internal/util"
)

func InputString(day int32) string {
	input, err := os.ReadFile(fmt.Sprintf("./input/%d", day))
	util.Check(err)
	return string(input)
}

func InputLines(day int32) []string {
	return strings.Split(InputString(day), "\n")
}

func InputWords(day int32) [][]string {
	var lines = InputLines(day)
	num_lines := len(lines)

	split_lines := make([][]string, num_lines)

	for i, line := range lines {
		split_lines[i] = strings.Split(line, " ")
	}

	return split_lines
}
