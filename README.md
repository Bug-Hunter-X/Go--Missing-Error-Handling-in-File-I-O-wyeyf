# Go: Missing Error Handling in File I/O

This repository demonstrates a common bug in Go programs: neglecting to check for errors returned by functions. This specific example focuses on file I/O operations, where failing to handle errors can lead to silent data corruption, unexpected program termination, or misleading results.

**Bug:** The `bug.go` file contains a Go program that attempts to read from a file. However, it fails to check the return value of the `os.ReadFile` function. If the file doesn't exist or there's another problem reading it, the program will proceed without reporting the error. This can make it difficult to diagnose issues later.

**Solution:** The `bugSolution.go` file shows the corrected version, which properly checks for errors.  Always check the error returned by file-related and network functions to ensure robust operation and easier debugging.