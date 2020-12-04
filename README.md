# pbgopy-clip

Automatically sync clipboard with multiple devices.

> You can basically CTRL C on one device and CTRL V on another :)

![Go](https://github.com/thealamu/pbgopy-clip/workflows/Go/badge.svg)

## Requisites

This program depends on [atotto/clipboard](https://github.com/atotto/clipboard).
You need to have a clipboard utitlity available. Please install xsel, xclip, wl-clipboard or Termux:API add-on for termux-clipboard-get/set.
<br>
<br>
Install [pbgopy](https://github.com/nakabonne/pbgopy), pbgopy-clip syncs clipboard data over the pbgopy server.

## Usage

You need to start the pbgopy server. It listens on port 9090 by default.

```shell
pbgopy serve
```

Export the address of the server on each device you want their clipboard synced.

```shell
export PBGOPY_SERVER=http://foo.bar:9090
```

Start `pbgopy-clip` on each of the devices too

```shell
pbgopy-clip
```

As long as pbgclip runs, the clipboard of the device will be kept in sync with the other devices,
using the pbgopy server as a backend.

## Build

### Go

```shell
go get github.com/thealamu/pbgopy-clip
```
