package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var logs_files = []string{
	"/var/log/messages",
	"/var/log/auth.log",
	"/var/log/kern.log",
	"/var/log/cron.log",
	"/var/log/maillog",
	"/var/log/boot.log",
	"/var/log/mysqld.log",
	"/var/log/qmail",
	"/var/log/httpd",
	"/var/log/lighttpd",
	"/var/log/secure",
	"/var/log/utmp",
	"/var/log/wtmp",
	"/var/log/yum.log",
	"/var/log/system.log",
	"/var/log/DiagnosticMessages",
	"/Library/Logs",
	"/Library/Logs/DiagnosticReports",
	"~/Library/Logs",
	"~/Library/Logs/DiagnosticReports",
}

func displayMenu() {
	menu := `Welcome to Cover my ass tool !

Select an option :

1) Clear logs for user $USER
2) Permenently disable auth & bash history
3) Restore settings to default
99) Exit tool
`

	fmt.Println(menu)
}

func readInput(separator string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(separator)
	input, _ := reader.ReadString('\n')

	return strings.TrimRight(input, "\n")
}

func clearLogs() {}

func enableAuthLogs()  {}
func disableAuthLogs() {}
func clearAuthLogs()   {}

func enableHistory()  {}
func disableHistory() {}
func clearHistory()   {}

func throwError(err string) {
	panic(err)
}

func main() {
	displayMenu()
	choice := readInput("> ")

	switch choice {
	case "1":
		clearLogs()
		clearHistory()
	case "2":
		disableAuthLogs()
		disableHistory()
	case "3":
		enableAuthLogs()
		enableHistory()
	case "99":
		return
	default:
		throwError("Unrecognized choice")
	}

	fmt.Println(logs_files)
}
