//go:build !windows && !go1.12
// +build !windows,!go1.12

package pty

import (
	"os"
	"reflect"
	"syscall"
)

func ioctl(f *os.File, cmd uintptr, ptr interface{}) error {
	return ioctlInner(f.Fd(), cmd, ptr) // fall back to blocking io (old behavior)
}

func ioctlInner(fd, cmd uintptr, ptr interface{}) error {
	var e syscall.Errno
	if ptr == nil {
		_, _, e = syscall.Syscall(syscall.SYS_IOCTL, fd, cmd, uintptr(0))
	} else {
		_, _, e = syscall.Syscall(syscall.SYS_IOCTL, fd, cmd, reflect.ValueOf(ptr).UnsafeAddr()) //nolint:gosec // ptr-to-uintptr at syscall site.
	}
	if e != 0 {
		return e
	}
	return nil
}
