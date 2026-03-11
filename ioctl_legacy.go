//go:build !windows && !go1.12
// +build !windows,!go1.12

package pty

import "os"

func ioctl(f *os.File, cmd uintptr, ptr any) error {
	return ioctlInner(f.Fd(), cmd, ptr) // fall back to blocking io (old behavior)
}
