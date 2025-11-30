package emulator

import "errors"

var (
	ErrPTYNotInitialized = errors.New("PTY not initialized")
	ErrInvalidSize       = errors.New("invalid terminal size")
	ErrProcessNotStarted = errors.New("process not started")
	ErrProcessExited     = errors.New("process already exited")
)
