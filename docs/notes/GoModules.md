# Go Modules

>[GoModule](https://blog.golang.org/using-go-modules)
>[GoDoc](https://golang.org/doc/modules/gomod-ref)

## Module
1. A module is a collection of related Go packages that are versioned together as a single unit with a go.mod file at its root.
1. Each Go module is defined by a go.mod file 
1. It describes the module's properties, including its dependencies on other modules and on versions of Go.
1. Go generates a go.mod file when you run the "go mod init command". 
   Example: go mod init example.com/mymodule

## Syntax in go.mod file
   >syntax: module module-path
   
  Here, module-path :The module's module path which is also the import path used for the  root directory
  The module path should be a path from which Go tools can download the module source. 

### Note:
    In addition to go.mod, the go command maintains a file named go.sum containing the expected cryptographic hashes of the content of specific module versions:

## MAIN COMMANDS FOR GO MODULES

1. go get:Add dependencies to current module and install them.
1. go mod init: Initializes a module in the current directory. If the current directory is available, the module path is initialized with the current path in the GOPATH.
1. go mod tidy: This command resets the module configuration to the source code. Dependencies that are no longer needed are removed, transitive dependencies are updated and cleaned up.
1. go list -m all  :   Prints the current module’s dependencies.
1. go mod vendor   :   Make vendored copy of dependencies
1. go mod download :   Download modules to local cache
1. go mod edit     :   Edit go.mod from tools or scripts
1. go mod graph    :   Print module requirement graph
1. go mod why      :   Explain why packages or modules are needed


## Gopls (update new version of gopls v0.6.11 for Multiple Modules at same time)

>[Multiple Module](https://github.com/golang/tools/tree/master/gopls)

1. gopls (pronounced "Go please") is the official Go language server developed by the Go team. It provides IDE features to any LSP-compatible editor.
1. GO111MODULE=on go get golang.org/x/tools/gopls@latest
1. gopls supports both Go module and GOPATH modes, but if you are working with multiple modules or uncommon project layouts, you will need to specifically configure our workspace.