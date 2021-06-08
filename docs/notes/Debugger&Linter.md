## Debugger:
A debugger is a computer program used to test and debug other programs.This is very useful when trying to identify incorrect code and analyze how a program "flows".

So,Debugging tool is the process of detecting and removing of existing and potential errors (also called as "bugs") in a software code.

## How :
1. Set a breakpoint 
1. Start the debugger.
1. Navigate code in the debugger.(with F8 key)

### NOTE:
Delve is a debugger delve debugger command line tool (third-party debugger for the Go programming language ),we navigate the program line by line or through breakpoints, inspect variables, functions and expressions values and finally analyse all our programs in detail.

#### Download and Install Go Delve : $ go get github.com/go-delve/delve/cmd/dlv

## Go Linting:
Linting is the automated checking of source code for programmatic errors. This is done by using a lint tool . It is a static code analysis tool used to flag programming errors, bugs, stylistic errors and suspicious constructs

### Download And Install GO linter: brew install golangci/tap/golangci-lint

### Note: Linting can prevent debugging by catching bugs before you manually run your program. It will run the code and check for errors. Debugging is something you manually do after a bug is found




