package cmd

import (
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/sundowndev/go-covermyass/config"
	"github.com/sundowndev/go-covermyass/utils"
)

/*
 * Clear	Function that clears log files
 */
func Clear(p *utils.FileProcessor, patterns []string) {
	files := []string{}

	for _, pattern := range patterns {
		results, _ := filepath.Glob(pattern)
		files = append(files, results...)
	}

	for _, path := range files {
		p.Register(path)
	}

	p.Proceed(func(path string) {
		// echo "" > /var/log/auth.log
		utils.LoggerService.Info("Clear " + path)
	})

	utils.LoggerService.Info("Exiting.")
}

var clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Clear the content of log files",
	Run: func(cmd *cobra.Command, args []string) {
		fileProcessor := &utils.FileProcessor{}

		Clear(fileProcessor, config.LogFiles)
	},
}
