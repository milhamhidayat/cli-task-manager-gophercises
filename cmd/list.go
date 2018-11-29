package cmd

import (
	"fmt"
	"os"
	"task/db"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List of task",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := db.GetAllTask()
		if err != nil {
			fmt.Println("Something went wrong", err)
			os.Exit(1)
		}
		if len(tasks) == 0 {
			fmt.Println("You have no tasks")
			return
		}
		fmt.Println("You have the following tasks")
		for i, task := range tasks {
			fmt.Printf("%d. %s\n", i+1, task.Value)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
