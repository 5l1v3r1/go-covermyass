package cmd

import (
	"github.com/spf13/cobra"
	"github.com/sundowndev/go-covermyass/config"
	"github.com/sundowndev/go-covermyass/utils"
)

var enableCmd = &cobra.Command{
	Use:   "enable",
	Short: "Enable logs",
	Run: func(cmd *cobra.Command, args []string) {
		utils.Unmock(config.GeneralLogs)
		utils.Unmock(config.HistoryLogs)
		utils.Unmock(config.AuthLogs)
	},
}
