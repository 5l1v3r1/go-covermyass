package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(clearCmd)
	rootCmd.AddCommand(enableCmd)
	rootCmd.AddCommand(disableCmd)
	rootCmd.AddCommand(versionCmd)
}

var rootCmd = &cobra.Command{
	Use:   "covermyass",
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
