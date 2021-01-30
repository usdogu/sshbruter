package functions

import (
	"flag"
	"fmt"
	"os"
	"time"
)

type Options struct {
	User     string
	Port     int
	Host     string
	Wordlist string
	Timeout  time.Duration
	Threads  int
	Verbose  bool
}

func Parse() *Options {
	Opt := &Options{}
	Opt.Timeout = 30 * time.Second
	flag.IntVar(&Opt.Port, "p", 22, "")
	flag.StringVar(&Opt.Host, "h", "localhost", "")
	flag.StringVar(&Opt.User, "u", "", "")
	flag.StringVar(&Opt.Wordlist, "w", "", "")
	flag.DurationVar(&Opt.Timeout, "t", Opt.Timeout, "")
	flag.IntVar(&Opt.Threads, "c", 100, "")

	flag.BoolVar(&Opt.Verbose, "v", false, "")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, `Usage:
		sshbruter [-p port] [-w wordlist.txt] [-t timeout]
				  [-c concurrent] [-u username] [-h hostname]
		
		
		
		Options:
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
		
		Example:
  		   sshbruter -w wordlist.txt -t 1m -c 1000 -u root -h localhost
		
		
		`)
	}

	flag.Parse()
	if flag.NFlag() < 2 {
		fmt.Println("Flags are lower than two exiting")
		os.Exit(1)
	}
	return Opt
}
