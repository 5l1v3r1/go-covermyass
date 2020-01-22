package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

/*
 * AutoConfirm Option used to skip user confirmation
 */
var AutoConfirm bool

func init() {
	// Register commands
	rootCmd.AddCommand(clearCmd)
	rootCmd.AddCommand(enableCmd)
	rootCmd.AddCommand(disableCmd)
	rootCmd.AddCommand(versionCmd)

	// Register flags
	clearCmd.PersistentFlags().BoolVarP(&AutoConfirm, "yes", "y", false, "Skip user confirmation")
	enableCmd.PersistentFlags().BoolVarP(&AutoConfirm, "yes", "y", false, "Skip user confirmation")
	disableCmd.PersistentFlags().BoolVarP(&AutoConfirm, "yes", "y", false, "Skip user confirmation")
}

var rootCmd = &cobra.Command{
	Use:   "covermyass [COMMANDS] [OPTIONS]",
	Short: "Cover your tracks on UNIX systems!",
	Long:  `Shell script to cover your tracks on UNIX systems. Designed for pen testing "covering tracks" phase, before exiting the infected server. Or, permanently disable system logs for post-exploitation.`,
}

/*
 * Execute execute root command
 */
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
