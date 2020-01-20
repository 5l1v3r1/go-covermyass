package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/sundowndev/go-covermyass/config"
	"github.com/sundowndev/go-covermyass/utils"
)

var enableCmd = &cobra.Command{
	Use:   "enable",
	Short: "Enable logs",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("enable")

		utils.Unmock(config.GeneralLogs)
		utils.Unmock(config.HistoryLogs)
		utils.Unmock(config.AuthLogs)
	},
}
