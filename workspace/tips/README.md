# Tips

## The Problem

As developers we need to remember several commands related to git, tdd, general development environment setup making it overwhelming

## Solution

Command Line Tool to provide tips on the command to be used based on the topic

## Usage

```
  Usage
      $ tip [options]
  Options
      --help    Provides usage help (Shows the current page)
      --all     Gives all the git tips
      <keyword> Gives the git tips consisting of the keyword
  Examples
      $ tip bypass

      1. Bypass pre-commit and commit-msg githooks
      => git commit --no-verify

      $ tip

      Git Tip of the Terminal
      -------------------------
      Saving current state of tracked files without commiting
      => git stash
```

##  Running Notes

1.	The tips project describes about the commands/tips (as in Linux/Git) used in command line, where users can get commands based on request.
1.	Tips project usually gives commands related to Linux/Git(as terminals in other editors)
1.	Code is written using VS code editor in Go programming language
1.	It covers unit tests using TDD (Test Driven Development).
1.	Tips project is designed by using a framework-MVC (Model View Controller) architectural design pattern.
1.	MVC consists of three main packages as below
*	**View**--cli package(Taking input(title for tip) from console and passing to controller package)
*	**Model**-model package(Taking title parameter from controller and based on title return proper tip to Controller)
*	**Controller**—controller package(Prints the tips on console according to user input(title))
 
 #### Example: We can run an app from main function (go run main.go) and get a tip on user input.

* Example of practical learnings link is given below.
[Code for git-tips](https://github.com/rajasoun/gophers/tree/mvc_design_pattern/workspace/tips)


## Git-tips
[Git-Tips]( https://github.com/rajasoun/tips/blob/main/GitTips.md)




