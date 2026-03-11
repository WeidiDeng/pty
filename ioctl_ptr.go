//go:build !windows
// +build !windows

package pty

import "reflect"

// ptrToUintptr returns the uintptr value of a pointer passed as any.
// ptr must be a pointer type or nil. It is intended to be called directly
// as an argument to syscall.Syscall so the conversion happens at the call site.
func ptrToUintptr(ptr any) uintptr {
	if ptr == nil {
		return 0
	}
	v := reflect.ValueOf(ptr)
	switch v.Kind() { //nolint:exhaustive
	case reflect.Ptr, reflect.UnsafePointer:
		return v.Pointer()
	default:
		panic("ptrToUintptr: ptr must be a pointer type")
	}
}
