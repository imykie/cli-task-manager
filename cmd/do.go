package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"strconv"
)

var doCmd = &cobra.Command{
	Use: "do",
	Short: "Sets command to done",

	Run: func(cmd *cobra.Command, args []string){
		var ids []int
		for _, arg := range args {
			id, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Println("Failed to parse the argument:", arg)
				return
			}
			ids = append(ids, id)
		}
		fmt.Println(ids)
	},
}

func init(){
	RootCmd.AddCommand(doCmd)
}