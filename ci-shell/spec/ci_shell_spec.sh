#shellcheck shell=bash

Describe "System Test : "
    Include lib/validators.sh
    Context "ci-shell - common-os-packages : "
        It "common-os-packages"
            When call check_common_os_packages 
            The output should include "All Passed!"
        End
    End
    Context "ci-shell - user-management : "
        It "user ci-shell exists"
            When call whoami
            The output should equal "vscode"
        End
    End
End

Describe "Unit Test : validators.sh"
    Include lib/validators.sh
    It "echo_std_err function"
        When call echo_std_err "Failure"
        The error should include "Failure"
    End
    It "check function with uname"
        When call check uname
        The output should include "Passed!"
    End
    It "check function with dummy should fail"
        When call check dummy
        The status should be failure
        The error should include "Failed!"
    End
    It "report_results with empty failures"
        FAILED=()
        When call report_results "${FAILED[@]}"
        The output should include "All Passed!"
    End
    It "report_results with failure"
        FAILED=(failure)
        When call report_results "${FAILED[@]}"
        The output should include "failure"
        The error should include "Failed tests: 1"
        The status should be failure
    End
End