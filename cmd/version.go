package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/sundowndev/go-covermyass/config"
)

var versionCmd = &cobra.Command{
	Use:     "version",
	Aliases: []string{"v", "ver"},
	Short:   "Print the version number of Covermyass",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Go-covermyass %v\n", config.Version)
		fmt.Println("Source code: https://github.com/sundowndev/go-covermyass")
	},
}
