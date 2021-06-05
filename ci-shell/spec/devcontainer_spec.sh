#shellcheck shell=bash

Describe "Intgeration Test : "
    Include src/load.sh
    docker(){
        echo "$@"
    }
    Context "devcontainer.sh"
        It "build_container ???"
            When call build_container
            The status should be success
            The output should include "Dockerfile"
        End
    End 
End