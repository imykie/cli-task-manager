package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var RootCmd = &cobra.Command{
	Use:   "task",
	Short: "Task is a CLI task manager application",
	Long:  `A simple CLI task manager that can be used to manage your day to day tasks`,
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func Must(err error) {
	if err != nil {
		fmt.Println("something went wrong:", err.Error())
		os.Exit(1)
	}
}
