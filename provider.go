package main

type provider interface {
	hasNew() bool
	get() (string, error)
	getTimestamp() (int64, error)
	put(string) error
}
