package cmd

import (
	"fmt"
	"strconv"

	"github.com/maxgonick/go-todo/utils"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(deleteCommand)
}

var deleteCommand = &cobra.Command{
	Use:   "delete <taskid>",
	Short: "Delete a task",
	Long:  "Delete a task",
	Run:   delete,
	Args:  cobra.ExactArgs(1),
}

func delete(cmd *cobra.Command, args []string) {
	todoList := utils.MarshallToJSON(utils.CfgFilePath)

	targetId, err := strconv.Atoi(args[0])

	if err != nil {
		panic(err)
	}

	for index, element := range todoList.Elements {
		if element.ID == targetId {
			fmt.Println("FOUND YOU", element)
			updatedList := utils.TodoList{
				Elements: append(todoList.Elements[:index], todoList.Elements[index+1:]...),
				NextId:   todoList.NextId,
			}
			utils.TodoListToDisk(updatedList)
			return
		}
	}
	fmt.Println("NO MATCHES FOUND!")

}
