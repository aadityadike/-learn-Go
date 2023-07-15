#### Sync goroutines with waitGroups:

In sync package we have waitGroup Struct that contain few function like Add, Wait & Done this functions can be useful to sync with main function. In our case if we remove for loop. the main function did not wait for other goroutine functions to execute it executes line by line and since the goroutine has created a new thread the main function skip that execution and follow the flow of program.

#### What is the difference between other languages and go in terms of concurrency?

#### Goroutine

- Go is using, what's called a "Green Thread".
- Abstraction of an actual Thread.
- Managed by the go runtime, we are only interacting with high level goroutine.
- cheaper & Lightweight

#### OS Thread

- Managed by kernel.
- Are Hardware dependent
- achieving concurrency can be complex compare to Go
- high start up time

#### Defer Keyword

Defer keyword means, we are telling function to execute the particular line of code or function at the end or before return statement.
The benefit is that no matter from where we going to return, the defer code will execute just before the function execution finishes.

### Go Build

Go Build creates, an executable file that we can run on any machine.
and we can execute that file with command `./executable_name`

### Go Install

Go Install will Install your code globally, in short you don't need to specify path and you can code file from any directory.
command after installing file in your machine to run executable `executable_name` from any directory.
