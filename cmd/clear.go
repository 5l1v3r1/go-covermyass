package cmd

import (
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/sundowndev/go-covermyass/config"
	"github.com/sundowndev/go-covermyass/utils"
)

var fileProcessor = &utils.FileProcessor{}

/*
 * Clear	Function that clears log files
 */
func Clear(patterns []string) {
	files := []string{}

	for _, pattern := range patterns {
		results, _ := filepath.Glob(pattern)
		files = append(files, results...)
	}

	for _, path := range files {
		fileProcessor.Register(path)
	}

	fileProcessor.Proceed(func(path string) {
		// echo "" > /var/log/auth.log
		utils.LoggerService.Success("Cleared " + path)
	})

	utils.LoggerService.Success("Succeeded. Exiting.")
}

var clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Clear the content of log files",
	Run: func(cmd *cobra.Command, args []string) {
		Clear(config.LogFiles)
	},
}
