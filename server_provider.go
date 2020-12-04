package main

import (
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

type serverProvider struct {
	lastServerTs int64
	addr         string
}

func (s *serverProvider) hasNew() bool {
	lastTsFromServer, err := s.getLastTimestamp()
	if err != nil {
		return false
	}
	if lastTsFromServer > s.lastServerTs {
		s.lastServerTs = lastTsFromServer
		return true
	}
	return false
}

func (s *serverProvider) get() (string, error) {
	resp, err := http.Get(s.addr + "/")
	if err != nil {
		return "", err
	}
	if !isPlainText(resp.Header.Get("Content-Type")) {
		return "", errors.New("server content is not plain text")
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (s *serverProvider) getLastTimestamp() (int64, error) {
	resp, err := http.Get(s.addr + "/lastupdated")
	if err != nil {
		return 0, err
	}
	serverTsBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}
	if err := resp.Body.Close(); err != nil {
		return 0, err
	}
	serverTs, err := strconv.ParseInt(string(serverTsBytes), 10, 64)
	if err != nil {
		return 0, err
	}
	return serverTs, nil
}

func (s *serverProvider) put(data string) error {
	req, err := http.NewRequest(http.MethodPut, s.addr+"/", strings.NewReader(data))
	if err != nil {
		return err
	}
	_, err = http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	return nil
}

func isPlainText(ctype string) bool {
	s := strings.Split(ctype, ";")
	if len(s) < 1 {
		return false
	}
	return s[0] == "text/plain"
}
