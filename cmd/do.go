package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var doCmd = &cobra.Command{
	Use: "do",
	Short: "Sets command to done",

	Run: func(cmd *cobra.Command, args []string){
		fmt.Println("Do called!")
	},
}

func init(){
	RootCmd.AddCommand(doCmd)
}