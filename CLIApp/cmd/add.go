package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add to TODO list",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Add Called")
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}
