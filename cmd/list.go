package cmd

import (
	"fmt"

	"github.com/matthosch/todo/internal/todo"
	"github.com/spf13/cobra"
)

func newListCommand(ds *todo.DataStore) *cobra.Command {
	listCmd := &cobra.Command{
		Use:   "list",
		Short: "List all todos",
		Run: func(cmd *cobra.Command, args []string) {
			todos, err := ds.Load()
			if err != nil {
				fmt.Println("Error loading todos:", err)
				return
			}
			for _, t := range todos {
				fmt.Printf("- %s\n", t.Description)
			}
		},
	}
	return listCmd
}
