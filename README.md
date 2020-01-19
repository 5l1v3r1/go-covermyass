# Covermyass

[![Build status](https://github.com/SundownDEV/go-covermyass/workflows/Build/badge.svg?style=flat-square)](https://github.com/sundowndev/go-covermyass/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/sundowndev/go-covermyass)](https://goreportcard.com/report/github.com/sundowndev/go-covermyass)
[![Maintainability](https://api.codeclimate.com/v1/badges/4b59f310775d23c85617/maintainability)](https://codeclimate.com/github/sundowndev/go-covermyass/maintainability)
[![Issues](https://codeclimate.com/github/sundowndev/go-covermyass/badges/issue_count.svg)](https://codeclimate.com/github/sundowndev/go-covermyass/issues)
[![Tag](https://img.shields.io/github/tag/SundownDEV/go-covermyass.svg?style=flat)](https://github.com/sundowndev/go-covermyass/releases)

Shell script to cover your tracks on UNIX systems. Designed for pen testing "covering tracks" phase, before exiting the infected server. Or, permanently disable system logs for post-exploitation. This is a Golang implementation of the original [Covermyass shell script](https://github.com/sundowndev/covermyass).

This tool allows you to clear log files such as :

```bash
# Linux
/var/log/messages # General message and system related stuff
/var/log/auth.log # Authenication logs
/var/log/kern.log # Kernel logs
/var/log/cron.log # Crond logs
/var/log/maillog # Mail server logs
/var/log/boot.log # System boot log
/var/log/mysqld.log # MySQL database server log file
/var/log/qmail # Qmail log directory
/var/log/httpd # Apache access and error logs directory
/var/log/lighttpd # Lighttpd access and error logs directory
/var/log/secure # Authentication log
/var/log/utmp # Login records file
/var/log/wtmp # Login records file
/var/log/yum.log # Yum command log file

# macOS
/var/log/system.log # System Log
/var/log/DiagnosticMessages # Mac Analytics Data
/Library/Logs # System Application Logs
/Library/Logs/DiagnosticReports # System Reports
~/Library/Logs # User Application Logs
~/Library/Logs/DiagnosticReports # User Reports
```

## Installation

#### Temporary (global)

```bash
curl -sSL https://github.com/sundowndev/go-covermyass/releases/download/1.0.0-alpha/covermyass -o /tmp/covermyass
chmod +x /tmp/covermyass
# You can then run it using /tmp/covermyass
```

#### Permanent (local)

```bash
curl -sSL https://github.com/sundowndev/go-covermyass/releases/download/1.0.0-alpha/covermyass -o /usr/local/bin/covermyass
chmod +x /usr/local/bin/covermyass
```

You can now use the tool using the executable.

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

Simply type :

```
covermyass # you may need to use sudo if you want to clean auth logs
```

Follow the instructions :

```
Welcome to Cover my ass tool !

Select an option :

1) Clear logs for user root
2) Permenently disable auth & bash history
3) Restore settings to default
99) Exit tool

>
```

*NOTE: don't forget to exit the terminal session since the bash history is cached.*

Clear logs instantly (requires *sudo* to be efficient) :

```
sudo covermyass now
```

### Using cron job

Clear bash history every day at 5am :

```bash
0 5 * * * covermyass now >/dev/null 2>&1
```
