package cmd

import (
	"fmt"
	"os"

	"github.com/maxgonick/go-todo/utils"
	"github.com/spf13/cobra"
)


var rootCmd = &cobra.Command{
	Use:   "todo",
	Short: "A simple, transparent TODO CLI tool",
	Long:  "A simple, transparent TODO CLI tool that includes add, delete, and list functionalities",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Root CMD")
	},
}

func Execute() {
	utils.InitConfig()
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
