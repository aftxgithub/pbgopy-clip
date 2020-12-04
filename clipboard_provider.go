package main

import "github.com/atotto/clipboard"

type clipboardProvider struct {
	lastClipboardContent string
	lastClipboardTs      int64
}

func (c *clipboardProvider) hasNew() bool {
	data, err := clipboard.ReadAll()
	if err != nil {
		return false
	}
	if data != c.lastClipboardContent {
		return true
	}
	return false
}

func (c *clipboardProvider) get() (error, string) {
	data, err := clipboard.ReadAll()
	if err != nil {
		return err, ""
	}
	c.lastClipboardContent = data
	return nil, data
}

func (c *clipboardProvider) getTimestamp() (error, int64) {
	return nil, c.lastClipboardTs
}

func (c *clipboardProvider) put(data string) error {
	return clipboard.WriteAll(data)
}
