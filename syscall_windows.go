//go:build windows

package lockfile

import (
	"os"

	"golang.org/x/sys/windows"
)

// Locks a given file
func LockFile(file *os.File) error {
	return windows.LockFileEx(windows.Handle(file.Fd()), 3, 0, 1, 0, &windows.Overlapped{})
}

// Unlocks a given file
func UnlockFile(file *os.File) error {
	return windows.UnlockFileEx(windows.Handle(file.Fd()), 0, 1, 0, &windows.Overlapped{})
}
