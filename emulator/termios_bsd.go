//go:build darwin || freebsd || openbsd || netbsd || dragonfly

package emulator

import "golang.org/x/sys/unix"

func termiosRequests() (reqGet, reqSet uint) {
	return unix.TIOCGETA, unix.TIOCSETA
}
