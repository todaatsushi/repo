/*
Copyright © 2022 Atsushi

*/
package cmd

import (
	"fmt"
	"internal/repoconf"

	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove items or tags",
	Run: func(cmd *cobra.Command, args []string) {
		tagCmd, _ := cmd.Flags().GetBool("tag")
		if tagCmd == true {
			removeTag(cmd, args)
		} else {
			removeItems(cmd, args)
		}
	},
}

func removeTag(cmd *cobra.Command, tags []string) {
	config := repoconf.ReadConfig()
	toRemove := make(map[string]bool)
	currentTags := config.Tags
	for _, tag := range tags {
		toRemove[tag] = true
	}

	newTags := make([]string, 0)
	for _, tag := range currentTags {
		if _, ok := toRemove[tag]; !ok {
			newTags = append(newTags, tag)
		}
	}
	config.Tags = newTags
	repoconf.WriteConfig(config)

	fmt.Println("Tags successfully removed.")
}

func removeItems(cmd *cobra.Command, itemNames []string) {
	fmt.Println("remove items")
}

func init() {
	rootCmd.AddCommand(removeCmd)

	flags := removeCmd.Flags()
	flags.BoolP("tag", "t", false, "Indicates that you want to add a tag, not a resource")
}
