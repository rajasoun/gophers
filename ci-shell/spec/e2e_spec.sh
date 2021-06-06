#shellcheck shell=bash

Describe "e2e System Test : " iaac
    Include src/load.sh
    Context "e2e - check packages : "
        It "check_common_packages"
            When call check_common_packages 
            The output should include "All Passed!"
        End
        It "check_dev_packages"
            When call check_dev_packages 
            The output should include "All Passed!"
        End
        It "check_toolz_packages"
            When call check_toolz_packages 
            The output should include "All Passed!"
        End
        It "check_system_details"
            When call check_system_details 
            The output should include "Passed!"
        End
        It "_common_packages function with valid package"
            PACKAGE_LIST="git less"
            # shellcheck disable=SC2086
            When call _common_packages "common-os-packages" ${PACKAGE_LIST}
            The status should be success
            The output should include "Passed!"
        End 
    End
    Context "e2e - user-management : "
        It "user vscode exists"
            When call whoami
            The output should equal "vscode"
        End
    End
End

Describe "Unit Test : e2e.sh" unit
    Include src/load.sh
    It "echo_std_err function"
        When call echo_std_err "Failure"
        The error should include "Failure"
    End
    It "check function with uname"
        When call check uname
        The output should include "Passed!"
    End
    It "check_deatil function with \"zsh\" zsh --version"
        When call check_detail "zsh" zsh --version
        The output should include "Passed!"
    End
    It "check_deatil function with zsh --version1 should fail"
        When call check_detail "zsh" zsh --version1
        The status should be failure
        The output should include "Testing zsh"
        The error should include "Failed!"
    End
    It "_common_packages function with invalid packages should fail"
        PACKAGE_LIST="dummy"
        When call _common_packages "common-os-packages" ${PACKAGE_LIST}
        The status should be failure
        The output should include "Testing"
        The error should include "Failed!"
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