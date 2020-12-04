package main

import (
	"time"

	"github.com/atotto/clipboard"
)

type clipboardProvider struct {
	lastClipboardContent string
	lastClipboardTs      int64
}

func newClipboardProvider() *clipboardProvider {
	return &clipboardProvider{
		lastClipboardTs: time.Now().UnixNano(),
	}
}

func (c *clipboardProvider) hasNew() bool {
	data, err := clipboard.ReadAll()
	if err != nil {
		return false
	}
	if data != c.lastClipboardContent {
		c.lastClipboardTs = time.Now().UnixNano()
		return true
	}
	return false
}

func (c *clipboardProvider) get() (string, error) {
	data, err := clipboard.ReadAll()
	if err != nil {
		return "", err
	}
	c.lastClipboardContent = data
	return data, nil
}

func (c *clipboardProvider) getLastTimestamp() int64 {
	return c.lastClipboardTs
}

func (c *clipboardProvider) put(data string) error {
	return clipboard.WriteAll(data)
}
