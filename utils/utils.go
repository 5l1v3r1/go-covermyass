package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/user"
)

// ThrowError throws an error and exit the tool
func ThrowError(err string) {
	LoggerService.Error(err)
	os.Exit(1)
}

// GetUserHomeDir returns the current user home directory
func GetUserHomeDir() string {
	usr, err := user.Current()
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
