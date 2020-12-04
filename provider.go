package main

type provider interface {
	hasNew() bool
	get() (error, string)
	getTimestamp() (error, int64)
	put(string) error
}
