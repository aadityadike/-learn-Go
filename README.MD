# Learning GoLang

## Sync Goroutines with WaitGroups:

In Go, the `sync` package provides the `WaitGroup` struct, which contains a few functions like `Add`, `Wait`, and `Done`. These functions are useful for synchronizing goroutines with the main function. If we remove the for loop, the main function won't wait for other goroutine functions to execute. It will execute line by line, and since the goroutine has created a new thread, the main function may skip that execution and follow the flow of the program.

## What is the difference between Go and other languages in terms of concurrency?

### Goroutine

- Go uses what's called a "Goroutine," which is an abstraction of an actual thread.
- Goroutines are managed by the Go runtime, and developers interact with high-level goroutines.
- Goroutines are cheaper and lightweight compared to OS threads.

### OS Thread

- OS threads are managed by the kernel.
- They are hardware-dependent and achieving concurrency with OS threads can be more complex compared to Go's goroutines.
- OS threads have a higher start-up time.

## Defer Keyword

The `defer` keyword in Go is used to instruct a function to execute a particular line of code or function at the end or just before the return statement. The benefit of using `defer` is that no matter where we return from within the function, the `defer` code will always execute just before the function execution finishes.

## Go Build

The `go build` command creates an executable file that can be run on any machine. To execute the compiled file, you can use the command `./executable_name`.

## Go Install

The `go install` command installs your Go code globally, making it accessible from any directory without the need to specify the path. After installing the file on your machine, you can run the executable using its name from any directory.

## Go Mod

Go Mod is a dependency management tool in Go. As Go allows importing dependencies from remote URLs, we need to specify the remote URL.

### "replace" Keyword in Go Mod

The `replace` keyword follows the following syntax: `remote_URL_of_package => PATH_of_local_package`. This directive tells Go to fetch the package from the specified local path instead of the remote URL. If the right-hand side of a replace directive is a filesystem path, then the target must have a `go.mod` file at that location. If the `go.mod` file is not present, you can create one with `go mod init`.

### "require" Keyword in Go Mod

The `require` keyword follows the following syntax: `remote_URL_of_package`. This is used to import another module or package into your project.

## Go Get

The `go get remote_URL` command is used to fetch third-party dependencies and download them into your local environment.

## Concurrency VS Synchronous

### Synchronous Code

Synchronous code is executed sequentially, one after the other. Each operation must wait for the previous one to complete before proceeding.

### Concurrent Code

Concurrent code, in Go, refers to code that can handle multiple tasks simultaneously through goroutines and channels, achieving concurrency.

## Channels

**_Creating A Channel:_**

```go
ch := make(chan int)
```

**_Sending Data To A Channel:_**

```go
ch <- 30
```

**_Receiving Data From A Channel:_**

```go
v := <- ch
```

Channels are used for communication and synchronization between goroutines. They allow data to be passed between goroutines securely.
