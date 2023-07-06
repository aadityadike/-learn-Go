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
