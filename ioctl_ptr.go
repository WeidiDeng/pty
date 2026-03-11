//go:build !windows
// +build !windows

package pty

import (
	"reflect"
	"unsafe"
)

// ptrToUnsafePointer converts a pointer value (passed as any) to an unsafe.Pointer for use in syscalls.
// The final uintptr conversion must be deferred to the syscall.Syscall argument to satisfy the Go GC rules.
func ptrToUnsafePointer(ptr any) unsafe.Pointer {
	if ptr == nil {
		return nil
	}
	return unsafe.Pointer(reflect.ValueOf(ptr).Pointer()) //nolint:gosec // Safe: reflect.Value keeps ptr alive; uintptr immediately converted back.
}
