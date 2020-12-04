package main

import (
	"fmt"
	"os"
)

var serverAddr string

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

func run() error {
	serverAddr = os.Getenv("PBGOPY_SERVER")
	if serverAddr == "" {
		return fmt.Errorf("put the pbgopy server's address into PBGOPY_SERVER environment variable")
	}

	return nil
}
