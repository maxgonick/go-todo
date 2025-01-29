package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var cfgDirPath string
var cfgFilePath string
var rootCmd = &cobra.Command{
	Use:   "todo",
	Short: "A simple, transparent TODO CLI tool",
	Long:  "A simple, transparent TODO CLI tool that includes add, delete, and list functionalities",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Root CMD")
	},
}

func Execute() {
	initConfig()
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func pathExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func initConfig() {
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	cfgDirPath = filepath.Join(home, ".config", "todo")
	if !pathExists(cfgDirPath) {
		os.Mkdir(cfgDirPath, 0777)
	}
	os.Chdir(cfgDirPath)

	cfgFilePath = filepath.Join(cfgDirPath, "config.json")
	if !pathExists(cfgFilePath) {
		f, err := os.Create("config.json")
		if err != nil {
			panic(err)
		}
		defer f.Close()
	}
}
