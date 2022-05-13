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
	toDisplay := make(map[string]*repoconf.Item)
	for name, item := range items.Items {
		includeItem := true
		for _, tag := range item.Tags {
			includeItem = utils.TagInItem(&tag, &tags)
		}
		if includeItem == true {
			if val := toDisplay[name]; val == nil {
				toDisplay[name] = &item
			}
		}
	}
	fmt.Println("Items:")
	for _, item := range toDisplay {
		fmt.Printf("%s    --    %v\n", item.Name, item.Tags)
	}
}

func init() {
	rootCmd.AddCommand(lsCmd)

	flags := lsCmd.Flags()
	flags.StringSliceP("tags", "t", make([]string, 0), "")
}
