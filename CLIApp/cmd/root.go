package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var RootCmd = &cobra.Command{
	Use:   "task",
	Short: "Task is cli manager",
	Long:  `Task is cli manager for managing tasks ...............`,
	//Run: func(cmd *cobra.Command, args []string) {
	//	// Do Stuff Here
	//},
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
