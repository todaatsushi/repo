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
		tagCmd, _ := cmd.Flags().GetBool("tag")
		description, _ := cmd.Flags().GetString("description")
		tags, _ := cmd.Flags().GetStringSlice("tags")

		if tagCmd == true {
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
	name, _ := cmd.Flags().GetString("name")
	content, _ := cmd.Flags().GetString("content")
	item := repoconf.NewMemo(content, name)

	repo := repoconf.ReadConfig()
	if _, exists := repo.Items[name]; !exists {
		repo.Items[name] = item
	} else {
		log.Fatal("add: can't add item as it already exists")
	}
	repoconf.WriteConfig(repo)
	fmt.Println("Added item:", name)
}

func init() {
	rootCmd.AddCommand(addCmd)
	flags := addCmd.Flags()

	// Tag
	flags.BoolP("tag", "t", false, "Indicates that you want to add a tag, not a resource")

	// Item
	flags.StringP("content", "c", "", "Can be a link, note or whatever you want to repo")
	flags.StringP("name", "n", "", "Name to reference the item with")
}
