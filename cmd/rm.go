package cmd

import (
	"fmt"
	"strconv"
	"task/db"

	"github.com/spf13/cobra"
)

// rmCmd represents the rm command
var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove a task",
	Run: func(cmd *cobra.Command, args []string) {
		ids := []int{}
		for _, arg := range args {
			id, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Println("Failed to parse the argument")
			} else {
				ids = append(ids, id)
			}
		}
		tasks, err := db.GetAllTask()
		if err != nil {
			fmt.Println("Something went wrong : ", tasks)
		}
		for _, id := range ids {
			if id <= 0 || id > len(tasks) {
				fmt.Println("Invalid task number : ", id)
				continue
			}
			task := tasks[id-1]
			err := db.DeleteTask(task.Key)
			if err != nil {
				fmt.Printf("Failed to delete \"%d\". Error : %s\n", id, err)
			} else {
				fmt.Printf("Success to delete \"%d\" \n", id)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(rmCmd)
}
