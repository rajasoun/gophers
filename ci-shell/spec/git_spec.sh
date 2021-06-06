#shellcheck shell=bash

Describe "Integration Test : "
    Include src/load.sh
    Context "os.sh"
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
    End
End