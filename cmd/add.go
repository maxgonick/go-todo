package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
)

type todoElement struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
	IsComplete  bool      `json:"isComplete"`
}

func init() {
	rootCmd.AddCommand(addCommand)
}

var addCommand = &cobra.Command{
	Use:   "add",
	Short: "Add a task",
	Long:  "Add a task",
	Run:   add,
}

func add(cmd *cobra.Command, args []string) {
	fmt.Println("TODO ADD")

	todoItem := todoElement{
		ID:          1,
		Description: "foo",
		CreatedAt:   time.Now(),
		IsComplete:  false,
	}

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
