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
	filename, exists := repoconf.GetLibPath()

	if exists == false {
		_ = os.MkdirAll(root, os.ModePerm)
		file, err := os.Create(filename)
		if err != nil {
			log.Fatalf("init: couldn't create %s", filename)
			return
		}
		defer file.Close()
		if err != nil {
			log.Fatal("init: couldn't create startup script")
		}
	} else {
		log.Fatalf("Could not create the root directory at '%s'. You may have to create the dir or check that nothing with the same name exists already.", root)
	}
	fmt.Printf("Init successful. New `repo` dir made at %s.", root)
}

func init() {
	rootCmd.AddCommand(initCmd)
}
