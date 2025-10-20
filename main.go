/*
Copyright © 2025 Matt Hosch mhosch24@gmail.com
*/
package main

import (
	"fmt"
	"os"

	"github.com/matthosch/todo/cmd"
	"github.com/matthosch/todo/internal/storage"
)

func main() {
	ds := storage.NewJSONStore("todos.json")
	rootCmd := cmd.NewRootCmd(ds)
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
