package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"golang.org/x/sys/unix"
)

func ThrowError(err string) {
	fmt.Println(err)
	os.Exit(1)
}

func ReadInput(separator string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(separator)
	input, _ := reader.ReadString('\n')

	return strings.TrimRight(input, "\n")
}

func isWritable(path string) bool {
	return unix.Access(path, unix.W_OK) == nil
}
