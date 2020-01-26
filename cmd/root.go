package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// AutoConfirm is an option used to skip user confirmation
var AutoConfirm bool

var rootCmd = &cobra.Command{
	Use:   "covermyass [COMMANDS] [OPTIONS]",
	Short: "Cover your tracks on UNIX systems!",
	Long:  `Shell script to cover your tracks on UNIX systems. Designed for pen testing "covering tracks" phase, before exiting the infected server. Or, permanently disable system logs for post-exploitation.`,
}

// Execute is a function that executes the root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
