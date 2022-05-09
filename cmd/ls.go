/*
Copyright © 2022 Atsushi

*/
package cmd

import (
	"fmt"
	"internal/repoconf"
	"internal/utils"

	"github.com/spf13/cobra"
)

// lsCmd represents the ls command
var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List all items in the repo, optionally by tag",
	Run: func(cmd *cobra.Command, args []string) {
		flags := cmd.Flags()
		tags, _ := flags.GetStringSlice("tags")
		listItems(cmd, tags)
	},
}

func listItems(cmd *cobra.Command, tags []string) {
	items := repoconf.ReadConfig()
	if len(tags) == 0 {
		tags = items.Tags
	}
	for name, item := range items.Items {
		for _, tag := range item.Tags {
			tagInItem := utils.TagInItem(&tag, &tags)
			if tagInItem {
				fmt.Println("tagged:", name, item)
			}
		}
	}
}

func init() {
	rootCmd.AddCommand(lsCmd)

	flags := lsCmd.Flags()
	flags.StringSlice("tags", make([]string, 0), "")
}
