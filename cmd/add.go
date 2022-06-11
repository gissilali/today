package cmd

import (
	"github.com/gissilali/today/services"
	"github.com/spf13/cobra"
)

var AddCmd = &cobra.Command{
	Use:   "add",
	Short: "add new task",
	Run: func(cmd *cobra.Command, args []string) {
		services.InitAddTasksProgram()
	},
}
