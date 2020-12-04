package main

type provider interface {
	hasNew() bool
	get() (string, error)
	getLastTimestamp() (int64, error)
	put(string) error
}
