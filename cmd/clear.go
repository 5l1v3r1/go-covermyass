package cmd

import (
	"io/ioutil"
	"os"

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

	if confirm := utils.PromptConfirmation("The following files will be wiped, proceed?"); confirm {
		utils.LoggerService.Info("Cancelling.")
		os.Exit(0)
	}

	p.Proceed(func(path string) {
		err := ioutil.WriteFile(path, []byte(""), 0644)

		if err != nil {
			utils.ThrowError(err.Error())
		}
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
