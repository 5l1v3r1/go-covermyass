package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/sundowndev/go-covermyass/config"
	"github.com/sundowndev/go-covermyass/utils"
)

var clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Clear logs",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("clear")

		utils.Clear(config.GeneralLogs)
		utils.Clear(config.HistoryLogs)
		utils.Clear(config.AuthLogs)
	},
}
