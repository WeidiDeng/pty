//go:build !windows && !solaris && !aix && go1.12
// +build !windows,!solaris,!aix,go1.12

package pty

import (
	"reflect"
	"syscall"
)

// Local syscall const values.
const (
	TIOCGWINSZ = syscall.TIOCGWINSZ
	TIOCSWINSZ = syscall.TIOCSWINSZ
)

func ioctlInner(fd, cmd uintptr, ptr any) error {
	var e syscall.Errno
	if ptr == nil {
		_, _, e = syscall.Syscall(syscall.SYS_IOCTL, fd, cmd, uintptr(0))
	} else {
		_, _, e = syscall.Syscall(syscall.SYS_IOCTL, fd, cmd, uintptr(reflect.ValueOf(ptr).UnsafePointer())) //nolint:gosec // ptr-to-uintptr at syscall site.
	}
	if e != 0 {
		return e
	}
	return nil
}
