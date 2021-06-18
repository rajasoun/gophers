## Packaging 
[Build Go Executable](https://www.digitalocean.com/community/tutorials/how-to-build-go-executables-for-multiple-platforms-on-ubuntu-16-04)
[packaging in go](https://github.com/golang/go/wiki/WindowsCrossCompiling)
  
The packaging is used to design and maintain a large number of programs by grouping related features together into single units
1. so that they can be easy to maintain and understand and independent of the other package programs.
1. This modularity allows them to share and reuse.

It is the process of bundling together all files, components, and information in an application executable file or package 
 that are required to install, start, execute, and run an application in a specific environment.

## Continuous Integration & Continuous Deployment
[CI/CD pipeline](https://brunopaz.dev/blog/building-a-basic-ci-cd-pipeline-for-a-golang-application-using-github-actions)

[Github actions](https://docs.github.com/en/actions/guides/about-continuous-integration)

Continuous integration (CI) is a software practice that requires frequently committing code to a shared repository.Frequent code updates also make it easier to merge changes from different members of a software development team. 
When you commit code to your repository, you can continuously build and test the code to make sure that the commit doesn't introduce errors. Your tests can include code linters (which check style formatting), security checks, code coverage, functional tests, and other custom checks.

#### Note:
 Building and testing your code requires a server. You can build and test updates locally before pushing code to a repository, or you can use a CI server that checks for new code commits in a repository.


### CI/CD with using Github action:

1. GitHub Actions are based on the concept of Workflows. A workflow is nothing more than a set of jobs and steps that are executed when some condition or event is met. (Ex: a push to the repository, a pull request, a deployment, etc).

1. GitHub runs  CI tests and provides the results of each test in the pull request, so we can see whether the change in your branch introduces an error. When all CI tests in a workflow pass, the changes you pushed are ready to be reviewed by a team member or merged. When a test fails, one of your changes may have caused the failure.

1. Use GitHub Actions to create workflows across the full software development life cycle. For example, we can use actions to deploy, package, or release your project.
