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
