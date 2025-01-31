package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

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
	fmt.Println("TODO ADD")
	fmt.Println(args[0])
	todoItem := todoElement{
		ID:          1,
		Description: args[0],
		CreatedAt:   time.Now().Format("2006-01-02 15:04:05"),
		IsComplete:  false,
	}
	fmt.Println(todoItem.CreatedAt)
	//Marshall into JSON
	configData, err := os.ReadFile(cfgFilePath)
	if err != nil {
		panic(err)
	}
	var todoList []todoElement

	if len(configData) == 0 {
		todoList = []todoElement{}
	} else {
		if err := json.Unmarshal(configData, &todoList); err != nil {
			panic(err)
		}
	}

	todoList = append(todoList, todoItem)

	updatedData, err := json.MarshalIndent(todoList, "", "	")
	if err != nil {
		panic(err)
	}

	if err := os.WriteFile(cfgFilePath, updatedData, 0777); err != nil {
		panic(err)
	}

	fmt.Println("Successfully added item")
}
