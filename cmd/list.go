package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(listCommand)
}

var listCommand = &cobra.Command{
	Use:   "list",
	Short: "List all active tasks",
	Long:  "List all active tasks",
	Run:   list,
	Args:  cobra.ExactArgs(0),
}

func list(cmd *cobra.Command, args []string) {
	fmt.Println("TODO LIST")

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

	w := tabwriter.NewWriter(os.Stdout, 0, 8, 1, '\t', 0)
	defer w.Flush()
	fmt.Fprintln(w, "ID\tDescription\tCreated At\tCompleted\t")
	for _, element := range todoList {
		fmt.Fprintf(
			w,
			"%d\t%s\t%s\t%t\t\n", // Match header order and tab count
			element.ID,
			element.Description,
			element.CreatedAt, // Formatted timestamp
			element.IsComplete,
		)
	}

}
