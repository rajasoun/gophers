#shellcheck shell=bash

Describe "Integration Test : " integration
    Include src/load.sh
    Context "git.sh"
        It "_is_git_repo should pass"
            When call  _is_git_repo 
            The status should be success
            The output should include "true"
        End
        It "_is_git_repo should fail for /tmp"
            cd /tmp || return 1
            When call  _is_git_repo 
            The status should be failure
            The error should include "false"
        End
        It "_get_git_workspace should pass for current git directory"
            When call _get_git_workspace
            The status should be success
            The output should include "workspace"
        End 
        It "_get_git_workspace should fail for /tmp directory"
            cd /tmp || return 1
            When call _get_git_workspace
            The status should be failure
            The error should include "false"
        End 
        It "_get_app_name_from_git_workspace equals git project name"
            When call _get_app_name_from_git_workspace
            GIT_CMD=$(git remote get-url origin | xargs basename)
            The status should be success
            The output should equal  "$GIT_CMD"
        End
        It "_get_app_name_from_git_workspace on /tmp fails"
            cd /tmp || return 1
            When run _get_app_name_from_git_workspace
            The status should be failure
            The error should include  "Not a Valid Git Repo!"
            The error should include  "App Name is Empty!"
        End
    End
End