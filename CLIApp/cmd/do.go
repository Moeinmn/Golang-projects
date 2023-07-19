package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"strconv"
)

var doneCmd = &cobra.Command{
	Use:   "do",
	Short: "Do the job from list",
	Run: func(cmd *cobra.Command, args []string) {
		var ids []int
		for _, arg := range args {
			intId, err := strconv.Atoi(arg)
			if err != nil {
				log.Panicf("error with #%v", arg)
			} else {
				ids = append(ids, intId)
			}
		}
		log.Println(ids)
	},
}

func init() {
	RootCmd.AddCommand(doneCmd)
}
