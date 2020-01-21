package config

import "github.com/sundowndev/go-covermyass/utils"

type logFiles []string

var userHomeDir = utils.GetUserHomeDir()

/*
 * GeneralLogs	Log files related to the whole system
 */
var GeneralLogs = logFiles{
	"/var/log/*.log",
	"/Library/Logs/*.log",
	"/Library/Logs/*.log",
	"/Library/Logs/**/*.core_analytics",
	"/Library/Logs/**/*.diag",
	"/var/log/messages",
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
	userHomeDir + "/.bash_history",
	userHomeDir + "/.zsh_history",
}
