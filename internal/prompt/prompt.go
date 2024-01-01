package prompt

import (
	"fmt"
	"os"
	"strings"

	"github.com/IAmRasputin/go-aoc23/internal/util"
)

func PromptString(day int32, part int32) string {
	filepath := fmt.Sprintf("./prompt/%d/%d", day, part)

	dat, err := os.ReadFile(filepath)
	util.Check(err)

	return string(dat)
}

func PromptLines(day int32, part int32) []string {
	return strings.Split(PromptString(day, part), " ")
}
