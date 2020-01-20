package utils

import (
	"fmt"
	"os"
)

func process(files []string, callback func(path string)) {
	for _, path := range files {
		if _, err := os.Stat(path); os.IsNotExist(err) {
			fmt.Println("[!] " + path + " does not exists. Skipping")
			break
		}

		if isWritable(path) {
			ThrowError("[!] " + path + " is unwritable! Please retry using sudo.")
		}

		callback(path)
	}
}

/*
 * Clear	Function that clears log files
 */
func Clear(files []string) {
	process(files, func(path string) {
		// echo "" > /var/log/auth.log
		fmt.Println("[+] Cleared " + path + " file")
	})
}

/*
 * Unmock	Function that delete any existing symbolic link
 */
func Unmock(files []string) {
	process(files, func(path string) {
		// rm -rf /var/log/auth.log
		// echo "" > /var/log/auth.log
		fmt.Println("[+] Delete symbolic link for " + path + " to /dev/null")
	})
}

/*
 * Mock	Function that transform log files into symbolic links
 */
func Mock(files []string) {
	process(files, func(path string) {
		// ln ...
		fmt.Println("[+] Create symbolic link for " + path + " to /dev/null")
	})
}
