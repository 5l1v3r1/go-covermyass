package utils

import (
	"os"

	"golang.org/x/sys/unix"
)

// FileProcessor is a type
type FileProcessor struct {
	Files []string
}

// FileExists is a FileProcessor method to check if a file exists
func (p *FileProcessor) FileExists(path string) bool {
	_, err := os.Stat(path)
	os.IsNotExist(err)

	return !os.IsNotExist(err)
}

// FileIsWritable is a FileProcessor method to check if a file is writable
func (p *FileProcessor) FileIsWritable(path string) bool {
	return unix.Access(path, unix.W_OK) == nil
}

// Register is a FileProcessor method to add a file to process
func (p *FileProcessor) Register(path string) *FileProcessor {
	if !p.FileIsWritable(path) {
		LoggerService.Warn(path + " was found but is unwritable!")
		return p
	}

	p.Files = append(p.Files, path)

	return p
}

// Proceed is a FileProcessor method to call an
// anonymous function for each registered file
func (p *FileProcessor) Proceed(cb func(path string)) {
	for _, file := range p.Files {
		cb(file)
	}
}
