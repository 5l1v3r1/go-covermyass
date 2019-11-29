package utils

import (
	"fmt"
	"os"
)

func ClearFiles(files []string) {
	for _, path := range files {
		if _, err := os.Stat(path); os.IsNotExist(err) {
			fmt.Println("[!] " + path + " does not exists. Skipping")
			break
		}

		if isWritable(path) {
			ThrowError("[!] " + path + " is not writable! Retry using sudo.")
		}

		// echo "" > /var/log/auth.log
		fmt.Println("[+] Cleared " + path + " file")
	}
}
