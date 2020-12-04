package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	var clipboard provider = &clipboardProvider{}

	addr := os.Getenv("PBGOPY_SERVER")
	if addr == "" {
		return errors.New("put the pbgopy server's address into PBGOPY_SERVER environment variable")
	}
	var server provider = newServerProvider(addr)

	for {
		time.Sleep(1 * time.Second)

		if clipboard.hasNew() && clipboard.getLastTimestamp() > server.getLastTimestamp() {
			fmt.Println("Updating server")
			data, err := clipboard.get()
			if err != nil {
				showError(err)
				continue
			}
			err = server.put(data)
			if err != nil {
				showError(err)
				continue
			}

		} else if server.hasNew() && server.getLastTimestamp() >= clipboard.getLastTimestamp() {
			fmt.Println("Updating clipboard")
			data, err := server.get()
			if err != nil {
				showError(err)
				continue
			}
			err = clipboard.put(data)
			if err != nil {
				showError(err)
				continue
			}

		}

	}
}

func showError(err error) {
	fmt.Fprintf(os.Stderr, err.Error())
}
