package cmd

import (
	"fmt"
	"github.com/gissilali/today/repositories"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "today",
	Short: "A todo app for what should be done today",
	Long:  `A todo app for what should be done today, designed to be easy to use and avoid clutter that comes with complex UI`,
	Run: func(cmd *cobra.Command, args []string) {
		var tasks []repositories.Task

		db := repositories.CurrentDB()
		db.Find(&tasks)
		for _, item := range tasks {
			fmt.Println("✔", item.Task)
		}
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.today.yaml)")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.AddCommand(AddCmd)
}
