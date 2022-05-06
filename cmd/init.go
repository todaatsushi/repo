/*
Copyright © 2022 Atsushi

*/
package cmd

import (
	"fmt"
	"internal/repoconf"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Create the initial env to use repo",
	Run: func(cmd *cobra.Command, args []string) {
		initConfig(cmd)
	},
}

func initConfig(cmd *cobra.Command) {
	root, _ := repoconf.GetRootPath()
	_, exists := repoconf.GetLibPath()

	emptyConfig := repoconf.NewConfig()
	emptyConfig.Tags = make([]string, 0)

	if exists == false {
		_ = os.MkdirAll(root, os.ModePerm)
		err := repoconf.WriteConfig(emptyConfig)
		if err != nil {
			log.Fatal("init: couldn't create startup script")
		}
	} else {
		force, _ := cmd.Flags().GetBool("force")
		if force == true {
			log.Fatalf("Implement delete file at %s", root)
		} else {
			log.Fatalf("Could not create the root directory at '%s'. You may have to create the dir or check that nothing with the same name exists already.", root)
		}
	}
	fmt.Printf("Init successful. New `repo` dir made at %s.", root)
}

func init() {
	rootCmd.AddCommand(initCmd)
	initCmd.Flags().BoolP("force", "f", false, "Delete and restart a new repo.")
}
