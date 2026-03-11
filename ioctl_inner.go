//go:build !windows && !solaris && !aix
// +build !windows,!solaris,!aix

package pty

import "syscall"

// Local syscall const values.
const (
	TIOCGWINSZ = syscall.TIOCGWINSZ
	TIOCSWINSZ = syscall.TIOCSWINSZ
)

func ioctlInner(fd, cmd uintptr, ptr any) error {
	_, _, e := syscall.Syscall(syscall.SYS_IOCTL, fd, cmd, ptrToUintptr(ptr))
	if e != 0 {
		return e
	}
	return nil
}
