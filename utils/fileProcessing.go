package utils

import (
	"os"

	"golang.org/x/sys/unix"
)

/*
 * FileProcessor	Check if a file exists
 */
type FileProcessor struct {
	Files []string
}

/*
 * FileExists	Check if a file exists
 */
func (p *FileProcessor) FileExists(path string) bool {
	_, err := os.Stat(path)
	os.IsNotExist(err)

	return !os.IsNotExist(err)
}

/*
 * FileIsWritable	Check if a file is writable
 */
func (p *FileProcessor) FileIsWritable(path string) bool {
	return unix.Access(path, unix.W_OK) == nil
}

/*
 * Register	Add a file to process
 */
func (p *FileProcessor) Register(path string) *FileProcessor {
	if !p.FileIsWritable(path) {
		LoggerService.Warn(path + " was found but is unwritable!")
		return p
	}

	p.Files = append(p.Files, path)

	return p
}

/*
 * Proceed	Use a callback for each registered file
 */
func (p *FileProcessor) Proceed(cb func(path string)) {
	if len(p.Files) == 0 {
		LoggerService.Info("No log file were found. Exiting")
		os.Exit(0)
	}

	for _, file := range p.Files {
		cb(file)
	}
}
