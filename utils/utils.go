package utils

import (
	"log"
	"os"
	"os/user"

	"golang.org/x/sys/unix"
)

/**
 * ThrowError	Throws an error and exit the tool.
 */
func ThrowError(err string) {
	LoggerService.Error(err)
	os.Exit(1)
}

/**
 * GetUserHomeDir	Get the current user home directory.
 */
func GetUserHomeDir() string {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	return usr.HomeDir
}

/**
 * isWritable	Function that tells you if a
 * file is writable or not.
 */
func isWritable(path string) bool {
	return unix.Access(path, unix.W_OK) == nil
}
