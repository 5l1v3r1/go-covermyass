package cmd

import (
	"github.com/spf13/cobra"
	"github.com/sundowndev/go-covermyass/config"
	"github.com/sundowndev/go-covermyass/utils"
)

func init() {
	// Register command
	rootCmd.AddCommand(enableCmd)

	// Register flags
	enableCmd.PersistentFlags().BoolVarP(&AutoConfirm, "yes", "y", false, "Skip user confirmation")
}

// Unmock is a function that delete any existing symbolic link
func Unmock(p *utils.FileProcessor, patterns []string) {
	files := utils.GetFilesFromGlobs(patterns)

	for _, path := range files {
		p.Register(path)
	}

	p.Proceed(func(path string) {
		// rm -rf /var/log/auth.log
		// echo "" > /var/log/auth.log
		utils.LoggerService.Info("Delete symbolic link for " + path + " to /dev/null")
	})

}

var enableCmd = &cobra.Command{
	Use:   "enable",
	Short: "Enable logs",
	Run: func(cmd *cobra.Command, args []string) {
		fileProcessor := &utils.FileProcessor{}

		Unmock(fileProcessor, config.LogFiles)

		utils.LoggerService.Success("Exiting.")
	},
}
