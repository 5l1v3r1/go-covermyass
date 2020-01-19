package config

type logFiles = []string

/*
 * GeneralLogs	Log files related to the whole system
 */
var GeneralLogs = logFiles{
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
	"~/.bash_history",
	"~/.zsh_history",
}
