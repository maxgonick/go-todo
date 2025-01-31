package cmd

import (
	"fmt"
	"strconv"

	"github.com/maxgonick/go-todo/utils"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(completeCommand)
}

var completeCommand = &cobra.Command{
	Use:   "complete <taskid>",
	Short: "Mark a task as completed",
	Long:  "Mark a task as completed",
	Run:   complete,
	Args:  cobra.ExactArgs(1),
}

func complete(cmd *cobra.Command, args []string) {
	todoList := utils.MarshallToJSON(utils.CfgFilePath)
	targetId, err := strconv.Atoi(args[0])
	if err != nil {
		panic(err)
	}
	for index, element := range todoList.Elements {
		if element.ID == targetId {
			todoList.Elements[index].IsComplete = true
			utils.TodoListToDisk(todoList)
			return
		}
	}
	fmt.Println("No Element found with ID", targetId)
}
