# Tips

## The Problem

As developers we need to remember several commands related to git, docker, tdd, general development environment setup making it overwhelming

## Solution

Command Line Tool to provide tips on the command to be used based on the topic

## Install
```
$ make build

```

## Features
```
A Tips tool with an easy terminal/User interface.

tips
tips <flag>
tips <command>
tips --help
tips --version
tips git/docker 
tips git/docker <command>
tips git/docker <command> <flag> 

```
## Usage

### Tips tool Usage
```
$ tips

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
### Docker Usage
```
$ tips docker

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

### Git Usage
```
$ tips git

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
###  User can get a tip , on giving a command

```

1. Bypass ps docker command in terminal like [tips docker <command>]

$ tips docker ps
"List all containers : docker ps -a" 

2. Bypass log git tool command in terminal [tips git <command>]

$ tips git log
"Search change by content : git log -S'<a term in the source>'"

3. Help Command (set the help flag)-> $ tips --help , $ tips git/docker -h 

```

## Libraries 

1. Cobra library  is used to build Tips command line app [cli].
2. Logrus library is used to set the log status (i.e debug).
3. We followed TDD design while building the Tips cli app, Also used Testify library for testing test cases.


##  Running Notes

1.	The tips project describes about the commands/tips (as in Linux/Git/docker) used in command line, where users can get commands based on request.
1.	Tips project usually gives commands related to Linux/Git(as terminals in other editors)
1.	Code is written using VS code editor in Go programming language
1.	It covers unit tests using TDD (Test Driven Development).
1.	Tips project is designed by using a framework-MVC (Model View Controller) architectural design pattern.
1.	MVC consists of three main packages as below
*	**View**--cmd package(Taking input(title for tip) from console and passing to controller package)
*	**Model**-model package(Taking title parameter from controller and based on title return proper tip to Controller)
*	**Controller**—controller package(Prints the tips on console according to user input(title))
 
 
## Git-tips
[Git-Tips]( https://github.com/rajasoun/tips/blob/main/GitTips.md)