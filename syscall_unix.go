//go:build !windows

package lockfile

import (
	"os"
	"syscall"
)

// Locks a given file
func LockFile(file *os.File) error {
	return syscall.Flock(int(file.Fd()), syscall.LOCK_EX|syscall.LOCK_NB)
}

// Unlocks a given file
func UnlockFile(file *os.File) error {
	return syscall.Flock(int(file.Fd()), syscall.LOCK_UN)
}
