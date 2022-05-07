/*
Copyright © 2022 Atsushi

*/
package cmd

import (
	"fmt"
	"internal/repoconf"
	"internal/utils"
	"log"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add [OPTIONS]",
	Short: "Store an object to the memo.",
	Run: func(cmd *cobra.Command, args []string) {
		addTag, _ := cmd.Flags().GetBool("tag")
		description, _ := cmd.Flags().GetString("description")
		tags, _ := cmd.Flags().GetStringSlice("tags")

		if addTag == true {
			if description != "" {
				log.Fatal("Can't add description when adding a new tag")
			}
			if len(tags) != 0 {
				log.Fatal("Can't use `--tags` when adding a new tag")
			}
			addNewTag(cmd, args)
		} else {
			addNewItem(cmd)
		}
	},
}

func addNewTag(cmd *cobra.Command, args []string) {
	repo := repoconf.ReadConfig()
	allTags := append(repo.Tags, args...)
	repo.Tags = utils.GetUniqueTags(allTags)
	repoconf.WriteConfig(repo)

	fmt.Println("Added tags ", args)
}

func addNewItem(cmd *cobra.Command) {
	fmt.Println("add new item")
}

func init() {
	rootCmd.AddCommand(addCmd)
	flags := addCmd.Flags()

	// Tag
	flags.BoolP("tag", "t", false, "Indicates that you want to add a tag, not a resource")

	// Item
	flags.StringP("description", "d", "", "Add description for item to add to the repo")
	flags.StringSlice("tags", make([]string, 0), "Add tags to new item addition")
}
