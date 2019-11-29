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
			ThrowError("[!] " + path + " is not writable! Retry using sudo.")
		}

		callback(path)
	}
}

func Clear(files []string) {
	process(files, func(path string) {
		// echo "" > /var/log/auth.log
		fmt.Println("[+] Cleared " + path + " file")
	})
}

func Unmock(files []string) {
	process(files, func(path string) {
		// rm -rf /var/log/auth.log
		// echo "" > /var/log/auth.log
		fmt.Println("[+] Disabled sending " + path + " logs to /dev/null")
	})
}

func Mock(files []string) {
	process(files, func(path string) {
		// ln ...
		fmt.Println("[+] Enabled sending " + path + " logs to /dev/null")
	})
}
