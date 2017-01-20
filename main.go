package main

import (
	"log"

	"github.com/spf13/cobra"
)

func newCommandServer() *cobra.Command {
	pr := NewPullRequestCommand()
	var rootCmd = &cobra.Command{Use: "changelog"}
	rootCmd.AddCommand(pr)
	return rootCmd
}

func main() {
	if err := newCommandServer().Execute(); err != nil {
		log.Fatal(err)
	}
}
