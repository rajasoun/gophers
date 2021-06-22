#shellcheck shell=bash

Describe "Integration Test : " integration
    Include src/load.sh
    Context "os.sh"
        It "raise_error should prompt and exit"
            When run raise_error "Error Occured"
            The status should be failure
            The error should include "Exiting..."
        End
        It "_file_exist LICENSE should pass"
            When call  _file_exist Dockerfile
            The status should be success
            The output should include "Found!"
        End
        It "_file_exist LICENSE should fail"
            When call  _file_exist LICENSE
            The status should be failure
            The error should include "Not Found!"
        End
    End
End