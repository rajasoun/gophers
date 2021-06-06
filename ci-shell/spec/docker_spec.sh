#shellcheck shell=bash

Describe "System Test : " system
    Include src/load.sh
    Context "devcontainer.sh"
        # It "build_container builds image for .devcontainer"
        #     export VERBOSE=1 
        #     When call build_container
        #     The status should be success
        # End
    End 
End