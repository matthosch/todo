/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package main

import "github.com/matthosch/todo/cmd"

func main() {
	rootCmd := cmd.NewRootCmd()
	rootCmd.Execute()
}
