/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strings"

	todo "github.com/matthosch/todo/internal"
	"github.com/spf13/cobra"
)

func newAddCommand(ds *todo.DataStore) *cobra.Command {
	addCmd := &cobra.Command{
		Use:   "add",
		Short: "Add a new todo",
		Run: func(cmd *cobra.Command, args []string) {
			description := strings.Join(args, " ")
			todo := todo.NewTodo(description)
			if err := ds.Add(*todo); err != nil {
				fmt.Println("Error adding todo:", err)
				return
			}
			fmt.Println("Todo added:", todo.Description)
		},
	}

	return addCmd
}
