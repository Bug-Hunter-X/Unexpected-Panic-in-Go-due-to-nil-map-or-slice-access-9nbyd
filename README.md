# Unexpected Panic in Go due to nil map or slice access
This repository demonstrates a common but subtle error in Go: panics caused by accessing non-existent keys in maps or indices in slices.  The program will panic if it tries to access an element that doesn't exist, unexpectedly terminating execution.  The solution provides a robust method for handling these cases using error checking and the `recover` function.

## Problem
Accessing a non-existent key in a map or an out-of-bounds index in a slice in Go causes a runtime panic.  This can be difficult to detect if not properly handled.

## Solution
The solution showcases how to gracefully handle these situations, preventing unexpected panics and ensuring the program continues execution.

### Key Concepts
* **Error Handling:** Explicitly check for the existence of keys in maps before accessing them.
* **`recover` Function:** Use the `recover` function within a `defer` statement to gracefully handle panics and prevent program termination. 

## How to run
1. Clone this repo.
2. Navigate to the repository directory in your terminal.
3. Run `go run bug.go` to see the unexpected panic.
4. Run `go run bugSolution.go` to observe the graceful error handling.