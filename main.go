package main

import (
	"fmt"
	"os"

	"covermyass/src/utils"
)

var generalLogs = []string{
	"/var/log/messages",
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

var authLogs = []string{
	"/var/log/auth.log",
}

var historyLogs = []string{
	"~/.bash_history",
	"~/.zsh_history",
}

func displayMenu() {
	menu := `Welcome to Cover My Ass !

Select an option :

1) Clear your own logs
2) Permenently disable auth & bash history
3) Restore settings to default
99) Exit tool
`

	fmt.Println(menu)
}

func main() {
	displayMenu()
	run()

	os.Exit(0)
}

func run() {
	choice := utils.ReadInput("> ")

	switch choice {
	case "1":
		utils.ClearFiles(generalLogs)
		utils.ClearFiles(historyLogs)
	case "2":
		utils.MockFiles(authLogs)
		utils.MockFiles(historyLogs)
	case "3":
		utils.UnmockFiles(authLogs)
		utils.UnmockFiles(historyLogs)
	case "99":
		os.Exit(0)
	default:
		fmt.Println("[!] Unreconized option.")
		run()
	}
}
