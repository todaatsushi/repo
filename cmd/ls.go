/*
Copyright © 2022 Atsushi

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// lsCmd represents the ls command
var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List all items in the repo, optionally by tag",
	Run: func(cmd *cobra.Command, args []string) {
		listItems(cmd, args)
	},
}

func listItems(cmd *cobra.Command, tags []string) {
	fmt.Println("ls called")
}

func init() {
	rootCmd.AddCommand(lsCmd)
}
