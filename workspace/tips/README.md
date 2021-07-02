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

## Steps to build an app
1. Read UseCases
1. Make a Flow of Requirements(useCases)
1. Do coding with TDD 
1. Refactor the code 
1. Add one by one requirements with TDD(make pass unit test case)
1. Follow MVC Pattern(Check the app flow with MVC Pattern)
1. Also write Integration and End to End tests if required.
1. Make 100% code coverage in Testing
1. Apply Continous Integration and Continous Developement with Github Action on GitHub.
1. Build a packaging (build an .exe file for different platform)
1. Deployment / Release an App

### NOTE: 
1. The Code Design should be loosely coupled and hightly cohesive .So, The Code will be maintainable .
[click](https://medium.com/clarityhub/low-coupling-high-cohesion-3610e35ac4a6)
[click](https://blog.learngoprogramming.com/packages-can-allow-or-disallow-for-reusability-2edb6bd18815)

1. Unit testing code should be separated from Integration testing code .
[click](https://mickey.dev/posts/go-build-tags-testing/)

1. When we need to do mocking in the code :
  > 1. when OS(operating system) packages involve in the code.(like readFile)
  > 1. when we may not able to do with DI(like io.Reader,io.writer) which can be used for both implementations i.e main functinality as well as testing .
  > 1. when we have to increase the code coverage.
  Because sometime, Mocking increases code complexity and less readability. 
  
1. About MakeFile[click](https://www.youtube.com/watch?v=QztvWSCbQLU)
