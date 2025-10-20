package cmd

import (
	"fmt"

	"github.com/matthosch/todo/internal/storage"
	"github.com/spf13/cobra"
)

func newListCommand(ds storage.DataStore) *cobra.Command {
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
