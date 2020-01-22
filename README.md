# Covermyass

[![Build status](https://github.com/SundownDEV/go-covermyass/workflows/Build/badge.svg?style=flat-square)](https://github.com/sundowndev/go-covermyass/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/sundowndev/go-covermyass)](https://goreportcard.com/report/github.com/sundowndev/go-covermyass)
[![Maintainability](https://api.codeclimate.com/v1/badges/4b59f310775d23c85617/maintainability)](https://codeclimate.com/github/sundowndev/go-covermyass/maintainability)
[![Issues](https://codeclimate.com/github/sundowndev/go-covermyass/badges/issue_count.svg)](https://codeclimate.com/github/sundowndev/go-covermyass/issues)
[![Tag](https://img.shields.io/github/tag/SundownDEV/go-covermyass.svg?style=flat)](https://github.com/sundowndev/go-covermyass/releases)

Shell script to cover your tracks on UNIX systems. Designed for pen testing "covering tracks" phase, before exiting the infected server. Or, permanently disable system logs for post-exploitation. This is a Golang implementation of the original [Covermyass shell script](https://github.com/sundowndev/covermyass). **This tool only work for UNIX operating systems**.

>*Enter, exploit, **clean up**, exit.*

This tool allows you to clear log files such as :

```bash
/var/log/*.log
/Library/Logs/*.log
/Library/Logs/*.log
/var/log/messages
~/.bash_history
~/.zsh_history
```

## Installation

This script will download the binary in your current directory. Assuming your rights are limited, and need to delete it easily.

```bash
curl -sSL https://github.com/sundowndev/go-covermyass/releases/download/1.0.0-alpha/covermyass -o ./covermyass
chmod +x ./covermyass
# You can then run it using ./covermyass
# To uninstall: rm -rf ./covermyass
```

Keep in mind that without sudo privileges, you *might* be unable to clear system-level log files (`/var/log`).

## Build from source

```shell
# Fetch project
git clone https://github.com/sundowndev/go-covermyass

# Install dependencies
go get -v -t -d ./...

# Compile source
go build -v .
```

## Usage

```
covermyass # you may need to use sudo if you want to clean system-level logs
```

Use `--help` option to display help message :

```
Shell script to cover your tracks on UNIX systems. Designed for pen testing "covering tracks" phase, before exiting the infected server. Or, permanently disable system logs for post-exploitation.

Usage:
  covermyass [command]

Available Commands:
  clear       Clear the content of log files
  disable     Disable logs
  enable      Enable logs
  help        Help about any command
  version     Print the version number of Covermyass

Flags:
  -h, --help   help for covermyass

Use "covermyass [command] --help" for more information about a command.
```

#### Clear logs

Clearing logs will simply empty them.

```
covermyass clear
```

#### Disable logs

Disabling logs will create symbolic link to temporary send log messages to /dev/null.

```
covermyass disable
```

#### Enable logs

Enabling logs reset files to their initial state. **WARNING**: content will be deleted and permissions might not be the same as before.

```
covermyass disable
```

*NOTE: don't forget to exit the terminal session since the bash history is cached.*

### Auto confirm

Skip the user confirmation using the `-y` flag.

```
covermyass clear -y
```

### Using cron job

Clear bash history every day at 5am :

```bash
0 5 * * * covermyass clear -y >/dev/null 2>&1
```
