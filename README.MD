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

we can block and wait until something is sent on a channel using the following syntax.

```go
<-ch
```

### Buffered Channels

you can provide a buffer length as the second argument to `make()` to create a buffered channel

```go
ch := make(chan int, 100)
```

sending on a buffered channel only blocks when the buffer is full.

### Closing Channels

Channels can be explicitly closed by a sender.

```go
ch:= make(chan int)

close(ch)
```

**_Checking if Channel is closed_**

```go
variable,ok := <-ch
```

Similar to the `ok` value when accessing data in a `map`, receivers can check the `ok` value when receiving from a channel to test if a channel was closed.
ok is `false` if channel is empty or closed.

### Select

Sometimes we have a single goroutine listening to multiple channels and want to process data in the order it comes through each channel

A select statement is used to listen to multiple channels at the same time. It is similar to a switch statement but for channels.

```go
select {
	case i, ok := <-ch:
		fmt.Println(i)
	case s, ok := <-ch:
		fmt.Println(s)
	}
```

The first channel with a value ready to be received will fire and its body will execute. If multiple channels a ready at the same time one these randomly. The ok variable in the example above refers to whether or not the channel has been closed by the sender yet.

#### Select Default Case

```go
select {
	case i, ok := <-ch:
		// use i
	default :=
		// receiving from ch would block
		// so do something
	}
```

The `default` case in select statement executes immediately if no other channel has a value ready, A default case stops the `select` statement from blocking.

**_Send Only Channel:_**

`var ch = make(chan<- int)` // send only channel

**_Receive Only Channel:_**

`var ch = make(<-chan int)` // receive only channel

### Mutex in Go

Mutexes allow us to lock access to data. This ensures that we can control which goroutine can access certain data at which time.

It's called 'mutual exclusion' because a mutex excludes different threads (or go Routines) from accessing the same data at the same time.

Go's standard library provides a builtin implementation of a mutex with the `sync.Mutex` type and its two methods.

- `Lock()`
- `Unlock()`

```go
func protected(){
	// This will lock the protected function for first go routine call.
	mux.lock()
	//  After the execution of function it will Unlock, and other routine can access it.
	defer mux.Unlock
	//conclusion is only on routine can access the function at a time.
}
```

we can protect block of code by surrounding it with a call to `Lock` and `Unlock` as shown on the `protected()` method below.

It's good practice to structure the protected code within a function so that defer can be used to ensure that we never forget to unlock the mutex.

## Read-Write Mutex

**IN ADDITION TO THESE METHOD**

- `Lock()`
- `Unlock()`

**The `sync.RWMutex` also has these methods:**

- `RLock()`
- `RUnlock()`

The `sync.RWMutex` can help with performance if we have a read-intensive process. Many goroutines can safely read from the map at the same time (multiple `Rlock()` calls can happen simultaneously). However, only one goroutine can hold a `Lock`() and all `RLock()`'s will also be excluded.
