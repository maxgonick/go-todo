package cmd

import (
	"time"

	"github.com/maxgonick/go-todo/utils"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(addCommand)
}

var addCommand = &cobra.Command{
	Use:   "add <description>",
	Short: "Add a task",
	Long:  "Add a task",
	Run:   add,
	Args:  cobra.ExactArgs(1),
}

func add(cmd *cobra.Command, args []string) {
	todoList := utils.MarshallToJSON(utils.CfgFilePath)

	todoItem := utils.TodoElement{
		ID:          todoList.NextId,
		Description: args[0],
		CreatedAt:   time.Now().Format("2006-01-02 15:04:05"),
		IsComplete:  false,
	}
	todoList.Elements = append(todoList.Elements, todoItem)
	todoList.NextId++
	utils.TodoListToDisk(todoList)
}
