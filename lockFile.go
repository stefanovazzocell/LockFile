package lockfile

import (
	"errors"
	"os"
)

var (
	ErrUnlocked = errors.New("this lockFile was already freed")
)

// lockFile represents a file that file-based lock
type lockFile struct {
	file *os.File
}

// NewLockFile tries to create and lock a file at the given path.
// It returns a lockFile object and a possible error.
// If the file already exists, it tries to lock it exclusively.
// If the file is already locked by another process, it returns an error.
func NewLockFile(path string) (lockFile, error) {
	// Attempt to open/create file
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0640)
	if err != nil {
		return lockFile{}, err
	}
	// Attempt to place a lock
	err = LockFile(file)
	if err != nil {
		_ = file.Close()
		return lockFile{}, err
	}
	// Done
	return lockFile{
		file: file,
	}, nil
}

// Free unlocks and closes the lockFile.
// It returns a possible error.
func (lock *lockFile) Free() error {
	// Check if locked
	if lock.file == nil {
		// No file to unlock
		return ErrUnlocked
	}
	// Unlock
	if err := UnlockFile(lock.file); err != nil {
		return err
	}
	// Close file
	if err := lock.file.Close(); err != nil {
		return err
	}
	// Optionally delete file
	_ = os.Remove(lock.file.Name())
	// Reset
	lock.file = nil
	return nil
}
