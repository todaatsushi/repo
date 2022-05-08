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
	"github.com/spf13/pflag"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add [OPTIONS]",
	Short: "Store an object to the memo.",
	Run: func(cmd *cobra.Command, args []string) {
		flags := cmd.Flags()
		tagCmd, _ := flags.GetBool("tag")
		description, _ := flags.GetString("description")
		tags, _ := flags.GetStringSlice("tags")

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
	flags := cmd.Flags()
	name, _ := flags.GetString("name")
	content, _ := flags.GetString("content")
	description, _ := flags.GetString("description")

	item := repoconf.NewMemo(content, name)
	if description != "" {
		item.Description = description
	}

	tags, _ := flags.GetStringSlice("tags")
	if len(tags) != 0 {
		item.Tags = tags
		addNewTag(cmd, tags)
	}

	repo := repoconf.ReadConfig()
	if _, exists := repo.Items[name]; !exists {
		repo.Items[name] = item
	} else {
		log.Fatal("add: can't add item as it already exists")
	}
	repoconf.WriteConfig(repo)
	fmt.Println("Added item:", name)
}

func addFlags(flags *pflag.FlagSet) {
	// Tag
	flags.BoolP("tag", "t", false, "Indicates that you want to add a tag, not a resource")

	// Item
	flags.StringP("content", "c", "", "Can be a link, note or whatever you want to repo")
	flags.StringP("name", "n", "", "Name to reference the item with")
	flags.StringP("description", "d", "", "Optional info regarding the resource")

	flags.StringSlice("tags", make([]string, 0), "Add tags to assign to the item. Call with --create to add tags that don't exist")
}

func init() {
	rootCmd.AddCommand(addCmd)

	flags := addCmd.Flags()
	addFlags(flags)
}
