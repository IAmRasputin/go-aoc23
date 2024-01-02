/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/IAmRasputin/go-aoc23/internal/util"
	"github.com/IAmRasputin/go-aoc23/pkg/day1"
	"github.com/spf13/cobra"
)

// solveCmd represents the solve command
var solveCmd = &cobra.Command{
	Use:   "solve",
	Short: "Return the solution for day/part",
	Long:  `Return the solution for a problem: ./aoc23 solve 1 1`,
	Run: func(_ *cobra.Command, args []string) {
		if len(args) < 2 {
			panic("need day and part")
		}

		day, err := strconv.Atoi(args[0])
		util.Check(err)
		part, err := strconv.Atoi(args[1])
		util.Check(err)

		switch day {
		case 1:
			day1.Solve(part)
		case 2:
			fmt.Println("Not implemented!")
		case 3:
			fmt.Println("Not implemented!")
		case 4:
			fmt.Println("Not implemented!")
		case 5:
			fmt.Println("Not implemented!")
		}
	},
}

func init() {
	rootCmd.AddCommand(solveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// solveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// solveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
