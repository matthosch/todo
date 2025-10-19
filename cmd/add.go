/*
Copyright © 2025 Matt Hosch mhosch24@gmail.com
*/
package cmd

import (
	"fmt"
	"strings"

	"github.com/matthosch/todo/internal/todo"
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
