package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"
)

func main() {
	l := log.New(ioutil.Discard, "PBGOPY-CLIP: ", log.LstdFlags)
	if os.Getenv("PBGOPY_CLIP_DEBUG") != "" {
		l.SetOutput(os.Stdout)
	}

	if err := run(l); err != nil {
		log.Fatal(err)
	}
}

func run(logging *log.Logger) error {
	var clipboard provider = newClipboardProvider()

	addr := os.Getenv("PBGOPY_SERVER")
	if addr == "" {
		return errors.New("put the pbgopy server's address into PBGOPY_SERVER environment variable")
	}
	var server provider = newServerProvider(addr)

	logging.Printf("Server is at %s", addr)

	for {
		logging.Printf("Sleeping for 1s")
		time.Sleep(1 * time.Second)
		logging.Printf("Awake")

		if clipboard.hasNew() && clipboard.getLastTimestamp() > server.getLastTimestamp() {
			logging.Println("Clipboard has recent data")
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
			logging.Println("Successfully put clipboard data in server")

		} else if server.hasNew() && server.getLastTimestamp() >= clipboard.getLastTimestamp() {
			logging.Println("Server has recent data")
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
			logging.Println("Successfully put server data in clipboard")
		}

	}
}

func showError(err error) {
	fmt.Fprintf(os.Stderr, err.Error())
}
