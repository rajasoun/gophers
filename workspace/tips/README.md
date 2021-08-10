# Tips

## The Problem

As developers we need to remember several commands related to git, docker, tdd, general development environment setup making it overwhelming

## Solution

Command Line Tool to provide tips on the command to be used based on the topic

## Usage
```
  tips provides help for docker and git cli commands

Usage:
  tips [flags]
  tips [command]

Examples:
-> tips <tool_name> <command>

tips git push
tips docker ps

Available Commands:
  completion  generate the autocompletion script for the specified shell
  docker      Docker provides the ability to package and run an application.
  git         Git is a DevOps tool used for source code management.
  help        Help about any command

Flags:
      --config string   config file (default is $HOME/.tips.yaml)
  -h, --help            help for tips
  -v, --version         version for tips

Use "tips [command] --help" for more information about a command.
```

```
"Docker is a software platform that simplifies the process of building, running,
managing and distributing applications."

Usage:
  tips docker [flags]

Examples:
tips docker <command>

tips docker ps
"List all containers : docker ps -a "

Flags:
  -h, --help      help for docker
  -v, --version   version for docker

Global Flags:
      --config string   config file (default is $HOME/.tips.yaml)
```
```
"Git is used to tracking changes in the source code,
 enabling multiple developers to work together on non-linear development"

Usage:
  tips git [flags]

Examples:
tips git <command>

tips git stash
"Saving current state of unstaged changes to tracked files : git stash -k" 

Flags:
  -h, --help      help for git
  -v, --version   version for git

Global Flags:
      --config string   config file (default is $HOME/.tips.yaml)
```
```
###  User can get a tip on giving a command of the tool (i.e git,docker)

1. Bypass ps docker command in terminal like [tips docker <command>]

tips docker ps
"List all containers : docker ps -a" 

2. Bypass log git tool command in terminal [tips git <command>]

tips git log
"Search change by content : git log -S'<a term in the source>'"

3. Help Command-> tips --help ,tips git/docker -h 
```
## Libraries 

1. We used Cobra library to build Tips command line app [cli].
2. We used Logrus library to set the log status (i.e debug).
3. We followed TDD design while building the Tips cli app, Also used Testify library for checking test cases.


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
