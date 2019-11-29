package utils

import (
	"fmt"
	"os"
)

func UnmockFiles(files []string) {
	for _, path := range files {
		if _, err := os.Stat(path); os.IsNotExist(err) {
			fmt.Println("[!] " + path + " does not exists. Skipping")
			break
		}

		if isWritable(path) {
			ThrowError("[!] " + path + " is not writable! Retry using sudo.")
		}

		// rm -rf /var/log/auth.log
		// echo "" > /var/log/auth.log
		fmt.Println("[+] Disabled sending " + path + " logs to /dev/null")
	}
}

func MockFiles(files []string) {
	for _, path := range files {
		if _, err := os.Stat(path); os.IsNotExist(err) {
			fmt.Println("[!] " + path + " does not exists. Skipping")
			break
		}

		if isWritable(path) {
			ThrowError("[!] " + path + " is not writable! Retry using sudo.")
		}

		// ln ...
		fmt.Println("[+] Enabled sending " + path + " logs to /dev/null")
	}
}
