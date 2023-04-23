package lockfile_test

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"testing"

	lockfile "github.com/stefanovazzocell/LockFile"
)

const (
	fileNameLetters = "abcdefghijklmnopqrstuvwxyz0123456789"
)

// Returns a path for a temporary file
func getTmpFile() string {
	currentDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	tmpPath := currentDir + string(os.PathSeparator)
	counter := 100
	for {
		randomFileName := strings.Builder{}
		for i := 0; i < 24; i++ {
			randomFileName.WriteByte(fileNameLetters[rand.Intn(len(fileNameLetters))])
		}
		randomFile := tmpPath + randomFileName.String()
		if _, err := os.Stat(randomFile); os.IsNotExist(err) {
			return randomFile
		}
		counter -= 1
		if counter == 0 {
			panic(fmt.Sprintf("failed to find a valid file, last tried is %q", randomFile))
		}
	}
}

// Removes a file
func cleanupFile(path string) {
	if err := os.Remove(path); err != nil && !os.IsNotExist(err) {
		panic(err)
	}
}

func TestLockFile(t *testing.T) {
	tmp := getTmpFile()
	defer cleanupFile(tmp)

	lf, err := lockfile.NewLockFile(tmp)
	if err != nil {
		t.Fatalf("failed to create lock file: %v", err)
	}

	if _, err := lockfile.NewLockFile(tmp); err == nil {
		t.Fatal("failed to error on duplicate lock")
	}

	if err := lf.Free(); err != nil {
		t.Fatalf("failed first free lock file: %v", err)
	}
	if err := lf.Free(); err != lockfile.ErrUnlocked {
		// We should ignore double free
		t.Fatalf("incorrect error for double free call: %v", err)
	}

	lf, err = lockfile.NewLockFile(tmp)
	if err != nil {
		t.Fatalf("failed to re-create lock file: %v", err)
	}
	if err := lf.Free(); err != nil {
		t.Fatalf("failed to free lock file: %v", err)
	}
}
