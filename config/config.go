package config

import "github.com/sundowndev/go-covermyass/utils"

/*
 * LogFilesType	Type for log files list
 */
type LogFilesType []string

var userHomeDir = utils.GetUserHomeDir()

/*
 * Logs	Log files related to the whole system
 */
var LogFiles = LogFilesType{
	"/var/log/*.log",
	"/var/log/**/*.log",
	"/Library/Logs/*.log",
	"/Library/Logs/**/*.log",
	// "/Library/Logs/**/*.core_analytics",
	// "/Library/Logs/**/*.diag",
	"/var/log/messages",
	userHomeDir + "/.bash_history",
	userHomeDir + "/.zsh_history",
}
