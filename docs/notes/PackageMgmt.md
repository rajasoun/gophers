# Go Package managment 

## The Problem

Organize files into different directories for my Go Project. Gain deeper understanding of  Go approach to package management

## Quick Undersatnding of Go Package Management - Theory

> 1. Go considers each directory to be a separate package.
> 1. So moving files actually changes the package it belongs to.  
> 1. Imports working based off directories not invidual files.
> 1. Once the code has been moved to a separate package as part of refactoring, you need to export the types that are accessible outside of the package, including types accessible to the parent package. 

## Quick Undersatnding of Go Package Management - Hands On Lab

Simple example, a command line utility called tips, that can be used as quick reference tool

1. Directory Setup

    ```
    mkdir tips
    mkdir tips/model
    mkdir tips/cli
    ```

    or 

    Use can use single line 

    ```
    mkdir -p tips/{model,cli}
    ```

1. Initialize Go Modules : Most important step as the rest of our setup hinges on this setup. 

    ```
    cd tips
    go mod init github.com/gophers/tips
    ```

    > `github.com/gophers/` is strictly a label and does not require the code be pushed to github before use. **Don't let this confuse you**

1. TDD Tips Project 
    1. Inside model directory create tips_test.go with below code
        ```Go
        package model

        import "strings"

        func GetTip(topic string) string {
            // pass:1 - Make test pass
            // hardcoded
            //:ToDo: Reafactor To Load Tips from JSON
            var result string
            if strings.Compare(topic, "git status") == 0 {
                result = "git status -s"
            } else {
                result = "Tips Not Available for Topic"
            }
            return result
        }
        ```
    1. Run the Test `go test tips/model/*`
    1. Fix the Failing Test 
        ```
        # command-line-arguments [command-line-arguments.test]
        tips/model/tips_test.go:15:10: undefined: GetTip
        FAIL    command-line-arguments [build failed]
        FAIL
        ```
    1. Create a File tips.go
        ```Go
        package model

        import "strings"

        func GetTip(topic string) string {
            // pass:1 - Make test pass
            // hardcoded
            //:ToDo: Reafactor To Load Tips from JSON
            var result string
            if strings.Compare(topic, "git status") == 0 {
                result = "git status -s"
            } else {
                result = "Tips Not Available for Topic"
            }
            return result
        }

        ```

    1. Run Test again `go test tips/model/* `
    1. Run Test again `go test tips/model/* --cover` to get coverage report
    1. Inside cli directory create cli_test.go with below code
        ```Go
        package cli

        import (
            "testing"
        )

        func TestGetTopicFromConsole(t *testing.T) {
            assertEquals := func(t testing.TB, got, want string) {
                t.Helper()
                if got != want {
                    t.Errorf("got %q want %q", got, want)
                }
            }
            t.Run("Get Topic String From Console", func(t *testing.T) {
                got := GetTopic()
                want := "git status"
                assertEquals(t, got, want)
            })
        }
        ```
    1. Run the Test `go test tips/cli/*` and notice like before test fails.
    1. Fix Failed test by creating cli.go inside cli directory
        ```Go
        package cli

        func GetTopic() string {
            // pass:1 - Make test pass
            // hardcoded
            //:ToDo: Get User Input - Mock GetInput in Test
            return "git status"
        }        
        ```
1. Bringing It All Together
    1. Create our main package and main.go file in the root of our directory. 
        > This file should live next to the go.mod file we created earlier by running go mod init. 
        ```Go
        package main

        import (
            "fmt"

            "github.com/gophers/tips/model"
            "github.com/gophers/tips/cli"
        )

        func main() {
            fmt.Println("Main package - main file")
            string topic := cli.GetTopic()
            string tip := model.GetTip(topic)
            fmt.Printf("Tip for %q is $q ", topic, tip)
        }
        ```
1. `go run main.go` to run main file

### Important Tips

1. We are importing entire drectories and not individual files. 
1. Any function in any file in those directories that begins with a capital letter will be exported and publicly available and imported in for use in the above file
1. It is hard to test the main function in Go, cause it doesn’t take any parameters nor returns any values. But if you have the urge to test your main logic anyways, it is a great approach to make the main function as trivial as possible and just call another function

## References

1. [Structure Your Go Project Into Multiple Directories](https://www.jodylecompte.com/posts/go-structure-your-go-project/)
1. [Organizing Go Projects In Sub-Directories With Nested Packages](https://www.jodylecompte.com/posts/go-structure-your-go-project/)