package cmd

import "github.com/spf13/cobra"

var RootCmd = &cobra.Command{
	Use: "task",
	Short: "Task is a CLI task manager application",
	Long: `A simple CLI task manager that can be used to manage your day to day tasks`,

}
