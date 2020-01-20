package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:     "version",
	Aliases: []string{"v", "ver"},
	Short:   "Print the version number of Covermyass",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Go-covermyass v1.0.0")
		fmt.Println("Source code: https://github.com/sundowndev/go-covermyass")
	},
}
