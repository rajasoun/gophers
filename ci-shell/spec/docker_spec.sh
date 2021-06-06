#shellcheck shell=bash

Describe "System Test : " ci-build
    Include src/load.sh
    Context "docker.sh"
        It "build_container builds image for .devcontainer" 
            When call build_container
            The status should be success
            The output should include "DONE"
        End
        It "build_container_and_time_it builds image for .devcontainer and time it" 
            When call build_container_and_time_it
            The status should be success
            The output should include "Local Build Success"
        End
        It "tear_down  .devcontainer " 
            When call tear_down
            The status should be success
            The output should include "Removing Container Image"
            The output should include "Deleted"
        End
    End 
End