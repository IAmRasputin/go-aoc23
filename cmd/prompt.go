/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/IAmRasputin/go-aoc23/internal/prompt"
	"github.com/IAmRasputin/go-aoc23/internal/util"
	"github.com/spf13/cobra"
)

var promptCmd = &cobra.Command{
	Use:   "prompt",
	Short: "Print the day's prompt",
	Long:  `Print the day's prompt, optionally a specific part.`,
	Run: func(_ *cobra.Command, args []string) {
		if len(args) >= 2 {
			day, err := strconv.Atoi(args[0])
			util.Check(err)
			part, err := strconv.Atoi(args[1])
			util.Check(err)

			fmt.Print(prompt.PromptString(int32(day), int32(part)))
		} else if len(args) == 1 {
			day, err := strconv.Atoi(args[0])
			util.Check(err)
			path1 := fmt.Sprintf("./prompt/%d/1", day)

			if _, err = os.Stat(path1); errors.Is(err, os.ErrNotExist) {
				panic(fmt.Sprintf("No prompts found for day %d", day))
			}

			fmt.Print(prompt.PromptString(int32(day), 1))

			path2 := fmt.Sprintf("./prompt/%d/2", day)

			if _, err := os.Stat(path2); err == nil {
				fmt.Print(prompt.PromptString(int32(day), 2))
			}
		} else {
			panic("no day specified")
		}
	},
}

func init() {
	rootCmd.AddCommand(promptCmd)

}
