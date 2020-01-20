package config

import "github.com/sundowndev/go-covermyass/utils"

type logFiles = []string

/*
 * GeneralLogs	Log files related to the whole system
 */
var GeneralLogs = logFiles{
	"/var/log/messages/*.log",
	"/var/log/kern.log",
	"/var/log/cron.log",
	"/var/log/maillog/*.log",
	"/var/log/boot.log",
	"/var/log/mysqld.log",
	"/var/log/qmail/*.log",
	"/var/log/httpd/*.log",
	"/var/log/lighttpd/*.log",
	"/var/log/secure/*.log",
	"/var/log/utmp/*.log",
	"/var/log/wtmp/*.log",
	"/var/log/yum.log",
	"/var/log/system.log",
	"/var/log/DiagnosticMessages/*.log",
	"/Library/Logs/*.log",
	"/Library/Logs/DiagnosticReports/*.diag",
}

/*
 * AuthLogs	Log files related to system authentication
 */
var AuthLogs = logFiles{
	"/var/log/auth.log",
}

/*
 * HistoryLogs	Log files related to user history
 */
var HistoryLogs = logFiles{
	utils.GetUserHomeDir() + "/.bash_history",
	utils.GetUserHomeDir() + "/.zsh_history",
}
