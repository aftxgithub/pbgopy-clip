package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type serverProvider struct {
	lastServerTs int64
	addr         string
}

func (s *serverProvider) hasNew() bool {
	resp, err := http.Get(s.addr + "/lastupdated")
	if err != nil {
		log.Println(err)
		return false
	}
	serverTsBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return false
	}
	if err := resp.Body.Close(); err != nil {
		log.Println(err)
	}
	serverTs, err := strconv.ParseInt(string(serverTsBytes), 10, 64)
	if err != nil {
		log.Println(err)
		return false
	}

	if serverTs > s.lastServerTs {
		s.lastServerTs = serverTs
		return true
	}
	return false
}

func (s *serverProvider) get() (error, string) {

}

func (s *serverProvider) getTimestamp() (error, int64) {

}

func (s *serverProvider) put(data string) error {

}
