package cmd

import (
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/sundowndev/go-covermyass/config"
	"github.com/sundowndev/go-covermyass/utils"
)

/*
 * Unmock	Function that delete any existing symbolic link
 */
func Unmock(p *utils.FileProcessor, patterns []string) {
	files := []string{}

	for _, pattern := range patterns {
		results, _ := filepath.Glob(pattern)
		files = append(files, results...)
	}

	for _, path := range files {
		p.Register(path)
	}

	p.Proceed(func(path string) {
		// rm -rf /var/log/auth.log
		// echo "" > /var/log/auth.log
		utils.LoggerService.Info("Delete symbolic link for " + path + " to /dev/null")
	})

	utils.LoggerService.Info("Exiting.")
}

var enableCmd = &cobra.Command{
	Use:   "enable",
	Short: "Enable logs",
	Run: func(cmd *cobra.Command, args []string) {
		fileProcessor := &utils.FileProcessor{}

		Unmock(fileProcessor, config.LogFiles)
	},
}
