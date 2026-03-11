//go:build !windows && !solaris && !aix
// +build !windows,!solaris,!aix

package pty

import (
	"syscall"
	"unsafe"
)

// Local syscall const values.
const (
	TIOCGWINSZ = syscall.TIOCGWINSZ
	TIOCSWINSZ = syscall.TIOCSWINSZ
)

func ioctlInner(fd, cmd uintptr, ptr unsafe.Pointer) error {
	_, _, e := syscall.Syscall(syscall.SYS_IOCTL, fd, cmd, uintptr(ptr)) //nolint:gosec // ptr-to-uintptr at syscall site.
	if e != 0 {
		return e
	}
	return nil
}
