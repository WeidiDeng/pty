//go:build !windows && go1.12
// +build !windows,go1.12

package pty

import (
	"os"
	"unsafe"
)

func ioctl(f *os.File, cmd uintptr, ptr unsafe.Pointer) error {
	sc, e := f.SyscallConn()
	if e != nil {
		return ioctlInner(f.Fd(), cmd, ptr) // Fall back to blocking io (old behavior).
	}

	ch := make(chan error, 1)
	defer close(ch)

	e = sc.Control(func(fd uintptr) { ch <- ioctlInner(fd, cmd, ptr) })
	if e != nil {
		return e
	}
	e = <-ch
	return e
}
