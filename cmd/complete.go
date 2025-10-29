/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/matthosch/todo/internal/storage"
	"github.com/spf13/cobra"
)

func newCompleteCommand(ds storage.DataStore) *cobra.Command {
	completeCmd := &cobra.Command{
		Use:   "complete [id]",
		Short: "Mark a todo as complete",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			id, err := strconv.Atoi(args[0])
			if err != nil {
				fmt.Println("Invalid Selection:", args[0])
				return
			}
			if err := ds.Complete(id); err != nil {
				fmt.Println("Error completing todo:", err)
				return
			}
			fmt.Printf("Todo %d marked as complete\n", id)
		},
	}
	return completeCmd
}
