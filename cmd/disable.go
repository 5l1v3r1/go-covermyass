package cmd

import (
	"os"

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
	for _, path := range utils.GetFilesFromGlobs(patterns) {
		p.Register(path)
	}

	if len(p.Files) == 0 {
		utils.LoggerService.Info("No log file were found. Exiting.")
		os.Exit(0)
	}

	for _, path := range p.Files {
		utils.LoggerService.Info(path)
	}

	if !AutoConfirm {
		confirm := utils.PromptConfirmation("The following files will be replaced by symbolic links to /dev/null, proceed?")

		if confirm {
			utils.LoggerService.Info("Cancelling.")
			os.Exit(0)
		}
	}

	p.Proceed(func(path string) {
		err := os.Symlink("/dev/null", path)

		if err != nil {
			utils.ThrowError(err.Error())
		}
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
