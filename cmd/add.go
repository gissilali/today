package cmd

import (
	"github.com/gissilali/today/programs/add"
	"github.com/spf13/cobra"
)

var AddCmd = &cobra.Command{
	Use:   "add",
	Short: "add new task",
	Run: func(cmd *cobra.Command, args []string) {
		add.InitAddTasksProgram()
	},
}
