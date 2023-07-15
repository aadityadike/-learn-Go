# Learning Go

## Sync goroutines with waitGroups:

In sync package we have waitGroup Struct that contain few function like Add, Wait & Done this functions can be useful to sync with main function. In our case if we remove for loop. the main function did not wait for other goroutine functions to execute it executes line by line and since the goroutine has created a new thread the main function skip that execution and follow the flow of program.

## What is the difference between other languages and go in terms of concurrency?

### Goroutine

- Go is using, what's called a "Green Thread".
- Abstraction of an actual Thread.
- Managed by the go runtime, we are only interacting with high level goroutine.
- cheaper & Lightweight

### OS Thread

- Managed by kernel.
- Are Hardware dependent
- achieving concurrency can be complex compare to Go
- high start up time

## Defer Keyword

Defer keyword means, we are telling function to execute the particular line of code or function at the end or before return statement.
The benefit is that no matter from where we going to return, the defer code will execute just before the function execution finishes.

## Go Build

Go Build creates, an executable file that we can run on any machine.
and we can execute that file with command `./executable_name`

## Go Install

Go Install will Install your code globally, in short you don't need to specify path and you can code file from any directory.
command after installing file in your machine to run executable `executable_name` from any directory.

## Go Mod

Go Mod is dependence management tool in go, since go accepts remote urls to import depnedinces we need specify remote url.

### "replace" Keyword in Go Mod

replace key word will follow the following syntax : `remote_URL_of_package_that you_want_to_import => PATH_of_local_package`

what this will do is, Go will not go on remote url to find package it will get package from the path of local environment that we specify.
if the right-hand side of a replace directive is a filesystem path, then the target must have a go.mod file at that location. If the go.mod file is not present, you can create one with go mod init.

### "require" Keyword in Go Mod

replace key word will follow the following syntax : `remote_URL_of_package_that you_want_to_import`

This is used to import a other module or package.

## Go get

`Go get remote_URL` this will basically get third party dependencies in your local environment.
