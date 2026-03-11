//go:build !windows
// +build !windows

package pty

import "reflect"

// ptrToUintptr converts a pointer value (passed as any) to its uintptr representation for use in syscalls.
func ptrToUintptr(ptr any) uintptr {
	if ptr == nil {
		return 0
	}
	return reflect.ValueOf(ptr).Pointer()
}
