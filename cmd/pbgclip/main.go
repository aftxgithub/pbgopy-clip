package main

import (
	"log"

	pbgclip "github.com/thealamu/pbgopy-clip"
)

func main() {
	if err := pbgclip.Run(); err != nil {
		log.Fatal(err)
	}
}
