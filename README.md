# SSHBRUTER

[![made with Go](https://forthebadge.com/images/badges/made-with-go.svg)](https://golang.org)
[![built with love](https://forthebadge.com/images/badges/built-with-love.svg)](#)

SSHBRUTER — A faster & simpler way to bruteforce SSH server.

## Installation

### Linux

Simply Type

```bash
wget https://raw.githubusercontent.com/usdogu/sshbruter/main/installer.sh
chmod +x ./installer.sh
./installer.sh
```

### Windows

Download latest release from [releases](https://github.com/usdogu/sshbruter/releases) page and run it.

## Usage

```bash
sshbruter [-p port] [-w wordlist.txt] [-t timeout]
				  [-c concurrent] [-u username] [-h hostname]
```
#### Options

```bash
-p port
		Port to connect to on the remote host (default 22).
-u username
		Username to connect to on the remote host.
-h host
		Host to connect.
-w wordlist
    Path to wordlist file.
-t timeout
    Connection timeout (default 30s).
-c concurrent
    Concurrency/threads level (default 100).
-v
    Verbose mode.
```

## Credits

Thanks a lot to @kitabisa for [ssb](https://github.com/kitabisa/ssb/) this project heavily inspired by ssb.
