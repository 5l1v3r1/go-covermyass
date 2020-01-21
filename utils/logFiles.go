package utils

import (
	"os"
	"path/filepath"
)

func processFile(path string, callback func(path string)) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		LoggerService.Info(path + " does not exists. Skipping")
		return
	}

	if !isWritable(path) {
		LoggerService.Error(path + " is unwritable! Please retry using sudo.")
		return
	}

	callback(path)
}

/*
 * Clear	Function that clears log files
 */
func Clear(patterns []string) {
	files := []string{}

	for _, pattern := range patterns {
		path, _ := filepath.Glob(pattern)
		files = append(files, path...)
	}

	for _, path := range files {
		processFile(path, func(path string) {
			// echo "" > /var/log/auth.log
			LoggerService.Success("Cleared " + path + " file")
		})
	}
}

/*
 * Unmock	Function that delete any existing symbolic link
 */
func Unmock(files []string) {
	for _, path := range files {
		processFile(path, func(path string) {
			// rm -rf /var/log/auth.log
			// echo "" > /var/log/auth.log
			LoggerService.Success("Delete symbolic link for " + path + " to /dev/null")
		})
	}
}

/*
 * Mock	Function that transform log files into symbolic links
 */
func Mock(files []string) {
	for _, path := range files {
		processFile(path, func(path string) {
			// ln ...
			LoggerService.Success("Create symbolic link for " + path + " to /dev/null")
		})
	}
}
