#shellcheck shell=bash

Describe "e2e System Test : " e2e system
 
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
