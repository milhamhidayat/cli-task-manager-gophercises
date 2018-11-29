package cmd

import (
	"fmt"
	"strings"
	"task/db"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add task",
	Long:  `Add task in golang`,
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		id, err := db.CreateTask(task)
		if err != nil {
			fmt.Println("Something went wrong : ", err)
			return
		}
		fmt.Printf("Added \"%s\" to your task list with id %d.\n", task, id)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
