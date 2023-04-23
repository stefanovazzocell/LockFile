#  LockFile
A cross-platform Go library that creates a lock file.

## Utilization

You can either use the full-auto mode or on a self-service basis.

Full-auto example:

```go
fLock, err := NewLockFile(path + ".lock")
if err != nil {
    return err
}
defer fLock.Free()

// Do work
```

If you want to modify a locked file or prefer a more self-service approach, here's an example for you:

```go
// Open File
file, err := os.Open("/path/to/file")
if err != nil {
	return err
}
// Lock file
if err := lockfile.LockFile(file); err != nil {
	return err
}
// Write data
if _, err := file.Write(data); err != nil {
	_ = lockfile.UnlockFile(file)
	_ = file.Close()
	return err
}
if err := file.Sync(); err != nil {
	_ = lockfile.UnlockFile(file)
	_ = file.Close()
	return err
}
// Free and unlock
_ = lockfile.UnlockFile(file)
return file.Close()
```

## Caveats

- The locking logic will only work if all entities accessing a given lock file use this API.
- This will likely not work on shared drives.
- This library won't work on `wasm`/`js` builds.

This library has been tested on:
- `Linux/amd64` (Fedora Linux)
- `windows/amd64` (Windows 10 VM)
- `darwin/arm64` (M1 MacBook)