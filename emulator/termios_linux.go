//go:build linux

package emulator

import "golang.org/x/sys/unix"

func termiosRequests() (reqGet, reqSet uint) {
	return unix.TCGETS, unix.TCSETS
}
