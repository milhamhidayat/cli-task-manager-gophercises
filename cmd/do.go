package cmd

import (
	"fmt"
	"strconv"
	"task/db"

	"github.com/spf13/cobra"
)

// doCmd represents the do command
var doCmd = &cobra.Command{
	Use:   "do",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		ids := []int{}
		for _, arg := range args {
			id, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Println("Failed to parse the argument", arg)
			} else {
				ids = append(ids, id)
			}
		}
		tasks, err := db.GetAllTask()
		if err != nil {
			fmt.Println("Something went wrong : ", err)
			return
		}
		for _, id := range ids {
			if id <= 0 || id > len(tasks) {
				fmt.Println("Invalid task number : ", id)
				continue
			}
			task := tasks[id-1]
			err := db.DeleteTask(task.Key)
			if err != nil {
				fmt.Printf("Failed to mark \"%d\" as completed. Error : %s\n", id, err)
			} else {
				fmt.Printf("Marked \"%d\" as completed \n", id)
			}
		}

	},
}

func init() {
	rootCmd.AddCommand(doCmd)
}
