package cmd

import (
	"github.com/spf13/cobra"
	"github.com/sundowndev/go-covermyass/config"
	"github.com/sundowndev/go-covermyass/utils"
)

func init() {
	// Register command
	rootCmd.AddCommand(disableCmd)

	// Register flags
	disableCmd.PersistentFlags().BoolVarP(&AutoConfirm, "yes", "y", false, "Skip user confirmation")
}

// Mock is a function that transforms log files into symbolic links
func Mock(p *utils.FileProcessor, patterns []string) {
	files := utils.GetFilesFromGlobs(patterns)

	for _, path := range files {
		p.Register(path)
	}

	p.Proceed(func(path string) {
		// ln ...
		utils.LoggerService.Info("Create symbolic link for " + path + " to /dev/null")
	})

}

var disableCmd = &cobra.Command{
	Use:   "disable",
	Short: "Disable logs",
	Run: func(cmd *cobra.Command, args []string) {
		fileProcessor := &utils.FileProcessor{}

		Mock(fileProcessor, config.LogFiles)

		utils.LoggerService.Success("Exiting.")
	},
}
