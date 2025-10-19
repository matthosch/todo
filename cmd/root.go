/*
Copyright © 2025 Matt Hosch mhosch24@gmail.com
*/
package cmd

import (
	"github.com/matthosch/todo/internal/todo"

	"github.com/spf13/cobra"
)

func NewRootCmd(ds *todo.DataStore) *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "todo",
		Short: "A simple todo application",
	}
	rootCmd.AddCommand(newAddCommand(ds))

	return rootCmd
}
