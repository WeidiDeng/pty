//go:build !windows
// +build !windows

package pty

import (
	"os"
	"syscall"
)

// Winsize describes the terminal size.
type Winsize struct {
	Rows uint16 // ws_row: Number of rows (in cells).
	Cols uint16 // ws_col: Number of columns (in cells).
	X    uint16 // ws_xpixel: Width in pixels.
	Y    uint16 // ws_ypixel: Height in pixels.
}

// Setsize resizes t to s.
func Setsize(t *os.File, ws *Winsize) error {
	return ioctl(t, syscall.TIOCSWINSZ, ws)
}

// GetsizeFull returns the full terminal size description.
func GetsizeFull(t *os.File) (size *Winsize, err error) {
	var ws Winsize

	if err := ioctl(t, syscall.TIOCGWINSZ, &ws); err != nil {
		return nil, err
	}
	return &ws, nil
}
