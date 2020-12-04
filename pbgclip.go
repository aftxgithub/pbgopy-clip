package main

import (
	"errors"
	"fmt"
	"os"
	"time"
)

func run() error {
	clipboard := &clipboardProvider{}

	addr := os.Getenv("PBGOPY_SERVER")
	if addr == "" {
		return errors.New("put the pbgopy server's address into PBGOPY_SERVER environment variable")
	}
	server := &serverProvider{addr: addr}

	for {
		time.Sleep(1 * time.Second)

		if clipboard.hasNew() {
			if clipboard.lastClipboardTs > server.lastServerTs {
				fmt.Println("Updating server")
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
			}

		} else if server.hasNew() {
			if server.lastServerTs > clipboard.lastClipboardTs {
				fmt.Println("Updating clipboard")
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

			}

		}
	}
}

/*
func run() error {
	serverAddr = os.Getenv("PBGOPY_SERVER")
	if serverAddr == "" {
		return fmt.Errorf("put the pbgopy server's address into PBGOPY_SERVER environment variable")
	}

	for {
		time.Sleep(2 * time.Second)

		clipboardData, err := clipboard.ReadAll()
		if err != nil {
			showError(err)
			continue
		}

		//check when last the server was updated
		resp, err := http.Get(serverAddr + "/lastupdated")
		if err != nil {
			showError(err)
		}

		serverTsBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			showError(err)
		}
		if err := resp.Body.Close(); err != nil {
			showError(err)
		}

		serverTs, err := strconv.ParseInt(string(serverTsBytes), 10, 64)
		if err != nil {
			showError(err)
		}

		if clipboardData != lastClipboardContent {
			lastClipboardContent = clipboardData
			lastClipboardTs = time.Now().UnixNano()
		}

		if serverTs == 0 {
			continue
		}

		if lastClipboardTs > serverTs {
			//send the clipboard data to the server
			req, err := http.NewRequest(http.MethodPut, serverAddr+"/", strings.NewReader(lastClipboardContent))
			if err != nil {
				showError(err)
				continue
			}
			_, err = http.DefaultClient.Do(req)
			if err != nil {
				showError(err)
			}
			continue
		}

		//set the server data as clipboard data if text/plain
		resp, err = http.Get(serverAddr + "/")
		if err != nil {
			showError(err)
			continue
		}
		if resp.Header.Get("Content-Type") != "text/plain" {
			continue
		}
		data, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		clipboard.WriteAll(string(data))
	}
}
*/

func showError(err error) {
	fmt.Fprintf(os.Stderr, err.Error())
}
