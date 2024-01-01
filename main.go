/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"aoc23/cmd"
	"fmt"

	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()

	if err != nil {
		panic(fmt.Errorf("%w", err))
	}

	cmd.Execute()
}
