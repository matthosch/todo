/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	todo "github.com/matthosch/todo/internal"

	"github.com/spf13/cobra"
)

var ds = todo.NewDataStore("todos.json")

func NewRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "todo",
		Short: "A simple todo applicaton",
	}
	rootCmd.AddCommand(newAddCommand(ds))

	return rootCmd
}
