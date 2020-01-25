package cmd

import (
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/sundowndev/go-covermyass/config"
	"github.com/sundowndev/go-covermyass/utils"
)

// Mock is a function that transforms log files into symbolic links
func Mock(p *utils.FileProcessor, patterns []string) {
	files := []string{}

	for _, pattern := range patterns {
		results, _ := filepath.Glob(pattern)
		files = append(files, results...)
	}

	for _, path := range files {
		p.Register(path)
	}

	p.Proceed(func(path string) {
		// ln ...
		utils.LoggerService.Info("Create symbolic link for " + path + " to /dev/null")
	})

	utils.LoggerService.Success("Exiting.")
}

var disableCmd = &cobra.Command{
	Use:   "disable",
	Short: "Disable logs",
	Run: func(cmd *cobra.Command, args []string) {
		fileProcessor := &utils.FileProcessor{}

		Mock(fileProcessor, config.LogFiles)
	},
}
