package cmd

import (
	"github.com/spf13/cobra"
	"github.com/sundowndev/go-covermyass/config"
	"github.com/sundowndev/go-covermyass/utils"
)

func init() {
	// Register command
	rootCmd.AddCommand(clearCmd)

	// Register flags
	clearCmd.PersistentFlags().BoolVarP(&AutoConfirm, "yes", "y", false, "Skip user confirmation")
}

// Clear is a function that clears log files
func Clear(p *utils.FileProcessor, patterns []string) {
	files := utils.GetFilesFromGlobs(patterns)

	for _, path := range files {
		p.Register(path)
	}

	p.Proceed(func(path string) {
		// echo "" > /var/log/auth.log
		utils.LoggerService.Info("Clear " + path)
	})

}

var clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Clear the content of log files",
	Run: func(cmd *cobra.Command, args []string) {
		fileProcessor := &utils.FileProcessor{}

		Clear(fileProcessor, config.LogFiles)

		utils.LoggerService.Success("Exiting.")
	},
}
