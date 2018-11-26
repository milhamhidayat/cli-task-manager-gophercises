package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add task",
	Long:  `Add task in golang`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("add called")
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
