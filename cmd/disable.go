package cmd

import (
	"github.com/spf13/cobra"
	"github.com/sundowndev/go-covermyass/config"
	"github.com/sundowndev/go-covermyass/utils"
)

var disableCmd = &cobra.Command{
	Use:   "disable",
	Short: "Disable logs",
	Run: func(cmd *cobra.Command, args []string) {
		utils.Mock(config.GeneralLogs)
		utils.Mock(config.HistoryLogs)
		utils.Mock(config.AuthLogs)
	},
}
