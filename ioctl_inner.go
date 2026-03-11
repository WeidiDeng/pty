//go:build !windows && !solaris && !aix
// +build !windows,!solaris,!aix

package pty

import (
	"reflect"
	"syscall"
	"unsafe"
)

// Local syscall const values.
const (
	TIOCGWINSZ = syscall.TIOCGWINSZ
	TIOCSWINSZ = syscall.TIOCSWINSZ
)

func ioctlInner(fd, cmd uintptr, ptr any) error {
	var p unsafe.Pointer
	if ptr != nil {
		p = reflect.ValueOf(ptr).UnsafePointer()
	}
	_, _, e := syscall.Syscall(syscall.SYS_IOCTL, fd, cmd, uintptr(p)) //nolint:gosec // ptr-to-uintptr at syscall site.
	if e != 0 {
		return e
	}
	return nil
}
