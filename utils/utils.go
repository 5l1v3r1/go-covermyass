package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

// ThrowError throws an error and exit the tool
func ThrowError(err string) {
	LoggerService.Error(err)
	os.Exit(1)
}

// GetUserHomeDir returns the current user home directory
func GetUserHomeDir(userProfile func() (*user.User, error)) string {
	usr, err := userProfile()

	if err != nil {
		log.Fatal(err)
	}

	return usr.HomeDir
}

// ReadInput is a function that reads the user input
func ReadInput(separator string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(separator)
	input, _ := reader.ReadString('\n')

	return input
}

// PromptConfirmation is a function that ask for user confirmation
func PromptConfirmation(text string) bool {
	input := ReadInput(text + " [y/N] ")
	input = strings.ReplaceAll(input, "\n", "")

	return (input != "y" && input != "Y" && input != "yes")
}

// GetFilesFromGlobs is a function that finds files from glob patterns
func GetFilesFromGlobs(patterns []string) (files []string) {
	for _, pattern := range patterns {
		results, err := filepath.Glob(pattern)

		if err != nil {
			log.Fatal(err)
		}

		files = append(files, results...)
	}

	return files
}
