# Covermyass

[![Build status](https://github.com/sundowndev/go-covermyass/workflows/Build/badge.svg?style=flat-square)](https://github.com/sundowndev/go-covermyass/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/sundowndev/go-covermyass)](https://goreportcard.com/report/github.com/sundowndev/go-covermyass)
[![Maintainability](https://api.codeclimate.com/v1/badges/4b59f310775d23c85617/maintainability)](https://codeclimate.com/github/sundowndev/go-covermyass/maintainability)
[![Issues](https://codeclimate.com/github/sundowndev/go-covermyass/badges/issue_count.svg)](https://codeclimate.com/github/sundowndev/go-covermyass/issues)
[![Release](https://img.shields.io/github/release/sundowndev/go-covermyass.svg?style=flat)](https://github.com/sundowndev/go-covermyass/releases)

Shell script to cover your tracks on UNIX systems. Designed for pen testing "covering tracks" phase, before exiting the infected server. Or, permanently disable system logs for post-exploitation. This is a Golang implementation of the original [Covermyass shell script](https://github.com/sundowndev/covermyass). **This tool only work for UNIX operating systems**.

>*Enter, exploit, **clean up**, exit.*

This tool allows you to clear and mock log files such as :

```bash
/var/log/**/*.log
/Library/Logs/**/*.log
/var/log/messages
~/.bash_history
~/.zsh_history
```

## Installation

This script will download the tar archive containing the binary in your current directory. Assuming your rights are limited, and need to delete it easily.

```bash
# Using curl
# Get the latest version
COVERMYASS_VERSION=$(curl -s https://api.github.com/repos/sundowndev/go-covermyass/releases/latest | grep tag_name | cut -d '"' -f 4)
curl -sSL "https://github.com/sundowndev/go-covermyass/releases/download/$COVERMYASS_VERSION/go-covermyass_$(uname -s)_$(uname -m).tar.gz" -o ./go-covermyass.tar.gz
tar xfv go-covermyass.tar.gz

# Using wget
wget "https://github.com/sundowndev/go-covermyass/releases/download/$COVERMYASS_VERSION/go-covermyass_$(uname -s)_$(uname -m).tar.gz"
tar xfv go-covermyass.tar.gz

# You can then run it using ./go-covermyass
# To uninstall: rm -rf ./go-covermyass
```

If the download fails, it probably means your platform is not currently supported. Consider [opening an issue](https://github.com/sundowndev/go-covermyass/issues/new)!

### Build from source

```shell
# Fetch project
git clone https://github.com/sundowndev/go-covermyass

# Install dependencies
go get -v -t -d ./...

# Compile source
go build -v .
```

## Usage

Keep in mind that without sudo privileges, you *might* not be able to clear system-level log files (e.g: `/var/log`). **Other users will notice someone cleaned logs**.

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

Disabling logs will create symbolic links to temporary send log messages to /dev/null.

```
covermyass disable
```

#### Enable logs

Enabling logs reset files to their initial state. **WARNING**: *file content will be deleted and permissions might be altered.*

```
covermyass enable
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
