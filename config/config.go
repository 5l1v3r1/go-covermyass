package config

import "github.com/sundowndev/go-covermyass/utils"

import "os/user"

// LogFilesType is a type alias for array of strings
type LogFilesType []string

var userHomeDir = utils.GetUserHomeDir(user.Current)

// LogFiles is a list of log files related to the whole system
var LogFiles = LogFilesType{
	"/var/log/*.log",
	"/var/log/**/*.log",
	"/var/log/messages",
	"/Library/Logs/*.log",
	"/Library/Logs/**/*.log",
	"/Library/Logs/**/*.core_analytics",
	"/Library/Logs/**/*.diag",
	userHomeDir + "/*.log",
	userHomeDir + "/**/*.log",
	userHomeDir + "/.bash_history",
	userHomeDir + "/.zsh_history",
}
